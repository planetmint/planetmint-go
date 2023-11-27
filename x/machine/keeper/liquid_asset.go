package keeper

import (
	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreLiquidAttest(ctx sdk.Context, asset types.LiquidAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiquidAssetKey))
	appendValue := k.cdc.MustMarshal(&asset)
	store.Set(GetAssetBytes(asset.MachineID), appendValue)
}

func (k Keeper) LookupLiquidAsset(ctx sdk.Context, machineID string) (val types.LiquidAsset, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiquidAssetKey))
	liquidAsset := store.Get(GetAssetBytes(machineID))

	if liquidAsset == nil {
		return val, false
	}
	if err := k.cdc.Unmarshal(liquidAsset, &val); err != nil {
		return val, false
	}
	return val, true
}

func GetAssetBytes(pubKey string) []byte {
	return []byte(pubKey)
}
