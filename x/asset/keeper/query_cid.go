package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/x/asset/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetNotarizedAsset(goCtx context.Context, req *types.QueryGetNotarizedAssetRequest) (*types.QueryGetNotarizedAssetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	asset, found := k.GetAsset(ctx, req.GetCid())
	if !found {
		return nil, status.Error(codes.NotFound, "cid not found")
	}

	return &types.QueryGetNotarizedAssetResponse{Cid: asset.GetCid()}, nil
}
