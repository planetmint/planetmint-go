package keeper

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) getDistributionRequestPage(ctx sdk.Context, key []byte, offset uint64, page_size uint64, all bool, reverse bool) (distribution_orders []types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	if reverse {
		iterator = store.ReverseIterator(nil, nil)
		defer iterator.Close()
	}

	for ; iterator.Valid(); iterator.Next() {
		distribution_order := iterator.Value()
		var distribution_order_org types.DistributionOrder
		k.cdc.MustUnmarshal(distribution_order, &distribution_order_org)
		distribution_orders = append(distribution_orders, distribution_order_org)
	}
	return distribution_orders
}

func getLastPopBytes(height uint64) []byte {
	// Adding 1 because 0 will be interpreted as nil, which is an invalid key
	return big.NewInt(int64(height + 1)).Bytes()
}
