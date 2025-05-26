package keeper

import (
	dertypes "github.com/planetmint/planetmint-go/x/der/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStoreDerAssetAndLookupDerAsset(t *testing.T) {
	keeper, ctx := CreateTestKeeper(t)
	asset := dertypes.DER{
		ZigbeeID: "test-zigbee-id",
	}

	keeper.StoreDerAsset(ctx, asset)
	result, found := keeper.LookupDerAsset(ctx, "test-zigbee-id")
	require.True(t, found)
	require.Equal(t, asset.ZigbeeID, result.ZigbeeID)
}

func TestLookupDerAsset_NotFound(t *testing.T) {
	keeper, ctx := CreateTestKeeper(t)
	_, found := keeper.LookupDerAsset(ctx, "nonexistent-id")
	require.False(t, found)
}
