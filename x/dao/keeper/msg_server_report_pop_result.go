package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReportPopResult(goCtx context.Context, msg *types.MsgReportPopResult) (*types.MsgReportPopResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := util.ValidateStruct(*msg.Challenge)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidChallenge, err.Error())
	}

	err = k.issuePoPRewards(ctx, *msg.Challenge)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrFailedPoPRewardsIssuance, err.Error())
	}

	k.StoreChallenge(ctx, *msg.Challenge)

	return &types.MsgReportPopResultResponse{}, nil
}

func (k msgServer) issuePoPRewards(ctx sdk.Context, challenge types.Challenge) (err error) {
	cfg := config.GetConfig()
	amt := GetReissuanceAmount()
	amtUint, err := util.RDDLTokenStringToFloat(amt)
	if err != nil {
		return err
	}

	popAmt := uint64(amtUint * types.PercentagePop)
	stagedCRDDL := sdk.NewCoin(cfg.StagedDenom, sdk.NewIntFromUint64(popAmt))

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(stagedCRDDL))
	if err != nil {
		return err
	}

	challengerAmt := uint64(amtUint * types.PercentageChallenger)
	challengeeAmt := uint64(amtUint * types.PercentageChallengee)

	challengerCoin := sdk.NewCoin(cfg.StagedDenom, sdk.NewIntFromUint64(challengerAmt))
	challengeeCoin := sdk.NewCoin(cfg.StagedDenom, sdk.NewIntFromUint64(challengeeAmt))
	if challenge.Success {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(challenge.Challengee), sdk.NewCoins(challengeeCoin))
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(challenge.Challenger), sdk.NewCoins(challengerCoin))
		if err != nil {
			return err
		}
	} else {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(challenge.Challenger), sdk.NewCoins(stagedCRDDL))
		if err != nil {
			return err
		}
	}

	return err
}
