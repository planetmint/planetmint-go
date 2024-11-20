package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/clients"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	RegisterAssetServiceHTTPClient HTTPClient
)

func init() {
	RegisterAssetServiceHTTPClient = &http.Client{}
}

func IssueMachineNFT(goCtx context.Context, machine *types.Machine, scheme string, domain string, path string) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// asset registration is in order to have the contact published
	var notarizedAsset types.LiquidAsset
	notarizedAsset.Registered = true
	assetID, contract, hex, err := clients.IssueNFTAsset(goCtx, machine.Name, machine.Address, domain)
	if err != nil {
		GetAppLogger().Error(ctx, err, "")
		return err
	}
	assetRegistryEndpoint := fmt.Sprintf("%s://%s/%s", scheme, domain, path)

	GetAppLogger().Info(ctx, "Liquid Token Issuance assetID: "+assetID+" contract: "+contract+" tx: "+hex)
	err = RegisterAsset(goCtx, assetID, contract, assetRegistryEndpoint)
	if err != nil {
		GetAppLogger().Error(ctx, err, "")
		notarizedAsset.Registered = false
	}
	// issue message with:
	notarizedAsset.AssetID = assetID
	notarizedAsset.MachineID = machine.GetMachineId()
	notarizedAsset.MachineAddress = machine.Address

	SendLiquidAssetRegistration(goCtx, notarizedAsset)
	return err
}

func RegisterAsset(goCtx context.Context, assetID string, contract string, assetRegistryEndpoint string) error {
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

	req, err := http.NewRequestWithContext(goCtx, http.MethodPost, assetRegistryEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return errorsmod.Wrap(types.ErrAssetRegistryReqFailure, "Request creation: "+err.Error())
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	// Send request
	resp, err := RegisterAssetServiceHTTPClient.Do(req)
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
	if strings.Contains(resultObj, assetID) {
		return nil
	}
	return errorsmod.Wrap(types.ErrAssetRegistryRepsonse, "does not confirm asset registration")
}
