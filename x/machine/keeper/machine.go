package keeper

import (
	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreMachine(ctx sdk.Context, machine types.Machine) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MachineKey))
	appendValue := k.cdc.MustMarshal(&machine)
	store.Set(GetMachineBytes(machine.IssuerPlanetmint), appendValue)
}

func (k Keeper) GetMachine(ctx sdk.Context, index types.MachineIndex) (val types.Machine, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MachineKey))
	machine := store.Get(GetMachineBytes(index.IssuerPlanetmint))

	if machine == nil {
		return val, false
	}
	if err := k.cdc.Unmarshal(machine, &val); err != nil {
		return val, false
	}
	return val, true
}

func (k Keeper) StoreMachineIndex(ctx sdk.Context, machine types.Machine) {
	taIndexStore := prefix.NewStore(ctx.KVStore(k.taIndexStoreKey), types.KeyPrefix(types.TAIndexKey))
	issuerPlanetmintIndexStore := prefix.NewStore(ctx.KVStore(k.issuerPlanetmintIndexStoreKey), types.KeyPrefix(types.IssuerPlanetmintIndexKey))
	issuerLiquidIndexStore := prefix.NewStore(ctx.KVStore(k.issuerLiquidIndexStoreKey), types.KeyPrefix(types.IssuerLiquidIndexKey))

	index := types.MachineIndex{
		MachineId:        machine.MachineId,
		IssuerPlanetmint: machine.IssuerPlanetmint,
		IssuerLiquid:     machine.IssuerLiquid,
	}

	machineIdIndexKey := GetMachineBytes(machine.MachineId)
	issuerPlanetmintIndexKey := GetMachineBytes(machine.IssuerPlanetmint)
	issuerLiquidIndexKey := GetMachineBytes(machine.IssuerLiquid)
	indexAppendValue := k.cdc.MustMarshal(&index)
	taIndexStore.Set(machineIdIndexKey, indexAppendValue)
	issuerPlanetmintIndexStore.Set(issuerPlanetmintIndexKey, indexAppendValue)
	issuerLiquidIndexStore.Set(issuerLiquidIndexKey, indexAppendValue)
}

func (k Keeper) GetMachineIndex(ctx sdk.Context, pubKey string) (val types.MachineIndex, found bool) {
	taIndexStore := prefix.NewStore(ctx.KVStore(k.taIndexStoreKey), types.KeyPrefix(types.TAIndexKey))
	issuerPlanetmintIndexStore := prefix.NewStore(ctx.KVStore(k.issuerPlanetmintIndexStoreKey), types.KeyPrefix(types.IssuerPlanetmintIndexKey))
	issuerLiquidIndexStore := prefix.NewStore(ctx.KVStore(k.issuerLiquidIndexStoreKey), types.KeyPrefix(types.IssuerLiquidIndexKey))

	keyBytes := GetMachineBytes(pubKey)

	taIndex := taIndexStore.Get(keyBytes)
	if taIndex != nil {
		if err := k.cdc.Unmarshal(taIndex, &val); err != nil {
			return val, false
		}
		return val, true
	}

	ipIndex := issuerPlanetmintIndexStore.Get(keyBytes)
	if ipIndex != nil {
		if err := k.cdc.Unmarshal(ipIndex, &val); err != nil {
			return val, false
		}
		return val, true
	}

	ilIndex := issuerLiquidIndexStore.Get(keyBytes)
	if ilIndex != nil {
		if err := k.cdc.Unmarshal(ilIndex, &val); err != nil {
			return val, false
		}
		return val, true
	}

	return val, false
}

func GetMachineBytes(pubKey string) []byte {
	return []byte(pubKey)
}
