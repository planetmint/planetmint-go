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

func TestQueryMintRequestByAddress(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	items := createNMintRequests(keeper, ctx, sample.ConstBech32Addr, 10)

	var mintRequests types.MintRequests
	for i := range items {
		mintRequests.Requests = append(mintRequests.Requests, &items[i])
	}

	for _, tc := range []struct {
		desc     string
		request  *types.QueryMintRequestsByAddressRequest
		response *types.QueryMintRequestsByAddressResponse
		err      error
	}{
		{
			desc:     "mint requests found",
			request:  &types.QueryMintRequestsByAddressRequest{Address: sample.ConstBech32Addr},
			response: &types.QueryMintRequestsByAddressResponse{MintRequests: &mintRequests},
		},
		{
			desc:    "mint requests not found",
			request: &types.QueryMintRequestsByAddressRequest{Address: "invalid hash"},
			err:     status.Error(codes.NotFound, "mint requests not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			res, err := keeper.MintRequestsByAddress(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, res)
			}
		})
	}
}
