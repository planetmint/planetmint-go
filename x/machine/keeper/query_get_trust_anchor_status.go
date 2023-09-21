package keeper

import (
	"context"

	"planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetTrustAnchorStatus(goCtx context.Context, req *types.QueryGetTrustAnchorStatusRequest) (*types.QueryGetTrustAnchorStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	_, activated, found := k.GetTrustAnchor(ctx, req.Machineid)
	if !found {
		return nil, status.Error(codes.NotFound, "trust anchor not found")
	}

	return &types.QueryGetTrustAnchorStatusResponse{Machineid: req.Machineid, Isactivated: activated}, nil
}
