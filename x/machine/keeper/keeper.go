package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/crgimenes/go-osc"

	"planetmint-go/x/machine/types"
)

type (
	Keeper struct {
		cdc                           codec.BinaryCodec
		storeKey                      storetypes.StoreKey
		taIndexStoreKey               storetypes.StoreKey
		issuerPlanetmintIndexStoreKey storetypes.StoreKey
		issuerLiquidIndexStoreKey     storetypes.StoreKey
		memKey                        storetypes.StoreKey
		paramstore                    paramtypes.Subspace
		oscClient                     osc.Client
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	indexStoreKey,
	issuerPlanetmintIndexStoreKey,
	issuerLiquidIndexStoreKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	oscClient osc.Client,
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
		memKey:                        memKey,
		paramstore:                    ps,
		oscClient:                     oscClient,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
