package v2

// import (
// 	storetypes "cosmossdk.io/store/types"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/planetmint/planetmint-go/x/machine/types"
// )

// // MigrateStore migrates the x/machine module state from the consensus version
// // 1 to version 2. Specifically, it takes the default params and stores them
// // directly into the x/machine module state.
// func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
// 	store := ctx.KVStore(storeKey)
// 	params := types.DefaultParams()

// 	bz, err := cdc.Marshal(&params)
// 	if err != nil {
// 		return err
// 	}

// 	store.Set(types.KeyPrefix(types.ParamsKey), bz)

// 	return nil
// }
