package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefix(types.ParamsKey))
	if bz == nil {
		return params
	}
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.KeyPrefix(types.ParamsKey), bz)

	return nil
}

func (k Keeper) GetMintAddress(ctx sdk.Context) (mintAddress string) {
	return k.GetParams(ctx).MintAddress
}

func (k Keeper) GetTxGasLimit(ctx sdk.Context) (txGasLimit uint64) {
	return k.GetParams(ctx).TxGasLimit
}

func (k Keeper) GetClaimAddress(ctx sdk.Context) (claimAddress string) {
	return k.GetParams(ctx).ClaimAddress
}

func (k Keeper) GetValidatorPoPReward(ctx sdk.Context) (validatorPoPReward uint64) {
	return k.GetParams(ctx).ValidatorPopReward
}
