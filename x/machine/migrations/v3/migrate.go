package v3

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, _ codec.BinaryCodec) error {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.TrustAnchorKey))

	count := uint64(0)
	iterator := store.Iterator(nil, nil)
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
