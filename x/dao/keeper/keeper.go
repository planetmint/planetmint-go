package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/planetmint/planetmint-go/monitor"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

type (
	Keeper struct {
		cdc                   codec.BinaryCodec
		storeKey              storetypes.StoreKey
		memKey                storetypes.StoreKey
		challengeKey          storetypes.StoreKey
		mintRequestHashKey    storetypes.StoreKey
		mintRequestAddressKey storetypes.StoreKey
		accountKeeperKey      storetypes.StoreKey
		paramstore            paramtypes.Subspace

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		machineKeeper types.MachineKeeper
		authority     string
		RootDir       string
		MqttMonitor   *monitor.MqttMonitor
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	challengeKey storetypes.StoreKey,
	mintRequestHashKey storetypes.StoreKey,
	mintRequestAddressKey storetypes.StoreKey,
	accountKeeperKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	machineKeeper types.MachineKeeper,
	authority string,
	rootDir string,
	mqttMonitor *monitor.MqttMonitor,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:                   cdc,
		storeKey:              storeKey,
		memKey:                memKey,
		challengeKey:          challengeKey,
		mintRequestHashKey:    mintRequestHashKey,
		mintRequestAddressKey: mintRequestAddressKey,
		accountKeeperKey:      accountKeeperKey,
		paramstore:            ps,

		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		machineKeeper: machineKeeper,
		authority:     authority,
		RootDir:       rootDir,
		MqttMonitor:   mqttMonitor,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
