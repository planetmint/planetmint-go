package keeper_test

import (
	"fmt"
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func createNMintRequests(keeper *keeper.Keeper, ctx sdk.Context, beneficiary string, n int) []types.MintRequest {
	items := make([]types.MintRequest, n)
	for i := range items {
		items[i].Amount = uint64(i)
		items[i].LiquidTxHash = fmt.Sprintf("hash%v", i)
		items[i].Beneficiary = beneficiary
		keeper.StoreMintRequest(ctx, items[i])
	}
	return items
}

func TestGetMintRequestByHash(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNMintRequests(keeper, ctx, "beneficiary", 10)
	for _, item := range items {
		mintRequest, found := keeper.GetMintRequestByHash(ctx, item.LiquidTxHash)
		assert.True(t, found)
		assert.Equal(t, item, mintRequest)
	}
}

func TestGetMintRequestByAddress(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNMintRequests(keeper, ctx, "beneficiary", 10)
	mintRequests, found := keeper.GetMintRequestsByAddress(ctx, "beneficiary")
	assert.True(t, found)
	assert.Equal(t, items[0], *mintRequests.Requests[0])
	assert.Equal(t, items[1], *mintRequests.Requests[1])
	assert.Equal(t, items[2], *mintRequests.Requests[2])
	assert.Equal(t, items[3], *mintRequests.Requests[3])
	assert.Equal(t, items[4], *mintRequests.Requests[4])
	assert.Equal(t, items[5], *mintRequests.Requests[5])
	assert.Equal(t, items[6], *mintRequests.Requests[6])
	assert.Equal(t, items[7], *mintRequests.Requests[7])
	assert.Equal(t, items[8], *mintRequests.Requests[8])
	assert.Equal(t, items[9], *mintRequests.Requests[9])
}
