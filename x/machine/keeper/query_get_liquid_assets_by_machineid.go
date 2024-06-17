package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetLiquidAssetsByMachineId(goCtx context.Context, req *types.QueryGetLiquidAssetsByMachineIdRequest) (*types.QueryGetLiquidAssetsByMachineIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	liquidAsset, found := k.LookupLiquidAsset(ctx, req.GetMachineId())
	if !found {
		return nil, status.Error(codes.InvalidArgument, "no associated asset found")
	}
	return &types.QueryGetLiquidAssetsByMachineIdResponse{LiquidAssetEntry: &liquidAsset}, nil
}
