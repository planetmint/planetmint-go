package v3

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

func MigrateStore(ctx sdk.Context, taStoreKey storetypes.StoreKey, storeKey storetypes.StoreKey) error {
	store := ctx.KVStore(taStoreKey)

	count := uint64(0)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.TrustAnchorKey))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		if iterator.Value()[0] == 1 {
			count++
		}
	}

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	countStore := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.ActivatedTACounterPrefix))
	countStore.Set([]byte{1}, bz)

	return nil
}
