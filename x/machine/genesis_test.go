package machine_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "planetmint-go/testutil/keeper"
	"planetmint-go/testutil/nullify"
	"planetmint-go/x/machine"
	"planetmint-go/x/machine/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MachineKeeper(t)
	machine.InitGenesis(ctx, *k, genesisState)
	got := machine.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
