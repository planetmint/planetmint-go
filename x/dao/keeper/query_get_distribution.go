package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetDistribution(goCtx context.Context, req *types.QueryGetDistributionRequest) (*types.QueryGetDistributionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	distribution, found := k.LookupDistributionOrder(ctx, req.GetHeight())
	if !found {
		return nil, status.Error(codes.NotFound, "distribution not found")
	}

	return &types.QueryGetDistributionResponse{Distribution: &distribution}, nil
}
