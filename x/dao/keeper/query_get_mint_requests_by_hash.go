package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetMintRequestsByHash(goCtx context.Context, req *types.QueryGetMintRequestsByHashRequest) (*types.QueryGetMintRequestsByHashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	mintRequest, found := k.GetMintRequestByHash(ctx, req.GetHash())
	if !found {
		return nil, status.Error(codes.NotFound, "mint request not found")
	}

	return &types.QueryGetMintRequestsByHashResponse{MintRequest: &mintRequest}, nil
}
