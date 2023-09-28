package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/dao module sentinel errors
var (
	ErrInvalidChallenge = errorsmod.Register(ModuleName, 2, "invalid challenge")
)
