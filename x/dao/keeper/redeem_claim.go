package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// GetBeneficiaryRedeemClaimCount get the total number of RedeemClaim for a beneficiary
func (k Keeper) GetBeneficiaryRedeemClaimCount(ctx sdk.Context, beneficiary string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemClaimBeneficiaryCountKeyPrefix))
	byteKey := types.KeyPrefix(beneficiary)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetBeneficiaryRedeemClaimCount set the total number of RedeemClaim for beneficiary
func (k Keeper) SetBeneficiaryRedeemClaimCount(ctx sdk.Context, beneficiary string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemClaimBeneficiaryCountKeyPrefix))
	byteKey := types.KeyPrefix(beneficiary)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetRedeemClaim set a specific redeemClaim in the store from its index
func (k Keeper) SetRedeemClaim(ctx sdk.Context, redeemClaim types.RedeemClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemClaimKeyPrefix))
	b := k.cdc.MustMarshal(&redeemClaim)
	store.Set(types.RedeemClaimKey(
		redeemClaim.Beneficiary,
		redeemClaim.Id,
	), b)
}

// CreateRedeemClaim creates a specific redeemClaim in the store from its index
func (k Keeper) CreateNewRedeemClaim(ctx sdk.Context, redeemClaim types.RedeemClaim) uint64 {
	count := k.GetBeneficiaryRedeemClaimCount(ctx, redeemClaim.Beneficiary)

	redeemClaim.Id = count
	k.SetRedeemClaim(ctx, redeemClaim)

	// Update BeneficiaryRedeemCount
	k.SetBeneficiaryRedeemClaimCount(ctx, redeemClaim.Beneficiary, count+1)

	return count
}

// GetRedeemClaim returns a redeemClaim from its index
func (k Keeper) GetRedeemClaim(
	ctx sdk.Context,
	beneficiary string,
	id uint64,

) (val types.RedeemClaim, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RedeemClaimKeyPrefix))

	b := store.Get(types.RedeemClaimKey(
		beneficiary,
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
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
