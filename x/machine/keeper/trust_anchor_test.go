package keeper_test

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"

	"github.com/planetmint/planetmint-go/x/machine/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func createNTrustAnchor(t *testing.T, keeper keeper.Keeper, ctx sdk.Context, n int) []types.TrustAnchor {
	items := make([]types.TrustAnchor, n)
	for i := range items {
		pk := fmt.Sprintf("pubkey%v", i)
		if i%2 == 1 {
			pk = strings.ToUpper(pk)
		}

		items[i].Pubkey = hex.EncodeToString([]byte(pk))
		var activated bool
		if i%2 == 1 {
			activated = true
		} else {
			activated = false
		}
		err := keeper.StoreTrustAnchor(ctx, items[i], activated)
		assert.False(t, (err != nil))
	}
	return items
}

func TestGetTrustAnchor(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.MachineKeeper(t)
	items := createNTrustAnchor(t, keeper, ctx, 10)
	for i, item := range items {
		ta, activated, found := keeper.GetTrustAnchor(ctx, item.Pubkey)
		assert.True(t, found)
		assert.Equal(t, item, ta)
		if i%2 == 1 {
			assert.True(t, activated)
		} else {
			assert.False(t, activated)
		}
	}
}

func TestUpdateTrustAnchor(t *testing.T) {
	t.Parallel()
	keeper, ctx := keepertest.MachineKeeper(t)
	items := createNTrustAnchor(t, keeper, ctx, 10)
	for _, item := range items {
		ta, activated, _ := keeper.GetTrustAnchor(ctx, item.Pubkey)
		if !activated {
			err := keeper.StoreTrustAnchor(ctx, ta, true)
			assert.False(t, (err != nil))
		}
	}

	for _, item := range items {
		_, activated, _ := keeper.GetTrustAnchor(ctx, item.Pubkey)
		assert.True(t, activated)
	}
}
