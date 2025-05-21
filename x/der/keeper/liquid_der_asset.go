package keeper

import (
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/der/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreLiquidDerAttest(ctx sdk.Context, asset types.LiquidDerAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiquidDerAssetKey))
	appendValue := k.cdc.MustMarshal(&asset)
	store.Set(util.SerializeString(asset.ZigbeeID), appendValue)
}

func (k Keeper) LookupLiquidDerAsset(ctx sdk.Context, zigbeeID string) (val types.LiquidDerAsset, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiquidDerAssetKey))
	derAsset := store.Get(util.SerializeString(zigbeeID))

	if derAsset == nil {
		return val, false
	}
	if err := k.cdc.Unmarshal(derAsset, &val); err != nil {
		return val, false
	}
	return val, true
}
