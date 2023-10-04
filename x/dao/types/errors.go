package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/dao module sentinel errors
var (
	ErrInvalidMintAddress = errorsmod.Register(ModuleName, 2, "invalid mint address")
	ErrMintFailed         = errorsmod.Register(ModuleName, 3, "minting failed")
	ErrTransferFailed     = errorsmod.Register(ModuleName, 4, "transfer failed")
	ErrInvalidAddress     = errorsmod.Register(ModuleName, 5, "invalid address")
)
