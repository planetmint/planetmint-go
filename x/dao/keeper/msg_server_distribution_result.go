package keeper

import (
	"context"
	"sort"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// type Claims struct {
// 	challenger map[string]uint64
// 	challengee map[string]uint64
// 	initiator  map[string]uint64
// }

func (k msgServer) DistributionResult(goCtx context.Context, msg *types.MsgDistributionResult) (*types.MsgDistributionResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	distribution, found := k.LookupDistributionOrder(ctx, msg.GetLastPop())
	if !found {
		errorMessage := types.ErrDistributionNotFound.Error() + " for provided block height " + strconv.FormatInt(msg.GetLastPop(), 10)
		util.GetAppLogger().Error(ctx, errorMessage)
		return nil, errorsmod.Wrap(types.ErrDistributionNotFound, errorMessage)
	}

	distribution.DaoTxID = msg.DaoTxID
	distribution.PopTxID = msg.PopTxID
	distribution.InvestorTxID = msg.InvestorTxID
	distribution.EarlyInvAddr = msg.EarlyInvestorTxID
	distribution.StrategicTxID = msg.StrategicTxID

	if err := k.clearUnresolvedClaims(ctx, distribution.FirstPop); err != nil {
		util.GetAppLogger().Error(ctx, "error while clearing unresolved claims for heights %d-%d: %v", distribution.FirstPop, distribution.LastPop, err)
	}

	err := k.resolveStagedClaims(ctx, distribution.FirstPop, distribution.LastPop)
	if err != nil {
		util.GetAppLogger().Error(ctx, "%s for provided PoP heights: %d %d", types.ErrResolvingStagedClaims.Error(), distribution.FirstPop, distribution.LastPop)
		return nil, errorsmod.Wrap(types.ErrConvertClaims, err.Error())
	}
	util.GetAppLogger().Info(ctx, "staged claims successfully for provided PoP heights: %d %d", distribution.FirstPop, distribution.LastPop)
	k.StoreDistributionOrder(ctx, distribution)

	return &types.MsgDistributionResultResponse{}, nil
}

// TODO: only do this for challenger/challengee

// clearUnresolvedClaims checks for all Challenge participants starting from a given height.
// An accounts stagedDenom amount should always be 0 except for claims that have not yet been reissued.
// Calculate the difference for a set of participants and clear out all past unresolved staged claims.
func (k msgServer) clearUnresolvedClaims(ctx sdk.Context, start int64) (err error) {
	// calculate total amounts for current and future claims
	currentAmounts, err := k.getClaims(ctx, start, ctx.BlockHeight())
	if err != nil {
		return err
	}

	totalAmounts := make(map[string]uint64)
	for participantAddress := range currentAmounts {
		stagedBalance := k.bankKeeper.GetBalance(ctx, sdk.MustAccAddressFromBech32(participantAddress), k.GetParams(ctx).StagedDenom)
		totalAmounts[participantAddress] = stagedBalance.Amount.Uint64()
	}

	// calculate difference to account balance
	for participantAddress := range totalAmounts {
		totalAmounts[participantAddress] -= currentAmounts[participantAddress]
	}

	return k.convertOrderedClaim(ctx, totalAmounts)
}

// resolveStagedClaims converts staged claims to claims in an ordered fashion for a given range
func (k msgServer) resolveStagedClaims(ctx sdk.Context, start int64, end int64) (err error) {
	popParticipantAmounts, err := k.getClaims(ctx, start, end)
	if err != nil {
		return err
	}

	return k.convertOrderedClaim(ctx, popParticipantAmounts)
}

func (k msgServer) getClaims(ctx sdk.Context, start int64, end int64) (claims map[string]uint64, err error) {
	// lookup all challenges for a given range
	challenges, err := k.GetChallengeRange(ctx, start, end)
	if err != nil {
		return
	}

	claims = make(map[string]uint64)

	for _, challenge := range challenges {
		initiatorAddr, err := sdk.AccAddressFromBech32(challenge.Initiator)
		if err != nil {
			util.GetAppLogger().Error(ctx, "error converting initiator address")
		}
		validatorPopReward, found := k.getChallengeInitiatorReward(ctx, challenge.GetHeight())
		if !found {
			util.GetAppLogger().Error(ctx, "No PoP initiator reward found for height %v", challenge.GetHeight())
		}
		claims[initiatorAddr.String()] += validatorPopReward

		// if challenge not finished only initiator has claims
		if !challenge.GetFinished() {
			continue
		}
		_, challengerAmt, challengeeAmt := util.GetPopReward(challenge.Height, k.GetParams(ctx).PopEpochs)
		claims[challenge.Challenger] += challengerAmt
		if challenge.GetSuccess() {
			claims[challenge.Challengee] += challengeeAmt
		}
	}

	return
}

func (k msgServer) convertOrderedClaim(ctx sdk.Context, claims map[string]uint64) (err error) {
	// second data structure because map iteration order is not guaranteed in GO
	keys := make([]string, 0)
	for accountAddress := range claims {
		keys = append(keys, accountAddress)
	}

	sort.Strings(keys)
	for _, accountAddress := range keys {
		err = k.convertAccountClaim(ctx, accountAddress, claims[accountAddress])
		if err != nil {
			return err
		}
	}

	return
}

// convert per account
func (k msgServer) convertAccountClaim(ctx sdk.Context, participant string, amount uint64) (err error) {
	accAddr, err := sdk.AccAddressFromBech32(participant)
	if err != nil {
		return err
	}

	accStagedClaim := k.bankKeeper.GetBalance(ctx, accAddr, k.GetParams(ctx).StagedDenom)

	if accStagedClaim.Amount.GTE(sdk.NewIntFromUint64(amount)) {
		burnCoins, mintCoins := k.getConvertCoins(ctx, amount)
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, burnCoins)
		if err != nil {
			return err
		}

		err = k.convertCoins(ctx, burnCoins, mintCoins)
		if err != nil {
			return err
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddr, mintCoins)
		if err != nil {
			return err
		}
	}

	return
}

func (k msgServer) convertCoins(ctx sdk.Context, burnCoins sdk.Coins, mintCoins sdk.Coins) (err error) {
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if err != nil {
		return err
	}
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)
}

func (k msgServer) getConvertCoins(ctx sdk.Context, amount uint64) (burnCoins sdk.Coins, mintCoins sdk.Coins) {
	burnCoins = sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).StagedDenom, sdk.NewIntFromUint64(amount)))
	mintCoins = sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).ClaimDenom, sdk.NewIntFromUint64(amount)))
	return
}
