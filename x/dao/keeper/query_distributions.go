package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Distributions(goCtx context.Context, req *types.QueryDistributionsRequest) (*types.QueryDistributionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	distributions := make([]types.DistributionOrder, 0)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
	pageRes, err := query.Paginate(store, req.Pagination, func(_ []byte, value []byte) (err error) {
		var distribution types.DistributionOrder
		err = distribution.Unmarshal(value)
		distributions = append(distributions, distribution)
		return
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryDistributionsResponse{Distributions: distributions, Pagination: pageRes}, nil
}
