package keeper

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k Keeper) StoreDistributionOrder(ctx sdk.Context, distribution_order types.DistributionOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
	appendValue := k.cdc.MustMarshal(&distribution_order)
	store.Set(getLastPopBytes(distribution_order.LastPop), appendValue)
}

func (k Keeper) LookupDistributionOrder(ctx sdk.Context, lastPopHeight uint64) (val types.DistributionOrder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
	distribution_order := store.Get(getLastPopBytes(lastPopHeight))
	if distribution_order == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(distribution_order, &val)
	return val, true
}

func (k Keeper) GetLastDistributionOrder(ctx sdk.Context) (val types.DistributionOrder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))

	iterator := store.ReverseIterator(nil, nil)
	defer iterator.Close()
	found = iterator.Valid()
	if found {
		distribution_order := iterator.Value()
		k.cdc.MustUnmarshal(distribution_order, &val)
	}
	return val, found
}

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

func getLastPopBytes(height uint64) []byte {
	// Adding 1 because 0 will be interpreted as nil, which is an invalid key
	return big.NewInt(int64(height + 1)).Bytes()
}

func ComputeDistribution(lastReissuance uint64, BlockHeight int64, amount uint64) (distribution types.DistributionOrder) {
	conf := config.GetConfig()
	distribution.FirstPop = lastReissuance
	distribution.LastPop = uint64(BlockHeight)

	distribution.DaoAddr = conf.DistributionAddrDAO
	distribution.InvestorAddr = conf.DistributionAddrInv
	distribution.PopAddr = conf.DistributionAddrPoP

	distribution.DaoAmount = strconv.FormatUint(uint64(float64(amount)*types.PercentageDao), 10)
	distribution.InvestorAmount = strconv.FormatUint(uint64(float64(amount)*types.PercentageInvestor), 10)
	distribution.PopAmount = strconv.FormatUint(uint64(float64(amount)*types.PercentagePop), 10)

	return distribution
}

func getUintFromTXString(ctx sdk.Context, tx string) (amount uint64, err error) {
	sub_strs := strings.Split(tx, " ")
	if len(sub_strs) < 3 {
		ctx.Logger().Error("Reissue TX string is shorter than expected. " + tx)
	} else {
		value := sub_strs[2]

		amount, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			ctx.Logger().Error("Reissue TX string value is invalid " + sub_strs[2])
		}
	}
	return amount, err
}

func (k Keeper) GetDistributenForReissuedTokens(ctx sdk.Context, blockHeight int64) (distribution types.DistributionOrder, err error) {
	var lastPoP uint64 = 0
	lastDistributionOrder, found := k.GetLastDistributionOrder(ctx)
	if found {
		lastPoP = lastDistributionOrder.LastPop
	}

	reissuances := k.getReissuancesRange(ctx, lastPoP)
	var overall_amount uint64 = 0
	for index, obj := range reissuances {
		if (index == 0 && lastPoP == 0 && obj.BlockHeight == 0) || //corner case (beginning of he chain)
			(int64(lastPoP) < int64(obj.BlockHeight) && int64(obj.BlockHeight) <= blockHeight) {
			amount, err := getUintFromTXString(ctx, obj.Rawtx)
			if err == nil {
				overall_amount = overall_amount + amount
			}
		} else {
			ctx.Logger().Info("%u %u %u", lastPoP, obj.BlockHeight, blockHeight)
		}
	}
	distribution = ComputeDistribution(lastPoP, blockHeight, overall_amount)
	return distribution, err
}
