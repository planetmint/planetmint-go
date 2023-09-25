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

func TestGetNotarizedAssetByPubKey(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	_ = createNAsset(keeper, ctx, 10)
	assets, _ := keeper.GetCidsByPublicKey(ctx, "pubkey_search")
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCIDsByPubKeyRequest
		response *types.QueryGetCIDsByPubKeyResponse
		err      error
	}{
		{
			desc:     "cid found",
			request:  &types.QueryGetCIDsByPubKeyRequest{ExtPubKey: "pubkey_search"},
			response: &types.QueryGetCIDsByPubKeyResponse{CIDs: assets},
		},
		{
			desc:    "cid not found",
			request: &types.QueryGetCIDsByPubKeyRequest{ExtPubKey: "invalid key"},
			err:     status.Error(codes.NotFound, "no CIDs found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetCIDsByPubKey(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
