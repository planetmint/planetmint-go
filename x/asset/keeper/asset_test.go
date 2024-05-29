package keeper_test

import (
	"crypto/sha256"
	"strconv"
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/util"
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
}

func TestAssetCount(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.AssetKeeper(t)
	numItems := 10
	items := createNAsset(keeper, ctx, numItems)
	count, found := keeper.GetAddressAssetCount(ctx, items[0].Creator)
	assert.True(t, found)
	assert.Equal(t, uint64(5), count)
	count, found = keeper.GetAddressAssetCount(ctx, items[1].Creator)
	assert.True(t, found)
	assert.Equal(t, uint64(5), count)
}

func TestGetAssetByAddressAndID(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.AssetKeeper(t)
	items := createNAsset(keeper, ctx, 1)
	cid, found := keeper.GetAssetByAddressAndID(ctx, items[0].Creator, 1)
	assert.True(t, found)
	assert.Equal(t, items[0].Cid, cid)
}

func TestGetAssetsByAddress(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.AssetKeeper(t)
	items := createNAsset(keeper, ctx, 10)
	cids, found := keeper.GetAssetsByAddress(ctx, items[0].Creator, nil, nil)
	assert.True(t, found)
	assert.Equal(t, items[0].Cid, cids[0])
	assert.Equal(t, items[4].Cid, cids[2])
	cids, found = keeper.GetAssetsByAddress(ctx, items[1].Creator, nil, nil)
	assert.True(t, found)
	assert.Equal(t, items[1].Cid, cids[0])
	assert.Equal(t, items[5].Cid, cids[2])

	cids, found = keeper.GetAssetsByAddress(ctx, items[0].Creator, util.SerializeUint64(3), nil)
	assert.True(t, found)
	assert.Equal(t, 3, len(cids))
	assert.Equal(t, items[4].Cid, cids[0])
}
