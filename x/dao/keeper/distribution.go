package keeper

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k Keeper) StoreDistributionOrder(ctx sdk.Context, distributionOrder types.DistributionOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
	appendValue := k.cdc.MustMarshal(&distributionOrder)
	store.Set(util.SerializeInt64(distributionOrder.LastPop), appendValue)
}

func (k Keeper) LookupDistributionOrder(ctx sdk.Context, height int64) (val types.DistributionOrder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
	distributionOrder := store.Get(util.SerializeInt64(height))
	if distributionOrder == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(distributionOrder, &val)
	return val, true
}

func (k Keeper) GetLastDistributionOrder(ctx sdk.Context) (val types.DistributionOrder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))

	iterator := store.ReverseIterator(nil, nil)
	defer iterator.Close()
	found = iterator.Valid()
	if found {
		distributionOrder := iterator.Value()
		k.cdc.MustUnmarshal(distributionOrder, &val)
	}
	return val, found
}

func (k Keeper) ComputeDistribution(ctx sdk.Context, lastReissuance int64, blockHeight int64, amount uint64) (distribution types.DistributionOrder) {
	distribution.FirstPop = lastReissuance
	distribution.LastPop = blockHeight

	distribution.DaoAddr = k.GetParams(ctx).DistributionAddressDao
	distribution.EarlyInvAddr = k.GetParams(ctx).DistributionAddressEarlyInv
	distribution.InvestorAddr = k.GetParams(ctx).DistributionAddressInvestor
	distribution.StrategicAddr = k.GetParams(ctx).DistributionAddressStrategic
	distribution.PopAddr = k.GetParams(ctx).DistributionAddressPop

	// PoP rewards subtracted from DaoAmount and added to PoPAmount for later distribution
	validatorPoPRewards, err := k.accumulateValidatorPoPRewardsForDistribution(ctx, lastReissuance, blockHeight)
	if err != nil {
		util.GetAppLogger().Error(ctx, err, "calculating Validator PoP rewards from height %v to %v", lastReissuance, blockHeight)
	}

	distribution.DaoAmount = util.UintValueToRDDLTokenString(uint64(float64(amount)*types.PercentageDao) - validatorPoPRewards)
	distribution.EarlyInvAmount = util.UintValueToRDDLTokenString(uint64(float64(amount) * types.PercentageEarlyInvestor))
	distribution.InvestorAmount = util.UintValueToRDDLTokenString(uint64(float64(amount) * types.PercentageInvestor))
	distribution.StrategicAmount = util.UintValueToRDDLTokenString(uint64(float64(amount) * types.PercentageStrategic))
	distribution.PopAmount = util.UintValueToRDDLTokenString(uint64(float64(amount)*types.PercentagePop) + validatorPoPRewards)

	return distribution
}

func (k Keeper) accumulateValidatorPoPRewardsForDistribution(ctx sdk.Context, firstPop int64, lastPop int64) (amount uint64, err error) {
	challenges, err := k.GetChallengeRange(ctx, firstPop, lastPop)
	if err != nil {
		return 0, err
	}
	for _, challenge := range challenges {
		reward, found := k.getChallengeInitiatorReward(ctx, challenge.GetHeight())
		if found {
			amount += reward
		}
	}
	return amount, nil
}

func getUint64FromTxString(ctx sdk.Context, tx string) (amount uint64, err error) {
	subStrings := strings.Split(tx, " ")
	if len(subStrings) < 3 {
		ctx.Logger().Error("Reissue tx string is shorter than expected. " + tx)
	} else {
		value := subStrings[2]
		amount, err = util.RDDLTokenStringToUint(value)
		if err != nil {
			ctx.Logger().Error("Reissue tx string value is invalid " + subStrings[2])
		}
	}
	return amount, err
}

func (k Keeper) GetDistributionForReissuedTokens(ctx sdk.Context, blockHeight int64) (distribution types.DistributionOrder, err error) {
	var lastPoP int64
	lastDistributionOrder, found := k.GetLastDistributionOrder(ctx)
	if found {
		lastPoP = lastDistributionOrder.LastPop
	}

	reissuances := k.getReissuancesRange(ctx, lastPoP)
	var overallAmount uint64
	for index, obj := range reissuances {
		if (index == 0 && lastPoP == 0 && obj.BlockHeight == 0) || // corner case (beginning of he chain)
			(lastPoP < obj.BlockHeight && obj.BlockHeight <= blockHeight) {
			amount, err := getUint64FromTxString(ctx, obj.GetCommand())
			if err == nil {
				overallAmount += amount
			}
		} else {
			ctx.Logger().Info("%u %u %u", lastPoP, obj.BlockHeight, blockHeight)
		}
	}
	distribution = k.ComputeDistribution(ctx, lastPoP, blockHeight, overallAmount)
	return distribution, err
}
