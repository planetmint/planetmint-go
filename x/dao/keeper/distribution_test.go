package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func createNDistributionOrder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DistributionOrder {
	items := make([]types.DistributionOrder, n)
	for i := range items {
		amount := 10000
		items[i].FirstPop = int64(i)*1000 + 1
		items[i].LastPop = int64(i+1) * 1000
		items[i].DaoAddr = fmt.Sprintf("DAO%v", i)
		items[i].DaoAmount = fmt.Sprintf("%v", float64(amount)*types.PercentageDao)
		items[i].InvestorAddr = fmt.Sprintf("INVESTOR%v", i)
		items[i].InvestorAmount = strconv.FormatInt(int64(float64(amount)*types.PercentageInvestor), 10)
		items[i].PopAddr = fmt.Sprintf("POP%v", i)
		items[i].PopAmount = fmt.Sprintf("%v", float64(amount)*types.PercentagePop)
		keeper.StoreDistributionOrder(ctx, items[i])
	}
	return items
}

func TestDistributionOrder(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNDistributionOrder(keeper, ctx, 10)
	for _, item := range items {
		challenge, found := keeper.LookupDistributionOrder(ctx, item.LastPop)
		assert.True(t, found)
		assert.Equal(t, item, challenge)
	}

	lastDistribution, found := keeper.GetLastDistributionOrder(ctx)
	assert.True(t, found)
	assert.Equal(t, items[9], lastDistribution)
}

func TestTokenDistribution(t *testing.T) {
	t.Parallel()
	k, ctx := keepertest.DaoKeeper(t)
	var reissuanceValue uint64 = 99869000000
	reissuances := 1000
	var Amount1stBatch uint64 = 781
	var Amount2ndBatch uint64 = 219
	_ = createNReissuances(k, ctx, reissuances)
	distribution, err := k.GetDistributionForReissuedTokens(ctx, 780)
	assert.Nil(t, err)
	amount1, err1 := strconv.ParseUint(distribution.DaoAmount, 10, 64)
	amount2, err2 := strconv.ParseUint(distribution.InvestorAmount, 10, 64)
	amount3, err3 := strconv.ParseUint(distribution.PopAmount, 10, 64)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	sum := amount1 + amount2 + amount3
	expSum := reissuanceValue * Amount1stBatch // add the [0] of the
	assert.Equal(t, expSum, sum)

	var lastDistribution types.DistributionOrder
	lastDistribution.LastPop = 780
	k.StoreDistributionOrder(ctx, lastDistribution)
	lastDistribution, err0 := k.GetDistributionForReissuedTokens(ctx, 999)
	assert.Nil(t, err0)
	amount1, err1 = strconv.ParseUint(lastDistribution.DaoAmount, 10, 64)
	amount2, err2 = strconv.ParseUint(lastDistribution.InvestorAmount, 10, 64)
	amount3, err3 = strconv.ParseUint(lastDistribution.PopAmount, 10, 64)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	sum = amount1 + amount2 + amount3
	expSum = reissuanceValue * Amount2ndBatch // add the [0] of the
	assert.Equal(t, expSum, sum)
	assert.Equal(t, uint64(reissuances), Amount1stBatch+Amount2ndBatch)
}
