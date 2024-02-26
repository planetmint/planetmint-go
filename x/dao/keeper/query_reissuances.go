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

func (k Keeper) Reissuances(goCtx context.Context, req *types.QueryReissuancesRequest) (*types.QueryReissuancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	reissuances := make([]types.Reissuance, 0)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	pageRes, err := query.Paginate(store, req.Pagination, func(_ []byte, value []byte) (err error) {
		var reissuance types.Reissuance
		err = reissuance.Unmarshal(value)
		reissuances = append(reissuances, reissuance)
		return
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryReissuancesResponse{Reissuances: reissuances, Pagination: pageRes}, nil
}
