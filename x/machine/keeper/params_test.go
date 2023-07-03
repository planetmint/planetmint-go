package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "planetmint-go/testutil/keeper"
	"planetmint-go/x/machine/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MachineKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
