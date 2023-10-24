package keeper_test

import (
	"fmt"
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func createNReissuances(k *keeper.Keeper, ctx sdk.Context, n int) []types.Reissuance {
	items := make([]types.Reissuance, n)
	for i := range items {
		items[i].BlockHeight = int64(i)
		items[i].Proposer = fmt.Sprintf("proposer_%v", i)
		items[i].Rawtx = keeper.GetReissuanceCommand("asset_id", int64(i))
		items[i].TxId = ""
		k.StoreReissuance(ctx, items[i])
	}
	return items
}

func TestGetReissuances(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNReissuances(keeper, ctx, 10)
	for _, item := range items {
		reissuance, found := keeper.LookupReissuance(ctx, item.BlockHeight)
		assert.True(t, found)
		assert.Equal(t, item, reissuance)
	}
}
