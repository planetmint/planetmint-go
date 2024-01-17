package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// SetRedeemClaim set a specific redeemClaim in the store from its index
func (k Keeper) SetRedeemClaim(ctx sdk.Context, redeemClaim types.RedeemClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemClaimKeyPrefix))
	b := k.cdc.MustMarshal(&redeemClaim)
	store.Set(types.RedeemClaimKey(
		redeemClaim.Beneficiary,
		redeemClaim.LiquidTxHash,
	), b)
}

// GetRedeemClaim returns a redeemClaim from its index
func (k Keeper) GetRedeemClaim(
	ctx sdk.Context,
	beneficiary string,
	liquidTxHash string,

) (val types.RedeemClaim, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemClaimKeyPrefix))

	b := store.Get(types.RedeemClaimKey(
		beneficiary,
		liquidTxHash,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRedeemClaim removes a redeemClaim from the store
func (k Keeper) RemoveRedeemClaim(
	ctx sdk.Context,
	beneficiary string,
	liquidTxHash string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemClaimKeyPrefix))
	store.Delete(types.RedeemClaimKey(
		beneficiary,
		liquidTxHash,
	))
}

// GetAllRedeemClaim returns all redeemClaim
func (k Keeper) GetAllRedeemClaim(ctx sdk.Context) (list []types.RedeemClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemClaimKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RedeemClaim
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
