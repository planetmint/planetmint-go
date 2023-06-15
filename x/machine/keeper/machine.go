package keeper

import (
	"planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreMachine(ctx sdk.Context, machine types.Machine) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MachineKey))
	appendValue := k.cdc.MustMarshal(&machine)
	k.StoreMachineIndex(ctx, machine)
	store.Set(GetMachineBytes(machine.IssuerPlanetmint), appendValue)
}

func (k Keeper) GetMachine(ctx sdk.Context, pubKey string) (val types.Machine, found bool) {
	index, found := k.GetMachineIndex(ctx, pubKey)
	if !found {
		return val, false
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MachineKey))
	machine := store.Get(GetMachineBytes(index.IssuerPlanetmint))

	if machine == nil {
		return val, false
	}
	k.cdc.Unmarshal(machine, &val)
	return val, true
}

func (k Keeper) StoreMachineIndex(ctx sdk.Context, machine types.Machine) {
	indexStore := prefix.NewStore(ctx.KVStore(k.indexStoreKey), types.KeyPrefix(types.IndexKey))
	index := types.MachineIndex{
		MachineId:        machine.MachineId,
		IssuerPlanetmint: machine.IssuerPlanetmint,
		IssuerLiquid:     machine.IssuerPlanetmint,
	}

	machineIdIndexKey := GetMachineBytes(machine.MachineId)
	issuerPlanetmintIndexKey := GetMachineBytes(machine.IssuerPlanetmint)
	issuerLiquidIndexKey := GetMachineBytes(machine.IssuerLiquid)
	indexAppendValue := k.cdc.MustMarshal(&index)
	indexStore.Set(machineIdIndexKey, indexAppendValue)
	indexStore.Set(issuerPlanetmintIndexKey, indexAppendValue)
	indexStore.Set(issuerLiquidIndexKey, indexAppendValue)
}

func (k Keeper) GetMachineIndex(ctx sdk.Context, pubKey string) (val types.MachineIndex, found bool) {
	indexStore := prefix.NewStore(ctx.KVStore(k.indexStoreKey), types.KeyPrefix(types.IndexKey))
	index := indexStore.Get(GetMachineBytes(pubKey))

	if index == nil {
		return val, false
	}

	k.cdc.Unmarshal(index, &val)
	return val, true
}

func GetMachineBytes(pubKey string) []byte {
	return []byte(pubKey)
}
