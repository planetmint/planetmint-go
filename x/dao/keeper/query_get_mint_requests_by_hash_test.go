package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestQueryGetMintRequestByHash(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	items := createNMintRequests(keeper, ctx, sample.ConstBech32Addr, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMintRequestsByHashRequest
		response *types.QueryGetMintRequestsByHashResponse
		err      error
	}{
		{
			desc:     "mint request found",
			request:  &types.QueryGetMintRequestsByHashRequest{Hash: "hash0"},
			response: &types.QueryGetMintRequestsByHashResponse{MintRequest: &items[0]},
		},
		{
			desc:    "mint request not found",
			request: &types.QueryGetMintRequestsByHashRequest{Hash: "invalid hash"},
			err:     status.Error(codes.NotFound, "mint request not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			res, err := keeper.GetMintRequestsByHash(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, res)
			}
		})
	}
}
