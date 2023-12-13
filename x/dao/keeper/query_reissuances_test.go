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

func TestQueryReissuances(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	reissuances := createNReissuances(keeper, ctx, 20)

	for _, tc := range []struct {
		desc               string
		request            *types.QueryReissuancesRequest
		responseReissuance []types.Reissuance
		err                error
	}{
		{
			desc: "query 0 offset 10 limit",
			request: &types.QueryReissuancesRequest{
				Pagination: &query.PageRequest{
					Key:        nil,
					Offset:     0,
					Limit:      10,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseReissuance: reissuances[:10],
		},
		{
			desc: "query 1 offset 5 limit",
			request: &types.QueryReissuancesRequest{
				Pagination: &query.PageRequest{
					Key:        nil,
					Offset:     1,
					Limit:      5,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseReissuance: reissuances[1:6],
		},
		{
			desc: "query 5*ReIssuanceEpochs key 0 offset 10 limit",
			request: &types.QueryReissuancesRequest{
				Pagination: &query.PageRequest{
					Key:        util.SerializeInt64(4),
					Offset:     0,
					Limit:      10,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseReissuance: reissuances[4:14],
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := keeper.Reissuances(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, tc.err, err)
			} else {
				require.Equal(t, tc.responseReissuance, res.Reissuances)
			}
		})
	}
}
