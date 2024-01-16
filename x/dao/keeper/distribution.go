package keeper

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
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

// TODO to be integrated at a later stage
// func (k Keeper) getDistributionRequestPage(ctx sdk.Context, key []byte, offset uint64, page_size uint64, all bool, reverse bool) (distribution_orders []types.DistributionOrder) {
// 	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))

// 	iterator := store.Iterator(nil, nil)
// 	defer iterator.Close()
// 	if reverse {
// 		iterator = store.ReverseIterator(nil, nil)
// 		defer iterator.Close()
// 	}
// 	for ; iterator.Valid(); iterator.Next() {
// 		distribution_order := iterator.Value()
// 		var distribution_order_org types.DistributionOrder
// 		k.cdc.MustUnmarshal(distribution_order, &distribution_order_org)
// 		distribution_orders = append(distribution_orders, distribution_order_org)
// 	}
// 	return distribution_orders
// }

func ComputeDistribution(lastReissuance int64, blockHeight int64, amount uint64) (distribution types.DistributionOrder) {
	conf := config.GetConfig()
	distribution.FirstPop = lastReissuance
	distribution.LastPop = blockHeight

	distribution.DaoAddr = conf.DistributionAddrDAO
	distribution.InvestorAddr = conf.DistributionAddrInv
	distribution.PopAddr = conf.DistributionAddrPop

	distribution.DaoAmount = util.UintValueToRDDLTokenString(uint64(float64(amount) * types.PercentageDao))
	distribution.InvestorAmount = util.UintValueToRDDLTokenString(uint64(float64(amount) * types.PercentageInvestor))
	distribution.PopAmount = util.UintValueToRDDLTokenString(uint64(float64(amount) * types.PercentagePop))

	return distribution
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
	distribution = ComputeDistribution(lastPoP, blockHeight, overallAmount)
	return distribution, err
}
