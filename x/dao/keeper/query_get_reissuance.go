package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetReissuance(goCtx context.Context, req *types.QueryGetReissuanceRequest) (*types.QueryGetReissuanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	reissuance, found := k.LookupReissuance(ctx, req.GetBlockHeight())
	if !found {
		return nil, status.Error(codes.NotFound, "reissuance not found")
	}

	return &types.QueryGetReissuanceResponse{Reissuance: &reissuance}, nil
}
