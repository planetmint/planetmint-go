package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	return &types.MsgInitPopResponse{}, nil
}