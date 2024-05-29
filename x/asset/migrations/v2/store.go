package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/asset/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, _ codec.BinaryCodec) error {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.AssetKey))

	mapping := make(map[string][][]byte)

	// read all cids
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		addressBytes := iterator.Value()
		cidBytes := iterator.Key()

		// map all cids by address
		mapping[string(addressBytes)] = append(mapping[string(addressBytes)], cidBytes)
	}

	// store all cids with new key
	for address, cids := range mapping {
		assetByAddressStore := prefix.NewStore(ctx.KVStore(storeKey), types.AddressPrefix(address))
		for i, cid := range cids {
			assetByAddressStore.Set(util.SerializeUint64(uint64(i)), cid)
		}
		addressAssetCountStore := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.AssetKey))
		addressAssetCountStore.Set(types.AddressCountKey(address), util.SerializeUint64(uint64(len(cids))))
	}

	return nil
}
