package keeper

import (
	"testing"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
)

func CreateTestKeeper(t *testing.T) (Keeper, sdk.Context) {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	storeKey := sdk.NewKVStoreKey("der")
	memKey := storetypes.NewMemoryStoreKey("mem_der")
	cms.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(memKey, storetypes.StoreTypeMemory, nil)
	err := cms.LoadLatestVersion()
	require.NoError(t, err)

	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	ps := paramtypes.NewSubspace(cdc, codec.NewLegacyAmino(), storeKey, memKey, "DerParams")
	keeper := NewKeeper(cdc, storeKey, memKey, ps, nil, "")
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	return *keeper, ctx
}
