package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/dao module sentinel errors
var (
	ErrInvalidMintAddress = errorsmod.Register(ModuleName, 2, "invalid mint address")
)
