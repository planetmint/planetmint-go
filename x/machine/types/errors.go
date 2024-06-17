package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/machine module sentinel errors
var (
	ErrInvalidSigner                = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrMachineNotFound              = sdkerrors.Register(ModuleName, 2, "machine not found")
	ErrTrustAnchorNotFound          = sdkerrors.Register(ModuleName, 3, "trust anchor not found")
	ErrTrustAnchorAlreadyInUse      = sdkerrors.Register(ModuleName, 4, "trust anchor already in use")
	ErrMachineIsNotCreator          = sdkerrors.Register(ModuleName, 5, "the machine.address is not the message creator address")
	ErrInvalidKey                   = sdkerrors.Register(ModuleName, 6, "invalid key")
	ErrNFTIssuanceFailed            = sdkerrors.Register(ModuleName, 7, "an error occurred while issuing the machine NFT")
	ErrMachineTypeUndefined         = sdkerrors.Register(ModuleName, 8, "the machine type has to be defined")
	ErrInvalidTrustAnchorKey        = sdkerrors.Register(ModuleName, 9, "invalid trust anchor pubkey")
	ErrTrustAnchorAlreadyRegistered = sdkerrors.Register(ModuleName, 10, "trust anchor is already registered")
	ErrMachineNFTIssuance           = sdkerrors.Register(ModuleName, 11, "the machine NFT could not be issued")
	ErrAssetRegistryReqFailure      = sdkerrors.Register(ModuleName, 13, "request to asset registry could not be created")
	ErrAssetRegistryReqSending      = sdkerrors.Register(ModuleName, 14, "request to asset registry could not be sent")
	ErrAssetRegistryRepsonse        = sdkerrors.Register(ModuleName, 15, "request response issue")
	ErrInvalidAddress               = sdkerrors.Register(ModuleName, 16, "invalid address")
	ErrTransferFailed               = sdkerrors.Register(ModuleName, 17, "transfer failed")
)
