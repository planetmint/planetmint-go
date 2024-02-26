package keeper

import (
	"testing"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/testutil/moduleobject"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/x/asset/keeper"
	"github.com/planetmint/planetmint-go/x/asset/types"

	assettestutils "github.com/planetmint/planetmint-go/x/asset/testutil"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func AssetKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"AssetParams",
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	ctrl := gomock.NewController(t)
	mk := assettestutils.NewMockMachineKeeper(ctrl)
	sk, pk := sample.KeyPair()
	_, ppk := moduleobject.ExtendedKeyPair(config.PlmntNetParams)
	_, lpk := moduleobject.ExtendedKeyPair(config.LiquidNetParams)

	id := moduleobject.MachineIndex(pk, ppk, lpk)
	mk.EXPECT().GetMachineIndexByPubKey(ctx, pk).Return(id, true).AnyTimes()
	mk.EXPECT().GetMachineIndexByPubKey(ctx, ppk).Return(id, true).AnyTimes()
	mk.EXPECT().GetMachineIndexByPubKey(ctx, sk).Return(id, false).AnyTimes()
	mk.EXPECT().GetMachine(ctx, id).Return(moduleobject.Machine(pk, pk, sk, ""), true).AnyTimes()
	mk.EXPECT().GetMachine(ctx, sk).Return(moduleobject.Machine(pk, pk, sk, ""), false).AnyTimes()

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		mk,
	)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
