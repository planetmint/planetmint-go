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

func (k Keeper) RedeemClaimAll(goCtx context.Context, req *types.QueryAllRedeemClaimRequest) (*types.QueryAllRedeemClaimResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	var redeemClaims []types.RedeemClaim
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	redeemClaimStore := prefix.NewStore(store, types.KeyPrefix(types.RedeemClaimKeyPrefix))

	// revive linter: from func(key []byte, value []byte)  to func(_ []byte, value []byte)
	pageRes, err := query.Paginate(redeemClaimStore, req.Pagination, func(_ []byte, value []byte) error {
		var redeemClaim types.RedeemClaim
		if err := k.cdc.Unmarshal(value, &redeemClaim); err != nil {
			return err
		}

		redeemClaims = append(redeemClaims, redeemClaim)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRedeemClaimResponse{RedeemClaim: redeemClaims, Pagination: pageRes}, nil
}

func (k Keeper) RedeemClaim(goCtx context.Context, req *types.QueryGetRedeemClaimRequest) (*types.QueryGetRedeemClaimResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetRedeemClaim(
		ctx,
		req.Beneficiary,
		req.Id,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRedeemClaimResponse{RedeemClaim: val}, nil
}

func (k Keeper) RedeemClaimByLiquidTxHash(goCtx context.Context, req *types.QueryRedeemClaimByLiquidTxHashRequest) (*types.QueryRedeemClaimByLiquidTxHashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetRedeemClaimByLiquidTXHash(
		ctx,
		req.LiquidTxHash,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryRedeemClaimByLiquidTxHashResponse{RedeemClaim: &val}, nil
}
