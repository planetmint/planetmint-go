package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetTrustAnchorStatus(goCtx context.Context, req *types.QueryGetTrustAnchorStatusRequest) (*types.QueryGetTrustAnchorStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	_, activated, found := k.GetTrustAnchor(ctx, req.MachineId)
	if !found {
		return nil, status.Error(codes.NotFound, "trust anchor not found")
	}

	return &types.QueryGetTrustAnchorStatusResponse{MachineId: req.MachineId, IsActivated: activated}, nil
}
