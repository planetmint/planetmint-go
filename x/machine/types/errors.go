package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/machine module sentinel errors
var (
	ErrMachineNotFound = errorsmod.Register(ModuleName, 1, "machine not found")
)
