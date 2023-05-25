package keeper

import (
	"planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreMachine(ctx sdk.Context, machine types.Machine) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MachineKey))
	appendValue := k.cdc.MustMarshal(&machine)
	store.Set(GetMachineBytes(machine.IssuerPlanetmint), appendValue)
}

func (k Keeper) GetMachine(ctx sdk.Context, pubKey string) (val types.Machine, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MachineKey))
	machine := store.Get(GetMachineBytes(pubKey))
	if machine == nil {
		return val, false
	}
	k.cdc.Unmarshal(machine, &val)
	return val, true
}

func GetMachineBytes(pubKey string) []byte {
	return []byte(pubKey)
}
