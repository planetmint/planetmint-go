package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/planetmint/planetmint-go/errormsg"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/nullify"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRedeemClaimQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRedeemClaim(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetRedeemClaimRequest
		response *types.QueryGetRedeemClaimResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRedeemClaimRequest{
				Beneficiary: msgs[0].Beneficiary,
				Id:          msgs[0].Id,
			},
			response: &types.QueryGetRedeemClaimResponse{RedeemClaim: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRedeemClaimRequest{
				Beneficiary: msgs[1].Beneficiary,
				Id:          msgs[1].Id,
			},
			response: &types.QueryGetRedeemClaimResponse{RedeemClaim: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRedeemClaimRequest{
				Beneficiary: strconv.Itoa(100000),
				Id:          uint64(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, errormsg.InvalidRequest),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RedeemClaim(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestRedeemClaimQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRedeemClaim(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRedeemClaimRequest {
		return &types.QueryAllRedeemClaimRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RedeemClaimAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RedeemClaim), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RedeemClaim),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RedeemClaimAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RedeemClaim), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RedeemClaim),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RedeemClaimAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RedeemClaim),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RedeemClaimAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, errormsg.InvalidRequest))
	})
}
