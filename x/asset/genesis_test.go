package asset_test

import (
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/nullify"
	"github.com/planetmint/planetmint-go/x/asset"
	"github.com/planetmint/planetmint-go/x/asset/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	t.Parallel()
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AssetKeeper(t)
	asset.InitGenesis(ctx, *k, genesisState)
	got := asset.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
