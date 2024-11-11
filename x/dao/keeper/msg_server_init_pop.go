package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) InitPop(goCtx context.Context, msg *types.MsgInitPop) (*types.MsgInitPopResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var challenge types.Challenge
	challenge.Initiator = msg.GetInitiator()
	challenge.Challengee = msg.GetChallengee()
	challenge.Challenger = msg.GetChallenger()
	challenge.Height = msg.GetHeight()

	k.StoreChallenge(ctx, challenge)

	if util.IsValidatorBlockProposer(ctx, k.RootDir) {
		go util.SendMqttPopInitMessagesToServer(ctx, challenge)
	}

	amount := k.GetValidatorPoPReward(ctx)
	k.StoreChallangeInitiatorReward(ctx, msg.GetHeight(), amount)

	// TODO: expand err value in log
	initiatorAddr, err := sdk.AccAddressFromBech32(msg.GetInitiator())
	if err != nil {
		util.GetAppLogger().Error(ctx, "error converting initiator address: %v", err)
	}

	valReward := sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).StagedDenom, sdk.NewIntFromUint64(amount)))
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, valReward)
	if err != nil {
		util.GetAppLogger().Error(ctx, "error minting initiator rewards: %v", err)
	}

	if err := k.sendRewards(ctx, initiatorAddr.String(), amount); err != nil {
		util.GetAppLogger().Error(ctx, "failed to send rewards: %v", err)
	}

	return &types.MsgInitPopResponse{}, nil
}
