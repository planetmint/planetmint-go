package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/nullify"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRedeemClaim(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RedeemClaim {
	items := make([]types.RedeemClaim, n)
	for i := range items {
		items[i].Beneficiary = strconv.Itoa(i)
		id := keeper.CreateNewRedeemClaim(ctx, items[i])
		items[i].Id = id
	}
	return items
}

func TestRedeemClaimGet(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNRedeemClaim(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRedeemClaim(ctx,
			item.Beneficiary,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestRedeemClaimGetAll(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNRedeemClaim(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRedeemClaim(ctx)),
	)
}

func TestRedeemClaimByLiquidTXHash(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNRedeemClaim(keeper, ctx, 10)
	for i, item := range items {
		item.LiquidTxHash = strconv.Itoa(i)
		keeper.SetRedeemClaimByLiquidTXHash(ctx, item)
		rc, found := keeper.GetRedeemClaimByLiquidTXHash(ctx, strconv.Itoa(i))
		require.True(t, found)
		require.Equal(t, item, rc)
	}
}
