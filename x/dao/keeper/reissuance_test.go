package keeper_test

import (
	"fmt"
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/util"
	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	daokeeper "github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func createNReissuances(k *daokeeper.Keeper, ctx sdk.Context, n int64, popsPerEpoch int64) []types.Reissuance {
	items := make([]types.Reissuance, n)
	for j := range items {
		i := int64(j)
		items[i].BlockHeight = i
		items[i].Proposer = fmt.Sprintf("proposer_%v", i)
		items[i].Command = daokeeper.GetReissuanceCommand("asset_id", i, popsPerEpoch)
		items[i].TxID = ""
		items[i].FirstIncludedPop = i
		items[i].LastIncludedPop = i
		k.StoreReissuance(ctx, items[i])
	}
	return items
}

func TestReissuanceComputation(t *testing.T) {
	t.Parallel()
	k, ctx := keepertest.DaoKeeper(t)
	var reissuanceValue uint64 = 99885844748
	numChallenges := 1000
	popepoch := types.DefaultGenesis().GetParams().PopEpochs
	_ = createNChallenge(k, ctx, numChallenges, popepoch)

	reissuanceValue1, firstIncludedPop, lastIncludedPop, err := k.ComputeReissuanceValue(ctx, 0, 780*popepoch)
	assert.Nil(t, err)

	// explaining the numbers:
	// the Pops/Challenges start with 1*PopEpoch, ... n*PopEpoch
	indexFirst := firstIncludedPop / popepoch
	indexLast := lastIncludedPop / popepoch
	assert.Equal(t, indexFirst, int64(1))
	assert.Equal(t, indexLast, int64(778))
	expSum := reissuanceValue * uint64(indexLast-indexFirst+1) // add 1 to count for the one that is missing by subtraction
	assert.Equal(t, expSum, reissuanceValue1)

	var lastReissuance types.Reissuance
	lastReissuance.FirstIncludedPop = firstIncludedPop
	lastReissuance.LastIncludedPop = lastIncludedPop
	k.StoreReissuance(ctx, lastReissuance)
	lastReissuanceValue2nd, firstIncludedPop, lastIncludedPop, err0 := k.ComputeReissuanceValue(ctx, lastIncludedPop, 1000*types.DefaultGenesis().GetParams().PopEpochs)
	assert.Nil(t, err0)
	indexFirst2nd := firstIncludedPop / popepoch
	indexLast2nd := lastIncludedPop / popepoch
	assert.Equal(t, indexLast+1, indexFirst2nd)
	assert.Equal(t, int64(numChallenges-2), indexLast2nd)
	expSum = reissuanceValue * uint64(indexLast2nd-indexFirst2nd+1) // add the [0] of the
	assert.Equal(t, expSum, lastReissuanceValue2nd)
	expectedSum := uint64(numChallenges-2) * reissuanceValue
	computedSum := lastReissuanceValue2nd + reissuanceValue1
	assert.Equal(t, expectedSum, computedSum)
}

func TestGetReissuances(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.DaoKeeper(t)
	popepoch := types.DefaultGenesis().GetParams().PopEpochs
	items := createNReissuances(keeper, ctx, 10, popepoch)
	for _, item := range items {
		reissuance, found := keeper.LookupReissuance(ctx, item.BlockHeight)
		assert.True(t, found)
		assert.Equal(t, item, reissuance)
	}
}

func TestReissuanceValueComputation(t *testing.T) {
	t.Parallel()
	popsPerEpochFlt := float64(types.DefaultGenesis().GetParams().PopEpochs)
	assert.Equal(t, "998.85844748", daokeeper.GetReissuanceAsStringValue(1, types.DefaultGenesis().GetParams().PopEpochs))
	assert.Equal(t, "499.42922374", daokeeper.GetReissuanceAsStringValue(int64(util.PopsPerCycle*popsPerEpochFlt*1+1), types.DefaultGenesis().GetParams().PopEpochs))
	assert.Equal(t, "249.71461187", daokeeper.GetReissuanceAsStringValue(int64(util.PopsPerCycle*popsPerEpochFlt*2+1), types.DefaultGenesis().GetParams().PopEpochs))
	assert.Equal(t, "124.85730593", daokeeper.GetReissuanceAsStringValue(int64(util.PopsPerCycle*popsPerEpochFlt*3+1), types.DefaultGenesis().GetParams().PopEpochs))
	assert.Equal(t, "62.42865296", daokeeper.GetReissuanceAsStringValue(int64(util.PopsPerCycle*popsPerEpochFlt*4+1), types.DefaultGenesis().GetParams().PopEpochs))
	assert.Equal(t, "0.0", daokeeper.GetReissuanceAsStringValue(int64(util.PopsPerCycle*popsPerEpochFlt*5+1), types.DefaultGenesis().GetParams().PopEpochs))
}
