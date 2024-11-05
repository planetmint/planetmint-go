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

	amount := k.GetValidatorPoPReward(ctx)
	k.StoreChallangeInitiatorReward(ctx, msg.GetHeight(), amount)

	if util.IsValidatorBlockProposer(ctx, k.RootDir) {
		go util.SendMqttPopInitMessagesToServer(ctx, challenge)
	}

	return &types.MsgInitPopResponse{}, nil
}
