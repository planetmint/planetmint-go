package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/machine module sentinel errors
var (
	ErrMachineNotFound         = errorsmod.Register(ModuleName, 2, "machine not found")
	ErrTrustAnchorNotFound     = errorsmod.Register(ModuleName, 3, "trust anchor not found")
	ErrTrustAnchorAlreadyInUse = errorsmod.Register(ModuleName, 4, "trust anchor already in use")
	ErrMachineIsNotCreator     = errorsmod.Register(ModuleName, 5, "the machine.address is no the message creator address")
)
