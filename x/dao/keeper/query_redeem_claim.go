package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RedeemClaimAll(goCtx context.Context, req *types.QueryAllRedeemClaimRequest) (*types.QueryAllRedeemClaimResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var redeemClaims []types.RedeemClaim
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	redeemClaimStore := prefix.NewStore(store, types.KeyPrefix(types.RedeemClaimKeyPrefix))

	pageRes, err := query.Paginate(redeemClaimStore, req.Pagination, func(key []byte, value []byte) error {
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
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetRedeemClaim(
		ctx,
		req.Beneficiary,
		req.LiquidTxHash,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRedeemClaimResponse{RedeemClaim: val}, nil
}
