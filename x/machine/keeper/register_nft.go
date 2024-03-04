package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	config "github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/machine/types"
	elements "github.com/rddl-network/elements-rpc"
)

var (
	assetClientService     IAssetServiceClient
	initAssetServiceClient sync.Once
)

type IAssetServiceClient interface {
	IssueNFTAsset(goCtx context.Context, name string, machineAddress string, domain string) (assetID string, contract string, err error)
	IssueMachineNFT(goCtx context.Context, machine *types.Machine, scheme string, domain string, path string) error
	RegisterAsset(goCtx context.Context, assetID string, contract string, assetRegistryEndpoint string) error
}

type AssetServiceClient struct{}

func GetAssetServiceClient() IAssetServiceClient {
	initAssetServiceClient.Do(func() {
		assetClientService = &AssetServiceClient{}
	})
	return assetClientService
}

func SetAssetServiceClient(asc IAssetServiceClient) {
	assetClientService = asc
}

func (asc *AssetServiceClient) IssueNFTAsset(goCtx context.Context, name string, machineAddress string, domain string) (assetID string, contract string, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	conf := config.GetConfig()

	url := conf.GetRPCURL()
	address, err := elements.GetNewAddress(url, []string{``})
	if err != nil {
		return
	}

	addressInfo, err := elements.GetAddressInfo(url, []string{address})
	if err != nil {
		return
	}

	hex, err := elements.CreateRawTransaction(url, []string{`[]`, `[{"data":"00"}]`})
	if err != nil {
		return
	}

	fundRawTransactionResult, err := elements.FundRawTransaction(url, []string{hex, `{"feeRate":0.00001000}`})
	if err != nil {
		return
	}

	c := types.Contract{
		Entity: types.Entity{
			Domain: domain,
		},
		IssuerPubkey: addressInfo.Pubkey,
		MachineAddr:  machineAddress,
		Name:         name,
		Precision:    0,
		Version:      0,
	}
	contractBytes, err := json.Marshal(c)
	if err != nil {
		return
	}
	// e.g. {"entity":{"domain":"testnet-assets.rddl.io"}, "issuer_pubkey":"02...}
	contract = string(contractBytes)

	h := sha256.New()
	_, err = h.Write(contractBytes)
	if err != nil {
		return
	}
	// e.g. 7ca8bb403ee5dccddef7b89b163048cf39439553f0402351217a4a03d2224df8
	hash := h.Sum(nil)

	// Reverse hash, e.g. f84d22d2034a7a21512340f053954339cf4830169bb8f7decddce53e40bba87c
	for i, j := 0, len(hash)-1; i < j; i, j = i+1, j-1 {
		hash[i], hash[j] = hash[j], hash[i]
	}

	rawIssueAssetResults, err := elements.RawIssueAsset(url, []string{fundRawTransactionResult.Hex,
		`[{"asset_amount":0.00000001, "asset_address":"` + address + `", "blind":false, "contract_hash":"` + fmt.Sprintf("%+x", hash) + `"}]`,
	})
	if err != nil {
		return
	}

	rawIssueAssetResult := rawIssueAssetResults[len(rawIssueAssetResults)-1]
	hex, err = elements.BlindRawTransaction(url, []string{rawIssueAssetResult.Hex, `true`, `[]`, `false`})
	if err != nil {
		return
	}
	assetID = rawIssueAssetResult.Asset

	signRawTransactionWithWalletResult, err := elements.SignRawTransactionWithWallet(url, []string{hex})
	if err != nil {
		return
	}

	testMempoolAcceptResults, err := elements.TestMempoolAccept(url, []string{`["` + signRawTransactionWithWalletResult.Hex + `"]`})
	if err != nil {
		return
	}

	testMempoolAcceptResult := testMempoolAcceptResults[len(testMempoolAcceptResults)-1]
	if !testMempoolAcceptResult.Allowed {
		log.Fatalln("not accepted by mempool")
	}

	hex, err = elements.SendRawTransaction(url, []string{signRawTransactionWithWalletResult.Hex})
	if err != nil {
		return
	}

	util.GetAppLogger().Info(ctx, "Liquid Token Issuance assetID: "+assetID+" contract: "+contract+" tx: "+hex)
	return assetID, contract, err
}

func (asc *AssetServiceClient) IssueMachineNFT(goCtx context.Context, machine *types.Machine, scheme string, domain string, path string) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// asset registration is in order to have the contact published
	var notarizedAsset types.LiquidAsset
	notarizedAsset.Registered = true
	assetID, contract, err := asc.IssueNFTAsset(goCtx, machine.Name, machine.Address, domain)
	if err != nil {
		util.GetAppLogger().Error(ctx, err.Error())
		return err
	}
	assetRegistryEndpoint := fmt.Sprintf("%s://%s/%s", scheme, domain, path)
	fmt.Println(" Register Asset: " + assetRegistryEndpoint)
	fmt.Println(" CONTRACT: " + contract)
	err = asc.RegisterAsset(goCtx, assetID, contract, assetRegistryEndpoint)
	if err != nil {
		util.GetAppLogger().Error(ctx, err.Error())
		notarizedAsset.Registered = false
	}
	// issue message with:
	notarizedAsset.AssetID = assetID
	notarizedAsset.MachineID = machine.GetMachineId()
	notarizedAsset.MachineAddress = machine.Address

	util.SendLiquidAssetRegistration(goCtx, notarizedAsset)
	return err
}

func (asc *AssetServiceClient) RegisterAsset(goCtx context.Context, assetID string, contract string, assetRegistryEndpoint string) error {
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
	if strings.Contains(resultObj, assetID) {
		return nil
	}
	return errorsmod.Wrap(types.ErrAssetRegistryRepsonse, "does not confirm asset registration")
}
