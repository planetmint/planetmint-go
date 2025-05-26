package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/der/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Der(goCtx context.Context, req *types.QueryDerRequest) (*types.QueryDerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	der, found := k.LookupDerAsset(ctx, req.ZigbeeID)
	if !found {
		return nil, status.Error(codes.NotFound, "error zigbeeID not found: "+req.ZigbeeID)
	}

	return &types.QueryDerResponse{Der: &der}, nil
}
