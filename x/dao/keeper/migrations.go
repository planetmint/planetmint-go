package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/planetmint/planetmint-go/x/dao/migrations/v2"
	v3 "github.com/planetmint/planetmint-go/x/dao/migrations/v3"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateParams(ctx, m.keeper.storeKey, m.keeper.cdc)
}
