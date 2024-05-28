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

func createNAsset(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MsgNotarizeAsset {
	items := make([]types.MsgNotarizeAsset, n)
	for i := range items {
		hash := sha256.Sum256([]byte(strconv.FormatInt(int64(i), 2)))
		hashStr := string(hash[:])
		items[i].Cid = hashStr
		items[i].Creator = "plmnt_address"
		if i%2 == 0 {
			items[i].Creator = "plmnt_address1"
		}
		keeper.StoreAsset(ctx, items[i])
	}
	return items
}

func TestGetAssetbyCid(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.AssetKeeper(t)
	items := createNAsset(keeper, ctx, 10)
	for _, item := range items {
		asset, found := keeper.GetAsset(ctx, item.Cid)
		assert.True(t, found)
		assert.Equal(t, item, asset)
	}
}

func TestGetAssetByPubKeys(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.AssetKeeper(t)
	_ = createNAsset(keeper, ctx, 10)
	assets, found := keeper.GetCidsByAddress(ctx, "plmnt_address")
	assert.True(t, found)
	assert.Equal(t, len(assets), 1) // TODO: just for HF: before 5 
	assets, found = keeper.GetCidsByAddress(ctx, "plmnt_address1")
	assert.True(t, found)
	assert.Equal(t, len(assets), 1) // TODO: just for HF: before 5
