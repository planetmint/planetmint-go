package keeper

import (
	"testing"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/golang/mock/gomock"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/require"

	daotestutil "github.com/planetmint/planetmint-go/x/dao/testutil"
)

func DaoKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	cfg := config.GetConfig()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	challengeStoreKey := storetypes.NewMemoryStoreKey(types.ChallengeKey)
	mintRequestHashStoreKey := storetypes.NewMemoryStoreKey(types.MintRequestHashKey)
	mintRequestAddressStoreKey := storetypes.NewMemoryStoreKey(types.MintRequestAddressKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(challengeStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(mintRequestHashStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(mintRequestAddressStoreKey, storetypes.StoreTypeIAVL, db)

	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"DaoParams",
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	ctrl := gomock.NewController(t)
	bk := daotestutil.NewMockBankKeeper(ctrl)

	amt := sdk.NewCoins(sdk.NewCoin(cfg.TokenDenom, sdk.NewIntFromUint64(1000)))
	beneficiaryAddr, _ := sdk.AccAddressFromBech32(sample.ConstBech32Addr)

	bk.EXPECT().MintCoins(ctx, types.ModuleName, amt).Return(nil).AnyTimes()
	bk.EXPECT().SendCoinsFromModuleToAccount(ctx, types.ModuleName, beneficiaryAddr, amt).Return(nil).AnyTimes()

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		challengeStoreKey,
		mintRequestHashStoreKey,
		mintRequestAddressStoreKey,
		paramsSubspace,
		bk,
		nil,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// Initialize params
	err := k.SetParams(ctx, types.DefaultParams())
	if err != nil {
		panic(err)
	}

	return k, ctx
}
