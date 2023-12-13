package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/planetmint/planetmint-go/config"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/require"
)

func TestQueryChallenges(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	challenges := createNChallenge(keeper, ctx, 20)
	for _, tc := range []struct {
		desc               string
		request            *types.QueryChallengesRequest
		responseChallenges []types.Challenge
		err                error
	}{
		{
			desc: "query 0 offset 10 limit",
			request: &types.QueryChallengesRequest{
				Pagination: &query.PageRequest{
					Key:        nil,
					Offset:     0,
					Limit:      10,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseChallenges: challenges[:10],
		},
		{
			desc: "query 1 offset 5 limit",
			request: &types.QueryChallengesRequest{
				Pagination: &query.PageRequest{
					Key:        nil,
					Offset:     1,
					Limit:      5,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseChallenges: challenges[1:6],
		},
		{
			desc: "query 5*PopEpochs key 0 offset 10 limit",
			request: &types.QueryChallengesRequest{
				Pagination: &query.PageRequest{
					Key:        util.SerializeInt64(int64(config.GetConfig().PopEpochs * 5)),
					Offset:     0,
					Limit:      10,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseChallenges: challenges[4:14],
		},
		{
			desc: "query 2*PopEpochs-5 key 0 offset 10 limit",
			request: &types.QueryChallengesRequest{
				Pagination: &query.PageRequest{
					Key:        util.SerializeInt64(int64(config.GetConfig().PopEpochs*2 - 5)),
					Offset:     0,
					Limit:      10,
					CountTotal: true,
					Reverse:    false,
				},
			},
			responseChallenges: challenges[1:11],
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := keeper.Challenges(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, tc.err, err)
			} else {
				require.Equal(t, tc.responseChallenges, res.Challenges)
			}
		})
	}
}
