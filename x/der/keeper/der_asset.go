package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/der/types"
)

func (k Keeper) StoreDerAsset(ctx sdk.Context, asset types.DER) {
	appendValue := k.cdc.MustMarshal(&asset)
	k.storeAsset(ctx, types.KeyPrefix(types.DerAssetKey), asset.ZigbeeID, appendValue)
}

func (k Keeper) LookupDerAsset(ctx sdk.Context, zigbeeID string) (val types.DER, found bool) {
	bz, found := k.lookupAsset(ctx, types.KeyPrefix(types.DerAssetKey), zigbeeID)
	if !found {
		return val, false
	}
	if err := k.cdc.Unmarshal(bz, &val); err != nil {
		return val, false
	}
	return val, true
}
