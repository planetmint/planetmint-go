package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestQueryGetReissuance(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	items := createNReissuances(keeper, ctx, 1, types.DefaultGenesis().GetParams().PopEpochs)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetReissuanceRequest
		response *types.QueryGetReissuanceResponse
		err      error
	}{
		{
			desc:     "reissuance request found",
			request:  &types.QueryGetReissuanceRequest{BlockHeight: 0},
			response: &types.QueryGetReissuanceResponse{Reissuance: &items[0]},
		},
		{
			desc:    "reissuance request not found",
			request: &types.QueryGetReissuanceRequest{BlockHeight: 100},
			err:     status.Error(codes.NotFound, "reissuance not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := keeper.GetReissuance(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, res)
			}
		})
	}
}
