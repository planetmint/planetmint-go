package keeper_test

import (
	"crypto/sha256"
	"strconv"
	"testing"

	keepertest "planetmint-go/testutil/keeper"
	"planetmint-go/x/asset/keeper"
	"planetmint-go/x/asset/types"

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
