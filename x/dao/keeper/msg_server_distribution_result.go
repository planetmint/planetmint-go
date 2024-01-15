package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) DistributionResult(goCtx context.Context, msg *types.MsgDistributionResult) (*types.MsgDistributionResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	distribution, found := k.LookupDistributionOrder(ctx, msg.GetLastPop())
	if found {
		distribution.DaoTxID = msg.DaoTxID
		distribution.PopTxID = msg.PopTxID
		distribution.InvestorTxID = msg.InvestorTxID
		err := k.resolveStagedClaims(ctx, distribution.FirstPop, distribution.LastPop)
		if err != nil {
			util.GetAppLogger().Error(ctx, "%s for provided PoP heights: %d %d", types.ErrResolvingStagedClaims.Error(), distribution.FirstPop, distribution.LastPop)
		} else {
			util.GetAppLogger().Info(ctx, "staged claims successfully for provided PoP heights: %d %d", distribution.FirstPop, distribution.LastPop)
		}
		k.StoreDistributionOrder(ctx, distribution)
	} else {
		util.GetAppLogger().Error(ctx, "%s for provided block height %s", types.ErrDistributionNotFound.Error(), strconv.FormatInt(msg.GetLastPop(), 10))
	}

	return &types.MsgDistributionResultResponse{}, nil
}

func (k msgServer) resolveStagedClaims(ctx sdk.Context, start int64, end int64) (err error) {
	// lookup all challenges since the last distribution
	challenges, err := k.GetChallengeRange(ctx, start, end)
	if err != nil {
		return err
	}

	popParticipants := make(map[string]uint64)
	daoModuleAmt := uint64(0)

	for _, challenge := range challenges {
		_, challengerAmt, challengeeAmt := util.GetPopReward(challenge.Height)
		popParticipants[challenge.Challenger] += challengerAmt
		if challenge.GetSuccess() {
			popParticipants[challenge.Challengee] += challengeeAmt
		} else {
			daoModuleAmt += challengeeAmt
		}
	}

	// second data structure because map iteration order is not guaranteed in GO
	keys := make([]string, 0)
	for p := range popParticipants {
		keys = append(keys, p)
	}
	for _, p := range keys {
		err = k.convertAccountClaim(ctx, p, popParticipants[p])
		if err != nil {
			return err
		}
	}

	// convert dao module claims
	burnCoins, mintCoins := getConvertCoins(daoModuleAmt)
	err = k.convertCoins(ctx, burnCoins, mintCoins)

	return
}

// convert per account
func (k msgServer) convertAccountClaim(ctx sdk.Context, participant string, amount uint64) (err error) {
	conf := config.GetConfig()
	accAddr, err := sdk.AccAddressFromBech32(participant)
	if err != nil {
		return err
	}

	accStagedClaim := k.bankKeeper.GetBalance(ctx, accAddr, conf.StagedDenom)

	if accStagedClaim.Amount.GTE(sdk.NewIntFromUint64(amount)) {
		burnCoins, mintCoins := getConvertCoins(amount)
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
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)
	if err != nil {
		return err
	}
	return
}

func getConvertCoins(amount uint64) (burnCoins sdk.Coins, mintCoins sdk.Coins) {
	conf := config.GetConfig()
	burnCoins = sdk.NewCoins(sdk.NewCoin(conf.StagedDenom, sdk.NewIntFromUint64(amount)))
	mintCoins = sdk.NewCoins(sdk.NewCoin(conf.ClaimDenom, sdk.NewIntFromUint64(amount)))
	return
}
