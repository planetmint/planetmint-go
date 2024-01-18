package keeper_test

import (
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/asset/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetNotarizedAsset(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAsset(keeper, ctx, 1)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNotarizedAssetRequest
		response *types.QueryGetNotarizedAssetResponse
		err      error
	}{
		{
			desc:     "cid found",
			request:  &types.QueryGetNotarizedAssetRequest{Cid: msgs[0].GetCid()},
			response: &types.QueryGetNotarizedAssetResponse{Cid: msgs[0].GetCid()},
		},
		{
			desc:    "cid not found",
			request: &types.QueryGetNotarizedAssetRequest{Cid: "invalid key"},
			err:     status.Error(codes.NotFound, "cid not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetNotarizedAsset(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
