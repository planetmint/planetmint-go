package keeper_test

import (
	"testing"

	testkeeper "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	t.Parallel()
	k, ctx := testkeeper.MachineKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
