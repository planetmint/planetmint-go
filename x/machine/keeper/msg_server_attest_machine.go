package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

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

	if util.IsValidatorBlockProposer(ctx, ctx.BlockHeader().ProposerAddress) {
		util.GetAppLogger().Info(ctx, "Issuing Machine NFT")
		err := k.issueMachineNFT(goCtx, msg.Machine)
		if err != nil {
			return nil, types.ErrNFTIssuanceFailed
		}
	} else {
		util.GetAppLogger().Info(ctx, "skipping Machine NFT issuance")
	}

	k.StoreMachine(ctx, *msg.Machine)
	k.StoreMachineIndex(ctx, *msg.Machine)
	err = k.StoreTrustAnchor(ctx, ta, true)
	if err != nil {
		return nil, err
	}
	return &types.MsgAttestMachineResponse{}, err
}

func validateExtendedPublicKey(issuer string, cfg chaincfg.Params) bool {
	xpubKey, err := hdkeychain.NewKeyFromString(issuer)
	if err != nil {
		return false
	}
	isValidExtendedPublicKey := xpubKey.IsForNet(&cfg)
	return isValidExtendedPublicKey
}

func (k msgServer) issueNFTAsset(ctx sdk.Context, name string, machineAddress string) (assetID string, contract string, err error) {
	conf := config.GetConfig()
	cmdName := "poetry"
	cmdArgs := []string{"run", "python", "issuer_service/issue2liquid.py", name, machineAddress}

	// Create a new command
	cmd := exec.Command(cmdName, cmdArgs...)

	// If you want to set the working directory
	cmd.Dir = conf.IssuanceServiceDir

	// Capture the output
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	err = cmd.Run()
	if err != nil {
		util.GetAppLogger().Error(ctx, "Issue2Liquid.py failed with %s\n", err)
		err = errorsmod.Wrap(types.ErrMachineNFTIssuance, stderr.String())
	} else {
		util.GetAppLogger().Info(ctx, "Liquid Token Issuance: "+stdout.String())
		lines := strings.Split(stdout.String(), "\n")
		if len(lines) == 3 {
			assetID = lines[0]
			contract = lines[1]
		} else {
			err = errorsmod.Wrap(types.ErrMachineNFTIssuanceNoOutput, stderr.String())
		}
	}
	return assetID, contract, err
}

func (k msgServer) registerAsset(assetID string, contract string) error {
	conf := config.GetConfig()

	var contractMap map[string]interface{}
	err := json.Unmarshal([]byte(contract), &contractMap)
	if err != nil {
		return errorsmod.Wrap(types.ErrAssetRegistryReqFailure, "Unmarshal "+err.Error())
	}
	// Create your request payload
	data := map[string]interface{}{
		"asset_id": assetID,
		"contract": contractMap,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return errorsmod.Wrap(types.ErrAssetRegistryReqFailure, "Marshall "+err.Error())
	}

	req, err := http.NewRequest("POST", conf.AssetRegistryEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return errorsmod.Wrap(types.ErrAssetRegistryReqFailure, "Request creation: "+err.Error())
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errorsmod.Wrap(types.ErrAssetRegistryReqSending, err.Error())
	}
	defer resp.Body.Close()

	// Read response
	if resp.StatusCode > 299 {
		return errorsmod.Wrap(types.ErrAssetRegistryRepsonse, "Error reading response body:"+strconv.Itoa(resp.StatusCode))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errorsmod.Wrap(types.ErrAssetRegistryRepsonse, "Error reading response body:"+err.Error())
	}
	resultObj := string(body)
	if !strings.Contains(resultObj, assetID) {
		return errorsmod.Wrap(types.ErrAssetRegistryRepsonse, "does not confirm asset registration")
	}
	return nil
}

func (k msgServer) issueMachineNFT(goCtx context.Context, machine *types.Machine) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// asset registration is in order to have the contact published
	var notarizedAsset types.LiquidAsset
	notarizedAsset.Registered = true
	assetID, contract, err := k.issueNFTAsset(ctx, machine.Name, machine.Address)
	if err != nil {
		util.GetAppLogger().Error(ctx, err.Error())
		return err
	}
	err = k.registerAsset(assetID, contract)
	if err != nil {
		util.GetAppLogger().Error(ctx, err.Error())
		notarizedAsset.Registered = false
	}
	// issue message with:
	notarizedAsset.AssetID = assetID
	notarizedAsset.MachineID = machine.GetMachineId()
	notarizedAsset.MachineAddress = machine.Address

	err = util.SendLiquidAssetRegistration(goCtx, notarizedAsset)
	return err
}
