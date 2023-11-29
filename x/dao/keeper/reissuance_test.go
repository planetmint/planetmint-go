package keeper_test

import (
	"fmt"
	"testing"

	"github.com/planetmint/planetmint-go/config"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	daokeeper "github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func createNReissuances(k *daokeeper.Keeper, ctx sdk.Context, n int) []types.Reissuance {
	items := make([]types.Reissuance, n)
	for i := range items {
		items[i].BlockHeight = int64(i)
		items[i].Proposer = fmt.Sprintf("proposer_%v", i)
		items[i].RawTx = daokeeper.GetReissuanceCommand("asset_id", int64(i))
		items[i].TxID = ""
		items[i].FirstPop = int64(i)
		items[i].LastPop = int64(i)
		k.StoreReissuance(ctx, items[i])
	}
	return items
}

func TestGetReissuances(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNReissuances(keeper, ctx, 10)
	for _, item := range items {
		reissuance, found := keeper.LookupReissuance(ctx, item.BlockHeight)
		assert.True(t, found)
		assert.Equal(t, item, reissuance)
	}
}

func TestReissuanceValueComputation(t *testing.T) {
	t.Parallel()
	popsPerEpoch := float64(config.GetConfig().PopEpochs)
	assert.Equal(t, "998.69000000", daokeeper.GetReissuanceAsStringValue(1))
	assert.Equal(t, "499.34000000", daokeeper.GetReissuanceAsStringValue(int64(daokeeper.PopsPerCycle*popsPerEpoch*1+1)))
	assert.Equal(t, "249.67000000", daokeeper.GetReissuanceAsStringValue(int64(daokeeper.PopsPerCycle*popsPerEpoch*2+1)))
	assert.Equal(t, "124.83000000", daokeeper.GetReissuanceAsStringValue(int64(daokeeper.PopsPerCycle*popsPerEpoch*3+1)))
	assert.Equal(t, "62.42000000", daokeeper.GetReissuanceAsStringValue(int64(daokeeper.PopsPerCycle*popsPerEpoch*4+1)))
	assert.Equal(t, "0.0", daokeeper.GetReissuanceAsStringValue(int64(daokeeper.PopsPerCycle*popsPerEpoch*5+1)))
}
