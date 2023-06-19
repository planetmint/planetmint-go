package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"planetmint-go/x/machine/types"
)

func (k msgServer) AttestMachine(goCtx context.Context, msg *types.MsgAttestMachine) (*types.MsgAttestMachineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAttestMachineResponse{}, nil
}
