package keeper_test

import (
	"testing"

	keepertest "planetmint-go/testutil/keeper"
	"planetmint-go/x/machine/keeper"
	"planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func createNMachine(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Machine {
	items := make([]types.Machine, n)
	for i := range items {
		// fill out machine
		items[i].IssuerPlanetmint = "asd"
		items[i].IssuerLiquid = "dsa"
		keeper.StoreMachine(ctx, items[i])
	}
	return items
}

func TestGetMachine(t *testing.T) {
	keeper, ctx := keepertest.MachineKeeper(t)
	items := createNMachine(keeper, ctx, 10)
	for _, item := range items {
		machineById, found := keeper.GetMachine(ctx, item.IssuerPlanetmint)
		assert.True(t, found)
		assert.Equal(t, item, machineById)
	}
}
