package keeper

import (
	"context"

	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetMachineByPublicKey(goCtx context.Context, req *types.QueryGetMachineByPublicKeyRequest) (*types.QueryGetMachineByPublicKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	machineIndex, found := k.GetMachineIndexByPubKey(ctx, req.PublicKey)
	if !found {
		return nil, status.Error(codes.NotFound, "machine not found by public key: "+req.PublicKey)
	}

	machine, found := k.GetMachine(ctx, machineIndex)
	if !found {
		return nil, status.Error(codes.Internal, "error while fetching machine")
	}

	return &types.QueryGetMachineByPublicKeyResponse{Machine: &machine}, nil
}
