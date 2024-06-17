package keeper

import (
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/machine/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreLiquidAttest(ctx sdk.Context, asset types.LiquidAsset) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LiquidAssetKey))
	appendValue := k.cdc.MustMarshal(&asset)
	store.Set(util.SerializeString(asset.MachineID), appendValue)
}

func (k Keeper) LookupLiquidAsset(ctx sdk.Context, machineID string) (val types.LiquidAsset, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LiquidAssetKey))
	liquidAsset := store.Get(util.SerializeString(machineID))

	if liquidAsset == nil {
		return val, false
	}
	if err := k.cdc.Unmarshal(liquidAsset, &val); err != nil {
		return val, false
	}
	return val, true
}
