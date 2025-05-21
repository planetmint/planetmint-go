package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/der/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Nft(goCtx context.Context, req *types.QueryNftRequest) (*types.QueryNftResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	derNft, found := k.LookupLiquidDerAsset(ctx, req.ZigbeeID)
	if !found {
		return nil, status.Error(codes.NotFound, "error zigbeeID not found: "+req.ZigbeeID)
	}

	return &types.QueryNftResponse{DerNft: &derNft}, nil
}
