package keeper

import (
	"context"

	"github.com/planetmint/planetmint-go/monitor"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActiveTrustAnchorCount(goCtx context.Context, req *types.QueryActiveTrustAnchorCountRequest) (*types.QueryActiveTrustAnchorCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	count := monitor.GetActiveActorCount()

	return &types.QueryActiveTrustAnchorCountResponse{Count: count}, nil
}
