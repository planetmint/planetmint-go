package keeper

import (
	"context"

	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetTrustAnchorStatus(goCtx context.Context, req *types.QueryGetTrustAnchorStatusRequest) (*types.QueryGetTrustAnchorStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	_, activated, found := k.GetTrustAnchor(ctx, req.Machineid)
	if !found {
		return nil, status.Error(codes.NotFound, "trust anchor not found by machine id: "+req.Machineid)
	}

	return &types.QueryGetTrustAnchorStatusResponse{Machineid: req.Machineid, Isactivated: activated}, nil
}
