package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MintRequestsByAddress(goCtx context.Context, req *types.QueryMintRequestsByAddressRequest) (*types.QueryMintRequestsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	mintRequests, found := k.GetMintRequestsByAddress(ctx, req.GetAddress())
	if !found {
		return nil, status.Error(codes.NotFound, "mint requests not found")
	}

	return &types.QueryMintRequestsByAddressResponse{MintRequests: &mintRequests}, nil
}
