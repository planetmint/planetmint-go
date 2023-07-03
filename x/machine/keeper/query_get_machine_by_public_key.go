package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"planetmint-go/x/machine/types"
)

func (k Keeper) GetMachineByPublicKey(goCtx context.Context, req *types.QueryGetMachineByPublicKeyRequest) (*types.QueryGetMachineByPublicKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetMachineByPublicKeyResponse{}, nil
}
