package keeper

import (
	"context"

	"planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AttestMachine(goCtx context.Context, msg *types.MsgAttestMachine) (*types.MsgAttestMachineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.StoreMachine(ctx, *msg.Machine)
	k.StoreMachineIndex(ctx, *msg.Machine)

	return &types.MsgAttestMachineResponse{}, nil
}