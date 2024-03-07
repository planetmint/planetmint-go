package util

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/machine/types"
	elements "github.com/rddl-network/elements-rpc"
)

var (
	// this mutex has to protect all signing and crafting transactions so that UTXOs are not spend twice by accident
	elementsSyncAccess sync.Mutex
)

func ReissueAsset(reissueTx string) (txID string, err error) {
	conf := config.GetConfig()
	url := conf.GetRPCURL()
	cmdArgs := strings.Split(reissueTx, " ")
	elementsSyncAccess.Lock()
	defer elementsSyncAccess.Unlock()
	result, err := elements.ReissueAsset(url, []string{cmdArgs[1], cmdArgs[2]})
	if err != nil {
		return
	}
	txID = result.TxID
	return
}

func DistributeAsset(address string, amount string, reissuanceAsset string) (txID string, err error) {
	conf := config.GetConfig()
	url := conf.GetRPCURL()

	elementsSyncAccess.Lock()
	defer elementsSyncAccess.Unlock()
	txID, err = elements.SendToAddress(url, []string{
		address,
		`"` + amount + `"`,
		`""`,
		`""`,
		"false",
		"true",
		"null",
		`"unset"`,
		"false",
		`"` + reissuanceAsset + `"`,
	})
	return
}

func IssueNFTAsset(name string, machineAddress string, domain string) (assetID string, contract string, hexTx string, err error) {
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

	elementsSyncAccess.Lock()
	defer elementsSyncAccess.Unlock()
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

	return assetID, contract, hex, err
}
