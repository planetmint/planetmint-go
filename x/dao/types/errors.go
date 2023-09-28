package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/dao module sentinel errors
var (
	ErrInvalidChallenge         = errorsmod.Register(ModuleName, 2, "invalid challenge")
	ErrFailedPoPRewardsIssuance = errorsmod.Register(ModuleName, 3, "PoP rewards issuance failed")
)
