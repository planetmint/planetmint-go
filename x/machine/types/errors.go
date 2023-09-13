package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/machine module sentinel errors
var (
	ErrMachineNotFound         = errorsmod.Register(ModuleName, 1, "machine not found")
	ErrTrustAnchorNotFound     = errorsmod.Register(ModuleName, 2, "trust anchor not found")
	ErrTrustAnchorAlreadyInUse = errorsmod.Register(ModuleName, 3, "trust anchor already in use")
)
