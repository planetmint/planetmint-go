package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/stretchr/testify/assert"
)

func TestActivatedTrustAnchorCount(t *testing.T) {
	keeper, ctx := keepertest.MachineKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	createNTrustAnchor(t, keeper, ctx, 100)
	response, err := keeper.ActivatedTrustAnchorCount(wctx, &types.QueryActivatedTrustAnchorCountRequest{})
	assert.NoError(t, err)
	assert.Equal(t, uint64(50), response.Count)
}
