package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetLiquidAssetsByMachineid(goCtx context.Context, req *types.QueryGetLiquidAssetsByMachineidRequest) (*types.QueryGetLiquidAssetsByMachineidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	liquidAsset, found := k.LookupLiquidAsset(ctx, req.GetMachineID())
	if !found {
		return nil, status.Error(codes.InvalidArgument, "no associated asset found")
	}
	return &types.QueryGetLiquidAssetsByMachineidResponse{LiquidAssetEntry: &liquidAsset}, nil
}
