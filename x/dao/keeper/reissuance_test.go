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

func createNReissuance(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Reissuance {
	items := make([]types.Reissuance, n)
	for i := range items {
		items[i].BlockHeight = uint64(i)
		items[i].Proposer = fmt.Sprintf("proposer_%v", i)
		items[i].Rawtx = fmt.Sprintf("rawtransaction_%v", i)
		items[i].TxId = ""
		keeper.StoreReissuance(ctx, items[i])
	}
	return items
}

func TestGetResponse(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNChallenge(keeper, ctx, 10)
	for _, item := range items {
		challenge, found := keeper.GetChallenge(ctx, item.Height)
		assert.True(t, found)
		assert.Equal(t, item, challenge)
	}
}
