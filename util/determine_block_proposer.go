package util

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	cometcfg "github.com/cometbft/cometbft/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
)

type Key struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type KeyFile struct {
	Address string `json:"address"`
	PubKey  Key    `json:"pub-key"`
	PrivKey Key    `json:"priv-key"`
}

type KeyringEntry struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Pubkey  string `json:"pubkey"`
}

func getAddressFromKeyringEntry(jsonString string) (bech32Address string) {
	// Slice to hold the parsed data
	var keyringEntrys []KeyringEntry

	// Unmarshal the JSON string into the slice
	err := json.Unmarshal([]byte(jsonString), &keyringEntrys)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Iterate over the parsed data and print the addresses
	for _, v := range keyringEntrys {
		fmt.Printf("Address: %s\n", v.Address)
		bech32Address = v.Address
	}
	return
}

func GetValidatorAddress(rootDir string, keyringDir string) (bech32Address string) {
	// Sanity check
	listCmd := "planetmint-god keys list --home " + rootDir + " --keyring-backend test --keyring-dir " + keyringDir + " --output json"
	output, err := runCommandWithOutput(listCmd)
	if err != nil {
		fmt.Println("Error listing keys:", err)
		return
	}
	bech32Address = getAddressFromKeyringEntry(output)
	return
}

func CreateTestingKeyring(rootDir string, keyringDir string, mnemonic string) (bech32Address string) {
	// Define the variables (replace with actual values)
	filePath := rootDir + "/keyring-test/validator.info"
	inventoryHostname := "validator"

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File does not exist, recover the key
		recoverCmd := fmt.Sprintf("echo '%s' | planetmint-god keys --home %s add %s --recover=true --keyring-dir %s --keyring-backend test", mnemonic, rootDir, inventoryHostname, keyringDir)
		output, err := runCommandWithOutput(recoverCmd)
		if err != nil {
			fmt.Println("Error recovering key:", err)
			fmt.Println(" output : " + output)
			return
		}
		fmt.Println(" output : " + output)
	}

	// Sanity check
	listCmd := "planetmint-god keys list --home " + rootDir + " --keyring-backend test --keyring-dir " + keyringDir + " --output json"
	output, err := runCommandWithOutput(listCmd)
	if err != nil {
		fmt.Println("Error listing keys:", err)
		return
	}
	bech32Address = getAddressFromKeyringEntry(output)
	return
}

func runCommand(cmdString string) error {
	cmd := exec.Command("bash", "-c", cmdString)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runCommandWithOutput(cmdString string) (string, error) {
	cmd := exec.Command("bash", "-c", cmdString)
	outputBytes, err := cmd.Output()
	return string(outputBytes), err
}

func GetValidatorCometBFTIdentity(ctx sdk.Context) (string, bool) {
	conf := config.GetConfig()

	cfg := cometcfg.DefaultConfig()
	jsonFilePath := filepath.Join(conf.ConfigRootDir, cfg.PrivValidatorKey)

	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		GetAppLogger().Error(ctx, "error while opening config", err.Error())
		return "", false
	}
	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		GetAppLogger().Error(ctx, "error while reading file", err.Error())
		return "", false
	}

	var keyFile KeyFile
	err = json.Unmarshal(jsonBytes, &keyFile)
	if err != nil {
		GetAppLogger().Error(ctx, "error while unmarshaling key file", err.Error())
		return "", false
	}
	return strings.ToLower(keyFile.Address), true
}

func IsValidatorBlockProposer(ctx sdk.Context, proposerAddress []byte) bool {
	validatorIdentity, validResult := GetValidatorCometBFTIdentity(ctx)
	if !validResult {
		return false
	}
	hexProposerAddress := hex.EncodeToString(proposerAddress)
	return hexProposerAddress == validatorIdentity
}
