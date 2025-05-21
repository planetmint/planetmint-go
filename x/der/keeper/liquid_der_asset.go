package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/der/types"
)

func (k Keeper) StoreLiquidDerAttest(ctx sdk.Context, asset types.LiquidDerAsset) {
	appendValue := k.cdc.MustMarshal(&asset)
	k.storeAsset(ctx, types.KeyPrefix(types.LiquidDerAssetKey), asset.ZigbeeID, appendValue)
}

func (k Keeper) LookupLiquidDerAsset(ctx sdk.Context, zigbeeID string) (val types.LiquidDerAsset, found bool) {
	bz, found := k.lookupAsset(ctx, types.KeyPrefix(types.LiquidDerAssetKey), zigbeeID)
	if !found {
		return val, false
	}
	if err := k.cdc.Unmarshal(bz, &val); err != nil {
		return val, false
	}
	return val, true
}
