package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/machine module sentinel errors
var (
	ErrMachineNotFound              = errorsmod.Register(ModuleName, 2, "machine not found")
	ErrTrustAnchorNotFound          = errorsmod.Register(ModuleName, 3, "trust anchor not found")
	ErrTrustAnchorAlreadyInUse      = errorsmod.Register(ModuleName, 4, "trust anchor already in use")
	ErrMachineIsNotCreator          = errorsmod.Register(ModuleName, 5, "the machine.address is no the message creator address")
	ErrInvalidKey                   = errorsmod.Register(ModuleName, 6, "invalid key")
	ErrNFTIssuanceFailed            = errorsmod.Register(ModuleName, 7, "an error occurred while issuing the machine NFT")
	ErrMachineTypeUndefined         = errorsmod.Register(ModuleName, 8, "the machine type has to be defined")
	ErrInvalidTrustAnchorKey        = errorsmod.Register(ModuleName, 9, "invalid trust anchor pubkey")
	ErrTrustAnchorAlreadyRegistered = errorsmod.Register(ModuleName, 10, "trust anchor is already registered")
	ErrMachineNFTIssuance           = errorsmod.Register(ModuleName, 11, "the machine NFT could not be issued")
	ErrMachineNFTIssuanceNoOutput   = errorsmod.Register(ModuleName, 12, "the machine NFT issuing process derivated")
	ErrAssetRegistryReqFailure      = errorsmod.Register(ModuleName, 13, "request to asset registry could not be created")
	ErrAssetRegistryReqSending      = errorsmod.Register(ModuleName, 14, "request to asset registry could not be sent")
	ErrAssetRegistryRepsonse        = errorsmod.Register(ModuleName, 15, "request response issue")
)
