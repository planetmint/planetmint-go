package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) updateChallenge(ctx sdk.Context, msg *types.MsgReportPopResult) (err error) {
	challenge, found := k.LookupChallenge(ctx, msg.GetChallenge().GetHeight())
	if !found {
		err = errorsmod.Wrapf(types.ErrInvalidChallenge, "no challenge found for PoP report")
		return
	}
	if challenge.Challengee != msg.GetChallenge().Challengee ||
		challenge.Challenger != msg.GetChallenge().Challenger ||
		challenge.Initiator != msg.GetChallenge().Initiator ||
		challenge.Height != msg.GetChallenge().Height {
		err = errorsmod.Wrapf(types.ErrInvalidChallenge, "PoP report data does not match challenge")
		return
	}
	challenge.Success = msg.GetChallenge().GetSuccess()
	challenge.Finished = true

	k.StoreChallenge(ctx, challenge)
	return
}

func (k msgServer) ReportPopResult(goCtx context.Context, msg *types.MsgReportPopResult) (*types.MsgReportPopResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := util.ValidateStruct(*msg.Challenge)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidChallenge, err.Error())
	}

	// verify that the report origin is the challenger
	if msg.GetCreator() != msg.GetChallenge().GetChallenger() {
		err = errorsmod.Wrapf(types.ErrInvalidPopReporter, "PoP reporter is not the challenger")
		return nil, err
	}

	_, err = sdk.AccAddressFromBech32(msg.Challenge.GetInitiator())
	if err != nil {
		util.GetAppLogger().Error(ctx, "error converting initiator address")
		return nil, errorsmod.Wrap(types.ErrInvalidPoPInitiator, "PoP initiator not hex encoded")
	}

	// update valid PoP Result reports
	err = k.updateChallenge(ctx, msg)
	if err != nil {
		return nil, err
	}

	if msg.Challenge.GetSuccess() {
		util.GetAppLogger().Info(ctx, "PoP at height %v was successful", msg.Challenge.GetHeight())
	} else {
		util.GetAppLogger().Info(ctx, "PoP at height %v was unsuccessful", msg.Challenge.GetHeight())
	}

	err = k.issuePoPRewards(ctx, *msg.Challenge)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrFailedPoPRewardsIssuance, err.Error())
	}

	return &types.MsgReportPopResultResponse{}, nil
}

func (k msgServer) issuePoPRewards(ctx sdk.Context, challenge types.Challenge) (err error) {
	total, challengerAmt, _ := util.GetPopReward(challenge.Height, k.GetParams(ctx).PopEpochs)

	stagedCRDDL := sdk.NewCoin(k.GetParams(ctx).StagedDenom, sdk.ZeroInt())
	if challenge.GetSuccess() {
		stagedCRDDL = stagedCRDDL.AddAmount(sdk.NewIntFromUint64(total))
	} else {
		stagedCRDDL = stagedCRDDL.AddAmount(sdk.NewIntFromUint64(challengerAmt))
	}

	validatorPoPreward, found := k.getChallengeInitiatorReward(ctx, challenge.GetHeight())
	if !found {
		util.GetAppLogger().Error(ctx, "No PoP initiator reward found for height %v", challenge.GetHeight())
	}
	stagedCRDDL = stagedCRDDL.AddAmount(sdk.NewIntFromUint64(validatorPoPreward))

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(stagedCRDDL))
	if err != nil {
		return
	}

	return k.handlePoP(ctx, challenge)
}

func (k msgServer) handlePoP(ctx sdk.Context, challenge types.Challenge) (err error) {
	_, challengerAmt, challengeeAmt := util.GetPopReward(challenge.Height, k.GetParams(ctx).PopEpochs)

	err = k.sendRewards(ctx, challenge.GetChallenger(), challengerAmt)
	if err != nil {
		return
	}

	initiatorAddr, _ := sdk.AccAddressFromBech32(challenge.Initiator)
	err = k.sendRewards(ctx, initiatorAddr.String(), k.GetValidatorPoPReward(ctx))
	if err != nil {
		return
	}

	if !challenge.GetSuccess() {
		return
	}

	return k.sendRewards(ctx, challenge.GetChallengee(), challengeeAmt)
}

func (k msgServer) sendRewards(ctx sdk.Context, receiver string, amt uint64) (err error) {
	coins := sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).StagedDenom, sdk.NewIntFromUint64(amt)))
	receiverAddr, err := sdk.AccAddressFromBech32(receiver)
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, coins)
}
