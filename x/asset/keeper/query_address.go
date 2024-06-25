package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/asset/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetCIDsByAddress(goCtx context.Context, req *types.QueryGetCIDsByAddressRequest) (*types.QueryGetCIDsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	cids, found := k.GetAssetsByAddress(ctx, req.GetAddress(), nil, util.SerializeUint64(req.GetNumElements()))
	if !found {
		return nil, status.Error(codes.NotFound, "no CIDs found")
	}

	return &types.QueryGetCIDsByAddressResponse{Cids: cids}, nil
}
