package keeper

import (
	"github.com/planetmint/planetmint-go/x/asset/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreAsset(ctx sdk.Context, asset types.Asset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	appendValue := k.cdc.MustMarshal(&asset)
	store.Set(GetAssetHashBytes(asset.Hash), appendValue)
}

func (k Keeper) GetAsset(ctx sdk.Context, hash string) (val types.Asset, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	asset := store.Get(GetAssetHashBytes(hash))
	if asset == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(asset, &val)
	return val, true
}

func (k Keeper) GetAssetsByPublicKey(ctx sdk.Context, pubkey string) (assetArray []types.Asset, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))

	reverseIterator := store.ReverseIterator(nil, nil)
	defer reverseIterator.Close()
	var asset types.Asset
	for ; reverseIterator.Valid(); reverseIterator.Next() {
		lastValue := reverseIterator.Value()

		k.cdc.MustUnmarshal(lastValue, &asset)
		if asset.GetPubkey() == pubkey {
			assetArray = append(assetArray, asset)
		}
	}
	return assetArray, len(assetArray) > 0
}

func (k Keeper) GetCidsByPublicKey(ctx sdk.Context, pubkey string) (cids []string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))

	reverseIterator := store.ReverseIterator(nil, nil)
	defer reverseIterator.Close()
	var asset types.Asset
	for ; reverseIterator.Valid(); reverseIterator.Next() {
		lastValue := reverseIterator.Value()

		k.cdc.MustUnmarshal(lastValue, &asset)
		if asset.GetPubkey() == pubkey {
			cids = append(cids, asset.GetHash())
		}
	}
	return cids, len(cids) > 0
}

func GetAssetHashBytes(hash string) []byte {
	bz := []byte(hash)
	return bz
}
