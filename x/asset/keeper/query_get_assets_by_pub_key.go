package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/asset/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetCIDsByPubKey(goCtx context.Context, req *types.QueryGetCIDsByPubKeyRequest) (*types.QueryGetCIDsByPubKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	cids, found := k.GetCidsByPublicKey(ctx, req.GetExtPubKey())
	if !found {
		return nil, status.Error(codes.NotFound, "no CIDs found")
	}

	return &types.QueryGetCIDsByPubKeyResponse{CIDs: cids}, nil
}
