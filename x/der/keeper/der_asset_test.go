package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	dertypes "github.com/planetmint/planetmint-go/x/der/types"
	"github.com/stretchr/testify/require"

	dbm "github.com/cometbft/cometbft-db"

	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

func createTestKeeper(t *testing.T) (Keeper, sdk.Context) {
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
	keeper := NewKeeper(cdc, storeKey, memKey, ps, nil, "") // pass nil for MachineKeeper and empty rootDir for this test
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	return *keeper, ctx
}

func TestStoreDerAssetAndLookupDerAsset(t *testing.T) {
	keeper, ctx := createTestKeeper(t)
	asset := dertypes.DER{
		ZigbeeID: "test-zigbee-id",
		// Add other fields as needed for your DER struct
	}

	keeper.StoreDerAsset(ctx, asset)
	result, found := keeper.LookupDerAsset(ctx, "test-zigbee-id")
	require.True(t, found)
	require.Equal(t, asset.ZigbeeID, result.ZigbeeID)
	// Add more assertions for other fields as needed
}

func TestLookupDerAsset_NotFound(t *testing.T) {
	keeper, ctx := createTestKeeper(t)
	_, found := keeper.LookupDerAsset(ctx, "nonexistent-id")
	require.False(t, found)
}
