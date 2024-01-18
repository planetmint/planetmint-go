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
	// ensure the challenge is stored even without the token reward minting
	k.StoreChallenge(ctx, *msg.Challenge)

	if msg.Challenge.GetSuccess() {
		util.GetAppLogger().Info(ctx, "PoP at height %v was successful", msg.Challenge.GetHeight())
	} else {
		util.GetAppLogger().Info(ctx, "PoP at height %v was unsuccessful", msg.Challenge.GetHeight())
	}

	err = k.issuePoPRewards(ctx, *msg.Challenge)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrFailedPoPRewardsIssuance, err.Error())
	}

	return &types.MsgReportPopResultResponse{}, nil
}

func (k msgServer) issuePoPRewards(ctx sdk.Context, challenge types.Challenge) (err error) {
	conf := config.GetConfig()
	total, challengerAmt, _ := util.GetPopReward(challenge.Height)

	stagedCRDDL := sdk.NewCoin(conf.StagedDenom, sdk.ZeroInt())
	if challenge.GetSuccess() {
		stagedCRDDL = stagedCRDDL.AddAmount(sdk.NewIntFromUint64(total))
	} else {
		stagedCRDDL = stagedCRDDL.AddAmount(sdk.NewIntFromUint64(challengerAmt))
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(stagedCRDDL))
	if err != nil {
		return
	}

	return k.handlePoP(ctx, challenge)
}

func (k msgServer) handlePoP(ctx sdk.Context, challenge types.Challenge) (err error) {
	_, challengerAmt, challengeeAmt := util.GetPopReward(challenge.Height)

	err = k.sendRewards(ctx, challenge.GetChallenger(), challengerAmt)
	if err != nil {
		return
	}

	if !challenge.GetSuccess() {
		return
	}

	return k.sendRewards(ctx, challenge.GetChallengee(), challengeeAmt)
}

func (k msgServer) sendRewards(ctx sdk.Context, receiver string, amt uint64) (err error) {
	conf := config.GetConfig()
	coins := sdk.NewCoins(sdk.NewCoin(conf.StagedDenom, sdk.NewIntFromUint64(amt)))
	receiverAddr, err := sdk.AccAddressFromBech32(receiver)
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, coins)
}
