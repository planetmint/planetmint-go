package keeper

import (
	"planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreTrustAnchor(ctx sdk.Context, ta types.TrustAnchor, activated bool) {
	store := prefix.NewStore(ctx.KVStore(k.taStoreKey), types.KeyPrefix(types.TrustAnchorKey))
	appendValue := k.cdc.MustMarshal(&ta)
	store.Set(GetTrustAnchorBytes(ta.Pubkey), appendValue)
}

func (k Keeper) GetTrustAnchor(ctx sdk.Context, pubKey string) (val types.TrustAnchor, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.taStoreKey), types.KeyPrefix(types.TrustAnchorKey))
	trustAnchor := store.Get(GetTrustAnchorBytes(pubKey))

	if trustAnchor == nil {
		return val, false
	}
	if err := k.cdc.Unmarshal(trustAnchor, &val); err != nil {
		return val, false
	}
	return val, true
}

func GetTrustAnchorBytes(pubKey string) []byte {
	return []byte(pubKey)
}
