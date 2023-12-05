package keeper_test

import (
	"fmt"
	"testing"

	"github.com/planetmint/planetmint-go/config"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// this method returns a range of challenges, each with a blockheight * PopEpochs.
// be aware: the first element start with 1 instead of 0
func createNChallenge(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Challenge {
	items := make([]types.Challenge, n)
	for i := range items {
		blockHeight := (i + 1) * config.GetConfig().PopEpochs
		items[i].Height = int64(blockHeight)
		items[i].Initiator = fmt.Sprintf("initiator%v", blockHeight)
		items[i].Challenger = fmt.Sprintf("challenger%v", blockHeight)
		items[i].Challengee = fmt.Sprintf("challengee%v", blockHeight)
		items[i].Success = false
		items[i].Finished = false
		keeper.StoreChallenge(ctx, items[i])
	}
	return items
}

func TestGetChallenge(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNChallenge(keeper, ctx, 10)
	for _, item := range items {
		challenge, found := keeper.LookupChallenge(ctx, item.Height)
		assert.True(t, found)
		assert.Equal(t, item, challenge)
	}
}

func TestGetChallengeRange(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.DaoKeeper(t)
	createNChallenge(keeper, ctx, 10)
	challenges, err := keeper.GetChallengeRange(ctx, int64((0+1)*config.GetConfig().PopEpochs), int64((9+1)*config.GetConfig().PopEpochs))
	assert.NoError(t, err)
	assert.Equal(t, 10, len(challenges))
}
