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

	// TODO: develop a more resilient pattern: if the distribution does not work,
	//       the challenge shouldn't be discarded. it's most likely not the fault of the PoP participants.
	err = k.issuePoPRewards(ctx, *msg.Challenge)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrFailedPoPRewardsIssuance, err.Error())
	}

	k.StoreChallenge(ctx, *msg.Challenge)

	return &types.MsgReportPopResultResponse{}, nil
}

func (k msgServer) issuePoPRewards(ctx sdk.Context, challenge types.Challenge) (err error) {
	cfg := config.GetConfig()
	amt := GetReissuanceAsStringValue(challenge.GetHeight())
	amtFloat, err := util.RDDLTokenStringToFloat(amt)
	if err != nil {
		return err
	}

	popAmt := uint64(amtFloat * types.PercentagePop)
	stagedCRDDL := sdk.NewCoin(cfg.StagedDenom, sdk.NewIntFromUint64(popAmt))

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(stagedCRDDL))
	if err != nil {
		return err
	}

	if challenge.Success {
		err = k.handlePoPSuccess(ctx, challenge, amtFloat)
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

func (k msgServer) handlePoPSuccess(ctx sdk.Context, challenge types.Challenge, amount float64) (err error) {
	cfg := config.GetConfig()
	challengerAmt := uint64(amount * types.PercentageChallenger)
	challengeeAmt := uint64(amount * types.PercentageChallengee)

	challengerCoin := sdk.NewCoin(cfg.StagedDenom, sdk.NewIntFromUint64(challengerAmt))
	challengeeCoin := sdk.NewCoin(cfg.StagedDenom, sdk.NewIntFromUint64(challengeeAmt))
	challengee, err := sdk.AccAddressFromBech32(challenge.Challengee)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, challengee, sdk.NewCoins(challengeeCoin))
	if err != nil {
		return err
	}
	challenger, err := sdk.AccAddressFromBech32(challenge.Challenger)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, challenger, sdk.NewCoins(challengerCoin))
	if err != nil {
		return err
	}
	return
}
