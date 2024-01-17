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
		items[i].LiquidTxHash = strconv.Itoa(i)

		keeper.SetRedeemClaim(ctx, items[i])
	}
	return items
}

func TestRedeemClaimGet(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNRedeemClaim(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRedeemClaim(ctx,
			item.Beneficiary,
			item.LiquidTxHash,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRedeemClaimRemove(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNRedeemClaim(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRedeemClaim(ctx,
			item.Beneficiary,
			item.LiquidTxHash,
		)
		_, found := keeper.GetRedeemClaim(ctx,
			item.Beneficiary,
			item.LiquidTxHash,
		)
		require.False(t, found)
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
