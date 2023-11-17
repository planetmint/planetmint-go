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

func TestGetNotarizedAssetByAddress(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.AssetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	_ = createNAsset(keeper, ctx, 10)
	assets, _ := keeper.GetCidsByAddress(ctx, "plmnt_address")
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCIDsByAddressRequest
		response *types.QueryGetCIDsByAddressResponse
		err      error
	}{
		{
			desc:     "cid found",
			request:  &types.QueryGetCIDsByAddressRequest{Address: "plmnt_address"},
			response: &types.QueryGetCIDsByAddressResponse{Cids: assets},
		},
		{
			desc:    "cid not found",
			request: &types.QueryGetCIDsByAddressRequest{Address: "invalid key"},
			err:     status.Error(codes.NotFound, "no CIDs found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			response, err := keeper.GetCIDsByAddress(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
