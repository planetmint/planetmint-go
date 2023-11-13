package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetReissuances(goCtx context.Context, req *types.QueryGetReissuancesRequest) (*types.QueryGetReissuancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	reissuances := k.getReissuancesPage(ctx, req.Pagination.GetKey(),
		req.Pagination.GetOffset(), req.Pagination.GetLimit(),
		req.Pagination.GetCountTotal(), req.Pagination.GetReverse())

	if reissuances != nil {
		return &types.QueryGetReissuancesResponse{Reissuance: &reissuances[0]}, nil
	}
	return &types.QueryGetReissuancesResponse{}, nil

}
