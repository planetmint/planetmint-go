package keeper

import (
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreMachine(ctx sdk.Context, machine types.Machine) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MachineKey))
	appendValue := k.cdc.MustMarshal(&machine)
	store.Set(util.SerializeString(machine.IssuerPlanetmint), appendValue)
}

func (k Keeper) GetMachine(ctx sdk.Context, index types.MachineIndex) (val types.Machine, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MachineKey))
	machine := store.Get(util.SerializeString(index.IssuerPlanetmint))

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
	addressIndexStore := prefix.NewStore(ctx.KVStore(k.addressIndexStoreKey), types.KeyPrefix(types.AddressIndexKey))

	index := types.MachineIndex{
		MachineId:        machine.MachineId,
		IssuerPlanetmint: machine.IssuerPlanetmint,
		IssuerLiquid:     machine.IssuerLiquid,
		Address:          machine.Address,
	}

	machineIDIndexKey := util.SerializeString(machine.MachineId)
	issuerPlanetmintIndexKey := util.SerializeString(machine.IssuerPlanetmint)
	issuerLiquidIndexKey := util.SerializeString(machine.IssuerLiquid)
	addressIndexKey := util.SerializeString(machine.Address)

	indexAppendValue := k.cdc.MustMarshal(&index)
	taIndexStore.Set(machineIDIndexKey, indexAppendValue)
	issuerPlanetmintIndexStore.Set(issuerPlanetmintIndexKey, indexAppendValue)
	issuerLiquidIndexStore.Set(issuerLiquidIndexKey, indexAppendValue)
	addressIndexStore.Set(addressIndexKey, indexAppendValue)
}

func (k Keeper) GetMachineIndexByPubKey(ctx sdk.Context, pubKey string) (val types.MachineIndex, found bool) {
	taIndexStore := prefix.NewStore(ctx.KVStore(k.taIndexStoreKey), types.KeyPrefix(types.TAIndexKey))
	issuerPlanetmintIndexStore := prefix.NewStore(ctx.KVStore(k.issuerPlanetmintIndexStoreKey), types.KeyPrefix(types.IssuerPlanetmintIndexKey))
	issuerLiquidIndexStore := prefix.NewStore(ctx.KVStore(k.issuerLiquidIndexStoreKey), types.KeyPrefix(types.IssuerLiquidIndexKey))

	keyBytes := util.SerializeString(pubKey)

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

func (k Keeper) GetMachineIndexByAddress(ctx sdk.Context, address string) (val types.MachineIndex, found bool) {
	addressIndexStore := prefix.NewStore(ctx.KVStore(k.addressIndexStoreKey), types.KeyPrefix(types.AddressIndexKey))

	keyBytes := util.SerializeString(address)

	adIndex := addressIndexStore.Get(keyBytes)
	if adIndex != nil {
		if err := k.cdc.Unmarshal(adIndex, &val); err != nil {
			return val, false
		}
		return val, true
	}

	return val, false
}
