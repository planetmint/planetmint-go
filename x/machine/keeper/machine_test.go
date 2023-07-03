package keeper_test

import (
	"fmt"
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
		items[i].MachineId = fmt.Sprintf("machineId%v", i)
		items[i].IssuerPlanetmint = fmt.Sprintf("issuerPlanetmint%v", i)
		items[i].IssuerLiquid = fmt.Sprintf("issuerLiquid%v", i)
		keeper.StoreMachine(ctx, items[i])
		keeper.StoreMachineIndex(ctx, items[i])
	}
	return items
}

func TestGetMachine(t *testing.T) {
	keeper, ctx := keepertest.MachineKeeper(t)
	items := createNMachine(keeper, ctx, 10)
	for _, item := range items {
		index := types.MachineIndex{
			MachineId:        item.MachineId,
			IssuerPlanetmint: item.IssuerPlanetmint,
			IssuerLiquid:     item.IssuerLiquid,
		}
		machineById, found := keeper.GetMachine(ctx, index)
		assert.True(t, found)
		assert.Equal(t, item, machineById)
	}
}

func TestGetMachineIndex(t *testing.T) {
	keeper, ctx := keepertest.MachineKeeper(t)
	items := createNMachine(keeper, ctx, 10)
	for _, item := range items {
		expectedIndex := types.MachineIndex{
			MachineId:        item.MachineId,
			IssuerPlanetmint: item.IssuerPlanetmint,
			IssuerLiquid:     item.IssuerLiquid,
		}
		index, found := keeper.GetMachineIndex(ctx, item.MachineId)
		assert.True(t, found)
		assert.Equal(t, expectedIndex, index)
	}
}
