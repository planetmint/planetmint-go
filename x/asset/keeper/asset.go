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

func GetAssetHashBytes(hash string) []byte {
	bz := []byte(hash)
	return bz
}
