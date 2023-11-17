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

func createNChallenge(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Challenge {
	items := make([]types.Challenge, n)
	for i := range items {
		items[i].Height = uint64(i)
		items[i].Initiator = fmt.Sprintf("initiator%v", i)
		items[i].Challenger = fmt.Sprintf("challenger%v", i)
		items[i].Challengee = fmt.Sprintf("challengee%v", i)
		items[i].Success = true
		items[i].Description = fmt.Sprintf("expected %v got %v", i, i)
		keeper.StoreChallenge(ctx, items[i])
	}
	return items
}

func TestGetChallenge(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNChallenge(keeper, ctx, 10)
	for _, item := range items {
		challenge, found := keeper.GetChallenge(ctx, item.Height)
		assert.True(t, found)
		assert.Equal(t, item, challenge)
	}
}
