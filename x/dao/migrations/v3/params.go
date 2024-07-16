package v3

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func MigrateParams(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	paramBytes := store.Get(types.KeyPrefix(types.ParamsKey))

	var params types.Params
	err := cdc.Unmarshal(paramBytes, &params)
	if err != nil {
		return err
	}

	params.ValidatorPopReward = 100000000

	bz, err := cdc.Marshal(&params)
	if err != nil {
		return err
	}

	store.Set(types.KeyPrefix(types.ParamsKey), bz)

	return nil
}
