package keeper_test

import (
	"crypto/sha256"
	"strconv"
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/asset/keeper"
	"github.com/planetmint/planetmint-go/x/asset/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func createNAsset(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Asset {
	items := make([]types.Asset, n)
	for i := range items {
		hash := sha256.Sum256([]byte(strconv.FormatInt(int64(i), 2)))
		hashStr := string(hash[:])

		items[i].Hash = hashStr
		items[i].Pubkey = "pubkey"
		items[i].Signature = "sign"
		if i%2 == 1 {
			items[i].Pubkey = "pubkey_search"
		}
		keeper.StoreAsset(ctx, items[i])
	}
	return items
}

func TestGetAsset(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	items := createNAsset(keeper, ctx, 10)
	for _, item := range items {
		asset, found := keeper.GetAsset(ctx, item.Hash)
		assert.True(t, found)
		assert.Equal(t, item, asset)
	}
}
func TestGetCids(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	items := createNAsset(keeper, ctx, 10)
	for _, item := range items {
		asset, found := keeper.GetAsset(ctx, item.Hash)
		assert.True(t, found)
		assert.Equal(t, item, asset)
	}
}

func TestGetAssetByPubKeys(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	_ = createNAsset(keeper, ctx, 10)
	assets, found := keeper.GetCIDsByPublicKey(ctx, "pubkey_search")
	assert.True(t, found)
	assert.Equal(t, len(assets), 5)
	assets, found = keeper.GetCIDsByPublicKey(ctx, "pubkey")
	assert.True(t, found)
	assert.Equal(t, len(assets), 5)
}

func TestGetCidsByPubKeys(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	_ = createNAsset(keeper, ctx, 10)
	assets, found := keeper.GetCidsByPublicKey(ctx, "pubkey_search")
	assert.True(t, found)
	assert.Equal(t, len(assets), 5)
	assets, found = keeper.GetCidsByPublicKey(ctx, "pubkey")
	assert.True(t, found)
	assert.Equal(t, len(assets), 5)
}
