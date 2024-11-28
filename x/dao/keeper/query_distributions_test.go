package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/require"
)

func TestQueryDistribtions(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	//var popEpochs int64 = 24
	distributions := createNDistributionOrder(keeper, ctx, 20)
	for _, tc := range []struct {
		desc                  string
		request               *types.QueryDistributionsRequest
		responseDistributions []types.DistributionOrder
		err                   error
	}{
		{
			desc: "query 0 offset 10 limit",
			request: &types.QueryDistributionsRequest{
				Pagination: &query.PageRequest{
					Key:        nil,
					Offset:     0,
					Limit:      10,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseDistributions: distributions[:10],
		},
		{
			desc: "query 1 offset 5 limit",
			request: &types.QueryDistributionsRequest{
				Pagination: &query.PageRequest{
					Key:        nil,
					Offset:     1,
					Limit:      5,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseDistributions: distributions[1:6],
		},
		{
			desc: "query 5*1000 key 0 offset 10 limit",
			request: &types.QueryDistributionsRequest{
				Pagination: &query.PageRequest{
					Key:        util.SerializeInt64(5 * 1000),
					Offset:     0,
					Limit:      10,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseDistributions: distributions[4:14],
		},
		{
			desc: "query 2*1000 key 0 offset 10 limit",
			request: &types.QueryDistributionsRequest{
				Pagination: &query.PageRequest{
					Key:        util.SerializeInt64(2 * 1000),
					Offset:     0,
					Limit:      10,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseDistributions: distributions[1:11],
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := keeper.Distributions(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, tc.err, err)
			} else {
				require.Equal(t, tc.responseDistributions, res.Distributions)
			}
		})
	}
}
