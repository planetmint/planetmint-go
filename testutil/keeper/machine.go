package keeper

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/planetmint/planetmint-go/x/machine/keeper"
)

func MachineKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	taIndexStoreKey := storetypes.NewKVStoreKey(types.TAIndexKey)
	issuerPlanetmintIndexStoreKey := storetypes.NewKVStoreKey(types.IssuerPlanetmintIndexKey)
	issuerLiquidIndexStoreKey := storetypes.NewKVStoreKey(types.IssuerLiquidIndexKey)
	trustAnchorStoreKey := storetypes.NewKVStoreKey(types.TrustAnchorKey)
	addressStoreKey := storetypes.NewKVStoreKey(types.AddressIndexKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(taIndexStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(issuerPlanetmintIndexStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(issuerLiquidIndexStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(trustAnchorStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(addressStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	k := keeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
		"",
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	if err := k.SetParams(ctx, types.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx
}
