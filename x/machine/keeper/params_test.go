package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/planetmint/planetmint-go/x/machine/types"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MachineKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
