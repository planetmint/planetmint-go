package keeper

import (
	"testing"

	"github.com/planetmint/planetmint-go/x/machine/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
)

func MachineKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	taIndexStoreKey := sdk.NewKVStoreKey(types.TAIndexKey)
	issuerPlanetmintIndexStoreKey := sdk.NewKVStoreKey(types.IssuerPlanetmintIndexKey)
	issuerLiquidIndexStoreKey := sdk.NewKVStoreKey(types.IssuerLiquidIndexKey)
	trustAnchorStoreKey := sdk.NewKVStoreKey(types.TrustAnchorKey)
	addressStoreKey := sdk.NewKVStoreKey(types.AddressIndexKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
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

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"MachineParams",
	)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		taIndexStoreKey,
		issuerPlanetmintIndexStoreKey,
		issuerLiquidIndexStoreKey,
		trustAnchorStoreKey,
		addressStoreKey,
		memStoreKey,
		paramsSubspace,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	_ = k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
