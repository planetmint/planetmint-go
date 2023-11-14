package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/dao module sentinel errors
var (
	ErrInvalidMintAddress       = errorsmod.Register(ModuleName, 2, "invalid mint address")
	ErrMintFailed               = errorsmod.Register(ModuleName, 3, "minting failed")
	ErrTransferFailed           = errorsmod.Register(ModuleName, 4, "transfer failed")
	ErrInvalidAddress           = errorsmod.Register(ModuleName, 5, "invalid address")
	ErrAlreadyMinted            = errorsmod.Register(ModuleName, 6, "already minted")
	ErrWrongBlockHeight         = errorsmod.Register(ModuleName, 7, "wrong block height")
	ErrReissuanceNotFound       = errorsmod.Register(ModuleName, 8, "reissuance not found")
	ErrInvalidProposer          = errorsmod.Register(ModuleName, 9, "invalid proposer")
	ErrTXAlreadySet             = errorsmod.Register(ModuleName, 10, "tx already set")
	ErrReissuanceProposal       = errorsmod.Register(ModuleName, 11, "invalid reissuance proposal")
	ErrReissuanceFailed         = errorsmod.Register(ModuleName, 12, "reissuance of RDDL failed")
	ErrInvalidChallenge         = errorsmod.Register(ModuleName, 13, "invalid challenge")
	ErrFailedPoPRewardsIssuance = errorsmod.Register(ModuleName, 14, "PoP rewards issuance failed")
)
