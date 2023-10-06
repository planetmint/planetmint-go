package keeper

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k Keeper) StoreReissuance(ctx sdk.Context, reissuance types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	appendValue := k.cdc.MustMarshal(&reissuance)
	store.Set(getReissuanceBytes(reissuance.BlockHeight), appendValue)
}

func (k Keeper) LookupReissuance(ctx sdk.Context, height uint64) (val types.Reissuance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	reissuance := store.Get(getReissuanceBytes(height))
	if reissuance == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(reissuance, &val)
	return val, true
}

func (k Keeper) getReissuancesPage(ctx sdk.Context, key []byte, offset uint64, page_size uint64, all bool, reverse bool) (reissuances []types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	if reverse {
		iterator = store.ReverseIterator(nil, nil)
		defer iterator.Close()
	}

	for ; iterator.Valid(); iterator.Next() {
		reissuance := iterator.Value()
		var reissuance_org types.Reissuance
		k.cdc.MustUnmarshal(reissuance, &reissuance_org)
		reissuances = append(reissuances, reissuance_org)
	}
	return reissuances
}

func getReissuanceBytes(height uint64) []byte {
	// Adding 1 because 0 will be interpreted as nil, which is an invalid key
	return big.NewInt(int64(height + 1)).Bytes()
}
