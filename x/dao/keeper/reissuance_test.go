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
		items[i].FirstIncludedPop = int64(i)
		items[i].LastIncludedPop = int64(i)
		k.StoreReissuance(ctx, items[i])
	}
	return items
}

func TestReissuanceComputation(t *testing.T) {
	t.Parallel()
	k, ctx := keepertest.DaoKeeper(t)
	var reissuanceValue uint64 = 99869000000
	numChallenges := 1000
	popepoch := int64(config.GetConfig().PopEpochs)
	_ = createNChallenge(k, ctx, numChallenges)

	reIssuanceValue1, firstIncludedPop, lastIncludedPop, err := k.ComputeReIssuanceValue(ctx, 0, 780*popepoch)
	assert.Nil(t, err)

	// explaining the numbers:
	// the Pops/Challenges start with 1*PopEpoch, ... n*PopEpoch
	indexFirst := firstIncludedPop / popepoch
	indexLast := lastIncludedPop / popepoch
	assert.Equal(t, indexFirst, int64(1))
	assert.Equal(t, indexLast, int64(778))
	expSum := reissuanceValue * uint64(indexLast-indexFirst+1) // add 1 to count for the one that is missing by subtraction
	assert.Equal(t, expSum, reIssuanceValue1)

	var lastReIssuance types.Reissuance
	lastReIssuance.FirstIncludedPop = firstIncludedPop
	lastReIssuance.LastIncludedPop = lastIncludedPop
	k.StoreReissuance(ctx, lastReIssuance)
	lastReIssuanceValue2nd, firstIncludedPop, lastIncludedPop, err0 := k.ComputeReIssuanceValue(ctx, lastIncludedPop, 1000*int64(config.GetConfig().PopEpochs))
	assert.Nil(t, err0)
	indexFirst2nd := firstIncludedPop / popepoch
	indexLast2nd := lastIncludedPop / popepoch
	assert.Equal(t, indexLast+1, indexFirst2nd)
	assert.Equal(t, int64(numChallenges-2), indexLast2nd)
	expSum = reissuanceValue * uint64(indexLast2nd-indexFirst2nd+1) // add the [0] of the
	assert.Equal(t, expSum, lastReIssuanceValue2nd)
	expectedSum := uint64(numChallenges-2) * reissuanceValue
	computedSum := lastReIssuanceValue2nd + reIssuanceValue1
	assert.Equal(t, expectedSum, computedSum)
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
