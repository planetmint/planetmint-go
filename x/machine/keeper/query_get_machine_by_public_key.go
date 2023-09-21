package keeper

import (
	"context"

	"github.com/planetmint/planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetMachineByPublicKey(goCtx context.Context, req *types.QueryGetMachineByPublicKeyRequest) (*types.QueryGetMachineByPublicKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	machineIndex, found := k.GetMachineIndex(ctx, req.PublicKey)
	if !found {
		return nil, status.Error(codes.NotFound, "machine not found")
	}

	machine, _ := k.GetMachine(ctx, machineIndex)

	return &types.QueryGetMachineByPublicKeyResponse{Machine: &machine}, nil
}
