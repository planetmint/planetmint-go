package keeper

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func GetReissuanceCommand(assetID string, BlockHeight int64) string {
	return "reissueasset " + assetID + " 99869000000"
}

func IsValidReissuanceCommand(reissuanceStr string, assetID string, BlockHeight int64) bool {
	expected := "reissueasset " + assetID + " 99869000000"
	return reissuanceStr == expected
}

func (k Keeper) StoreReissuance(ctx sdk.Context, reissuance types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	appendValue := k.cdc.MustMarshal(&reissuance)
	store.Set(getReissuanceBytes(reissuance.BlockHeight), appendValue)
}

func (k Keeper) LookupReissuance(ctx sdk.Context, height int64) (val types.Reissuance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	reissuance := store.Get(getReissuanceBytes(height))
	if reissuance == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(reissuance, &val)
	return val, true
}

func (k Keeper) getReissuancesRange(ctx sdk.Context, from int64) (reissuances []types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))

	iterator := store.Iterator(getReissuanceBytes(from), nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		reissuance := iterator.Value()
		var reissuanceOrg types.Reissuance
		k.cdc.MustUnmarshal(reissuance, &reissuanceOrg)
		reissuances = append(reissuances, reissuanceOrg)
	}
	return reissuances
}

func (k Keeper) getReissuancesPage(ctx sdk.Context, key []byte, offset uint64, pageSize uint64, all bool, reverse bool) (reissuances []types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	if reverse {
		iterator = store.ReverseIterator(nil, nil)
		defer iterator.Close()
	}

	for ; iterator.Valid(); iterator.Next() {
		reissuance := iterator.Value()
		var reissuanceOrg types.Reissuance
		k.cdc.MustUnmarshal(reissuance, &reissuanceOrg)
		reissuances = append(reissuances, reissuanceOrg)
	}
	return reissuances
}

func getReissuanceBytes(height int64) []byte {
	// Adding 1 because 0 will be interpreted as nil, which is an invalid key
	return big.NewInt(height + 1).Bytes()
}
