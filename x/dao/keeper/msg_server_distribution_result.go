package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) DistributionResult(goCtx context.Context, msg *types.MsgDistributionResult) (*types.MsgDistributionResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	distribution, found := k.LookupDistributionOrder(ctx, msg.GetLastPop())
	if found {
		distribution.DaoTxID = msg.DaoTxID
		distribution.PopTxID = msg.PopTxID
		distribution.InvestorTxID = msg.InvestorTxID
		err := k.resolveStagedClaims(ctx, uint64(distribution.FirstPop), uint64(distribution.LastPop))
		if err != nil {
			return nil, errorsmod.Wrapf(types.ErrResolvingStagedClaims, " for provieded PoP heights: %d %d", distribution.FirstPop, distribution.LastPop)
		}
		k.StoreDistributionOrder(ctx, distribution)
	} else {
		return nil, errorsmod.Wrapf(types.ErrDistributionNotFound, " for provided block height %s", strconv.FormatInt(msg.GetLastPop(), 10))
	}

	return &types.MsgDistributionResultResponse{}, nil
}

func (k msgServer) resolveStagedClaims(ctx sdk.Context, start uint64, end uint64) (err error) {
	// lookup all challenges since the last distribution
	challenges, err := k.GetChallengeRange(ctx, start, end)
	if err != nil {
		return err
	}

	for _, challenge := range challenges {
		err = k.convertClaim(ctx, challenge.Challengee)
		if err != nil {
			return err
		}
		err = k.convertClaim(ctx, challenge.Challenger)
		if err != nil {
			return err
		}
	}

	return
}

func (k msgServer) convertClaim(ctx sdk.Context, addr string) (err error) {
	accAddress, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err
	}

	stagedClaim := k.bankKeeper.GetBalance(ctx, accAddress, "stagedCRDDL")

	if stagedClaim.Amount.GT(math.ZeroInt()) {
		claim := sdk.NewCoins(sdk.NewCoin("cRDDL", stagedClaim.Amount))
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddress, types.ModuleName, sdk.NewCoins(stagedClaim))
		if err != nil {
			return err
		}
		err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(stagedClaim))
		if err != nil {
			return err
		}
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, claim)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddress, claim)
		if err != nil {
			return err
		}
	}

	return
}
