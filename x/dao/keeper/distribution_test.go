package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/util"
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
	reissuanceValue := 998.85844748
	var reissuances int64 = 1000
	Amount1stBatch := 781.0
	Amount2ndBatch := 219.0

	_ = createNReissuances(k, ctx, reissuances, types.DefaultGenesis().GetParams().PopEpochs)
	distribution, err := k.GetDistributionForReissuedTokens(ctx, 780)
	assert.Nil(t, err)

	amount1, err1 := strconv.ParseFloat(distribution.DaoAmount, 64)
	amount2, err2 := strconv.ParseFloat(distribution.InvestorAmount, 64)
	amount3, err3 := strconv.ParseFloat(distribution.PopAmount, 64)
	amount4, err4 := strconv.ParseFloat(distribution.EarlyInvAmount, 64)
	amount5, err5 := strconv.ParseFloat(distribution.StrategicAmount, 64)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.Nil(t, err4)
	assert.Nil(t, err5)
	sum := amount1 + amount2 + amount3 + amount4 + amount5
	expSum := reissuanceValue * Amount1stBatch // add the [0] of the
	assert.True(t, expSum-sum < 0.000001)

	var lastDistribution types.DistributionOrder
	lastDistribution.LastPop = 780
	k.StoreDistributionOrder(ctx, lastDistribution)
	lastDistribution, err0 := k.GetDistributionForReissuedTokens(ctx, 999)
	assert.Nil(t, err0)
	amount1, err1 = strconv.ParseFloat(lastDistribution.DaoAmount, 64)
	amount2, err2 = strconv.ParseFloat(lastDistribution.InvestorAmount, 64)
	amount3, err3 = strconv.ParseFloat(lastDistribution.PopAmount, 64)
	amount4, err4 = strconv.ParseFloat(lastDistribution.EarlyInvAmount, 64)
	amount5, err5 = strconv.ParseFloat(lastDistribution.StrategicAmount, 64)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.Nil(t, err4)
	assert.Nil(t, err5)
	sum = amount1 + amount2 + amount3 + amount4 + amount5
	expSum = reissuanceValue * Amount2ndBatch // add the [0] of the
	assert.True(t, expSum-sum < 0.000001)
	assert.Equal(t, float64(reissuances), Amount1stBatch+Amount2ndBatch)
}

func TestValidatorRewardReallocation(t *testing.T) {
	t.Parallel()
	k, ctx := keepertest.DaoKeeper(t)

	// create reissuances
	_ = createNReissuances(k, ctx, 1000, types.DefaultParams().PopEpochs)
	distribution, err := k.GetDistributionForReissuedTokens(ctx, 780)
	assert.NoError(t, err)

	// create challenges
	_ = createNChallenge(k, ctx, 250, types.DefaultParams().PopEpochs)
	distributionWithPopReallocation, err := k.GetDistributionForReissuedTokens(ctx, 780)
	assert.NoError(t, err)

	// compare
	pops := 780 / types.DefaultParams().PopEpochs
	distributionDaoAmount, err := util.RDDLTokenStringToFloat(distribution.DaoAmount)
	assert.NoError(t, err)
	distributionPopAmount, err := util.RDDLTokenStringToFloat(distribution.PopAmount)
	assert.NoError(t, err)

	distributionWithValidatorPopRewardDao, err := util.RDDLTokenStringToFloat(distributionWithPopReallocation.DaoAmount)
	assert.NoError(t, err)
	distributionWithValidatorPopRewardPop, err := util.RDDLTokenStringToFloat(distributionWithPopReallocation.PopAmount)
	assert.NoError(t, err)

	// Some of DaoAmount is transferred to PopAmount
	assert.Greater(t, distributionDaoAmount, distributionWithValidatorPopRewardDao)
	assert.Greater(t, distributionWithValidatorPopRewardPop, distributionPopAmount)

	diffPopAmount := distributionWithValidatorPopRewardPop - distributionPopAmount
	diffDaoAmount := distributionDaoAmount - distributionWithValidatorPopRewardDao
	assert.Equal(t, diffPopAmount, diffDaoAmount)
	assert.Equal(t, pops, int64(diffPopAmount))
	assert.Equal(t, pops, int64(diffDaoAmount))
}
