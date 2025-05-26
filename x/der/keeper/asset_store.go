package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
)

// storeAsset is a helper for storing any asset type.
func (k Keeper) storeAsset(ctx sdk.Context, keyPrefix []byte, zigbeeID string, value []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), keyPrefix)
	store.Set(util.SerializeString(zigbeeID), value)
}

// lookupAsset is a helper for looking up any asset type.
func (k Keeper) lookupAsset(ctx sdk.Context, keyPrefix []byte, zigbeeID string) (bz []byte, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), keyPrefix)
	bz = store.Get(util.SerializeString(zigbeeID))
	return bz, bz != nil
}
