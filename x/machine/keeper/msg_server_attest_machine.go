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
)

func (k msgServer) AttestMachine(goCtx context.Context, msg *types.MsgAttestMachine) (*types.MsgAttestMachineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// the ante handler verifies that the MachineID exists. Additional result checks got moved to the ante-handler
	// and removed from here due to inconsistency or checking the same thing over and over again.
	ta, _, _ := k.GetTrustAnchor(ctx, msg.Machine.MachineId)

	isValidMachineID, err := util.ValidateSignature(msg.Machine.MachineId, msg.Machine.MachineIdSignature, msg.Machine.MachineId)
	if !isValidMachineID {
		return nil, err
	}

	isValidIssuerPlanetmint := validateExtendedPublicKey(msg.Machine.IssuerPlanetmint, config.PlmntNetParams)
	if !isValidIssuerPlanetmint {
		return nil, errorsmod.Wrap(types.ErrInvalidKey, "planetmint")
	}
	isValidIssuerLiquid := validateExtendedPublicKey(msg.Machine.IssuerLiquid, config.LiquidNetParams)
	if !isValidIssuerLiquid {
		return nil, errorsmod.Wrap(types.ErrInvalidKey, "liquid")
	}

	if msg.Machine.GetType() == 0 { // 0 == RDDL_MACHINE_UNDEFINED
		return nil, types.ErrMachineTypeUndefined
	}
	params := k.GetParams(ctx)
	if util.IsValidatorBlockProposer(ctx, ctx.BlockHeader().ProposerAddress, k.rootDir) {
		util.GetAppLogger().Info(ctx, "Issuing Machine NFT: "+msg.Machine.String())
		scheme := params.AssetRegistryScheme
		domain := params.AssetRegistryDomain
		path := params.AssetRegistryPath
		localErr := util.IssueMachineNFT(goCtx, msg.Machine, scheme, domain, path)
		if localErr != nil {
			util.GetAppLogger().Error(ctx, "Machine NFT issuance failed : "+localErr.Error())
		} else {
			util.GetAppLogger().Info(ctx, "Machine NFT issuance successful: "+msg.Machine.String())
		}

		k.sendInitialFundingTokensToMachine(goCtx, msg.GetMachine().GetAddress(), params)
	} else {
		util.GetAppLogger().Info(ctx, "Not block proposer: skipping Machine NFT issuance")
	}

	k.StoreMachine(ctx, *msg.Machine)
	k.StoreMachineIndex(ctx, *msg.Machine)
	err = k.StoreTrustAnchor(ctx, ta, true)
	if err != nil {
		return nil, err
	}

	return &types.MsgAttestMachineResponse{}, err
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
	util.SendPLMNTTokens(goCtx, machineAddress, keeperParams.GetDaoMachineFundingAmount(), keeperParams.DaoMachineFundingDenom)
}

func validateExtendedPublicKey(issuer string, cfg chaincfg.Params) bool {
	xpubKey, err := hdkeychain.NewKeyFromString(issuer)
	if err != nil {
		return false
	}
	isValidExtendedPublicKey := xpubKey.IsForNet(&cfg)
	return isValidExtendedPublicKey
}
