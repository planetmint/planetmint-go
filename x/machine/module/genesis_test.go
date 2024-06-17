package machine_test

import (
	"testing"

	"github.com/planetmint/planetmint-go/x/machine/types"

	machine "github.com/planetmint/planetmint-go/x/machine/module"

	"github.com/planetmint/planetmint-go/testutil/nullify"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MachineKeeper(t)
	machine.InitGenesis(ctx, k, genesisState)
	got := machine.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
