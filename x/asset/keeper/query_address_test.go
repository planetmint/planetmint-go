package keeper_test

import (
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/asset/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetNotarizedAssetByAddress(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	_ = createNAsset(keeper, ctx, 10)
	assets, _ := keeper.GetAssetsByAddress(ctx, "plmnt_address", nil, util.SerializeUint64(3+1))
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCIDsByAddressRequest
		response *types.QueryGetCIDsByAddressResponse
		err      error
	}{
		{
			desc:     "cid found",
			request:  &types.QueryGetCIDsByAddressRequest{Address: "plmnt_address", NumElements: 3},
			response: &types.QueryGetCIDsByAddressResponse{Cids: assets},
		},
		{
			desc:    "cid not found",
			request: &types.QueryGetCIDsByAddressRequest{Address: "invalid key"},
			err:     status.Error(codes.NotFound, "no CIDs found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetCIDsByAddress(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
