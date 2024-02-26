package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// MigrateStore migrates the x/dao module state from the consensus version 1 to
// version 2. Specifically, it takes the default params and stores them
// directly into the x/dao module state.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	params := types.DefaultParams()

	bz, err := cdc.Marshal(&params)
	if err != nil {
		return err
	}

	store.Set(types.KeyPrefix(types.ParamsKey), bz)

	return nil
}
