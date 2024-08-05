package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/planetmint/planetmint-go/x/machine/types"
)

type (
	Keeper struct {
		cdc                           codec.BinaryCodec
		storeKey                      storetypes.StoreKey
		taIndexStoreKey               storetypes.StoreKey
		issuerPlanetmintIndexStoreKey storetypes.StoreKey
		issuerLiquidIndexStoreKey     storetypes.StoreKey
		taStoreKey                    storetypes.StoreKey
		addressIndexStoreKey          storetypes.StoreKey
		memKey                        storetypes.StoreKey
		paramstore                    paramtypes.Subspace
		authority                     string
		rootDir                       string
		bankKeeper                    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	indexStoreKey,
	issuerPlanetmintIndexStoreKey,
	issuerLiquidIndexStoreKey,
	taStoreKey,
	addressIndexStoreKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	authority string,
	rootDir string,
	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:                           cdc,
		storeKey:                      storeKey,
		taIndexStoreKey:               indexStoreKey,
		issuerPlanetmintIndexStoreKey: issuerPlanetmintIndexStoreKey,
		issuerLiquidIndexStoreKey:     issuerLiquidIndexStoreKey,
		taStoreKey:                    taStoreKey,
		addressIndexStoreKey:          addressIndexStoreKey,
		memKey:                        memKey,
		paramstore:                    ps,
		authority:                     authority,
		rootDir:                       rootDir,
		bankKeeper:                    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
