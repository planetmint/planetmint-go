package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k Keeper) StorePoPDistribution(ctx sdk.Context, popdistribution types.PoPDistribution) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	appendValue := k.cdc.MustMarshal(&popdistribution)
	store.Set(GetKeyBytes("PoPDistribution"), appendValue)
}

func (k Keeper) LookupPoPDistribution(ctx sdk.Context) (val types.PoPDistribution, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	popdistribution := store.Get(GetKeyBytes("PoPDistribution"))
	if popdistribution == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(popdistribution, &val)
	return val, true
}

func GetKeyBytes(value string) []byte {
	return []byte(value)
}
