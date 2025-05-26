package keeper

import (
	"testing"

	dertypes "github.com/planetmint/planetmint-go/x/der/types"
	"github.com/stretchr/testify/require"
)

func TestStoreLiquidDerAssetAndLookupLiquidDerAsset(t *testing.T) {
	keeper, ctx := CreateTestKeeper(t)
	asset := dertypes.LiquidDerAsset{
		ZigbeeID: "liquid-test-zigbee-id",
	}
	keeper.StoreLiquidDerAsset(ctx, asset)
	result, found := keeper.LookupLiquidDerAsset(ctx, "liquid-test-zigbee-id")
	require.True(t, found)
	require.Equal(t, asset.ZigbeeID, result.ZigbeeID)
}

func TestLookupLiquidDerAsset_NotFound(t *testing.T) {
	keeper, ctx := CreateTestKeeper(t)
	_, found := keeper.LookupLiquidDerAsset(ctx, "nonexistent-id")
	require.False(t, found)
}
