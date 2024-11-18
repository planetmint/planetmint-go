package keeper

import (
	"context"
	"fmt"

	config "github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rddl-network/go-utils/signature"
)

func (k msgServer) AttestMachine(goCtx context.Context, msg *types.MsgAttestMachine) (*types.MsgAttestMachineResponse, error) {
	if err := k.validateMachineAttestation(msg.Machine); err != nil {
		return nil, err
	}

	if err := k.processMachineAttestation(goCtx, msg.Machine); err != nil {
		return nil, err
	}

	return &types.MsgAttestMachineResponse{}, nil
}

func (k msgServer) validateMachineAttestation(machine *types.Machine) error {
	// Validate machine signature
	if err := k.validateMachineSignature(machine); err != nil {
		return err
	}

	// Validate issuer keys
	if err := k.validateIssuerKeys(machine); err != nil {
		return err
	}

	// Validate machine type
	if machine.GetType() == 0 { // 0 == RDDL_MACHINE_UNDEFINED
		return types.ErrMachineTypeUndefined
	}

	return nil
}

func (k msgServer) validateMachineSignature(machine *types.Machine) error {
	isValidSecp256r1, errR1 := signature.ValidateSECP256R1Signature(
		machine.MachineId,
		machine.MachineIdSignature,
		machine.MachineId,
	)

	if errR1 == nil && isValidSecp256r1 {
		return nil
	}

	isValidSecp256k1, errK1 := signature.ValidateSignature(
		machine.MachineId,
		machine.MachineIdSignature,
		machine.MachineId,
	)

	if errK1 == nil && isValidSecp256k1 {
		return nil
	}

	return fmt.Errorf("invalid machine signature: %s, %s", errR1.Error(), errK1.Error())
}

func (k msgServer) validateIssuerKeys(machine *types.Machine) error {
	if !validateExtendedPublicKey(machine.IssuerPlanetmint, config.PlmntNetParams) {
		return errorsmod.Wrap(types.ErrInvalidKey, "planetmint")
	}

	if !validateExtendedPublicKey(machine.IssuerLiquid, config.LiquidNetParams) {
		return errorsmod.Wrap(types.ErrInvalidKey, "liquid")
	}

	return nil
}

func (k msgServer) processMachineAttestation(goCtx context.Context, machine *types.Machine) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	// Process NFT issuance if validator is block proposer
	if util.IsValidatorBlockProposer(ctx, k.rootDir) {
		if err := k.handleNFTIssuance(goCtx, machine, params); err != nil {
			return err
		}
		k.sendInitialFundingTokensToMachine(goCtx, machine.GetAddress(), params)
	} else {
		util.GetAppLogger().Info(ctx, "Not block proposer: skipping Machine NFT issuance")
	}

	// Store machine data
	k.StoreMachine(ctx, *machine)
	k.StoreMachineIndex(ctx, *machine)

	// Store trust anchor
	ta, _, _ := k.GetTrustAnchor(ctx, machine.MachineId)
	if err := k.StoreTrustAnchor(ctx, ta, true); err != nil {
		return err
	}

	return nil
}

func (k msgServer) handleNFTIssuance(goCtx context.Context, machine *types.Machine, params types.Params) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := util.GetAppLogger()
	logger.Info(ctx, "Issuing Machine NFT: "+machine.String())

	err := util.IssueMachineNFT(goCtx, machine,
		params.AssetRegistryScheme,
		params.AssetRegistryDomain,
		params.AssetRegistryPath,
	)

	if err != nil {
		logger.Error(ctx, "Machine NFT issuance failed: "+err.Error())
		return err
	}

	logger.Info(ctx, "Machine NFT issuance successful: "+machine.String())
	return nil
}

func (k msgServer) sendInitialFundingTokensToMachine(goCtx context.Context, machineAddressString string, keeperParams types.Params) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	machineAddress, err := sdk.AccAddressFromBech32(machineAddressString)
	if err != nil {
		util.GetAppLogger().Error(ctx, "error: for provided address "+machineAddress.String())
		return
	}

	logMsg := fmt.Sprintf("transferring %v tokens to address %s", keeperParams.GetDaoMachineFundingAmount(), machineAddress.String())
	util.GetAppLogger().Info(ctx, logMsg)
	util.SendTokens(goCtx, machineAddress, keeperParams.GetDaoMachineFundingAmount(), keeperParams.GetDaoMachineFundingDenom())
}

func validateExtendedPublicKey(issuer string, cfg chaincfg.Params) bool {
	xpubKey, err := hdkeychain.NewKeyFromString(issuer)
	if err != nil {
		return false
	}
	isValidExtendedPublicKey := xpubKey.IsForNet(&cfg)
	return isValidExtendedPublicKey
}
