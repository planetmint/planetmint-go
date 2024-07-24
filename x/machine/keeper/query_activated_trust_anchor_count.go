package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActivatedTrustAnchorCount(goCtx context.Context, req *types.QueryActivatedTrustAnchorCountRequest) (*types.QueryActivatedTrustAnchorCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	count := k.GetActivatedTACount(ctx)

	return &types.QueryActivatedTrustAnchorCountResponse{Count: count}, nil
}
