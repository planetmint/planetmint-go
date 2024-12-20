package util

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/planetmint/planetmint-go/errormsg"

	cometcfg "github.com/cometbft/cometbft/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func GetValidatorCometBFTIdentity(ctx sdk.Context, rootDir string) (validatorIdentity string, err error) {
	cfg := cometcfg.DefaultConfig()
	jsonFilePath := filepath.Join(rootDir, cfg.PrivValidatorKey)

	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		GetAppLogger().Error(ctx, err, "error while opening config: %v", jsonFilePath)
		return
	}
	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		GetAppLogger().Error(ctx, err, "error while reading file: %v", jsonFile)
		return
	}

	var keyFile KeyFile
	err = json.Unmarshal(jsonBytes, &keyFile)
	if err != nil {
		GetAppLogger().Error(ctx, err, "error while unmarshaling key file")
		return
	}
	validatorIdentity = strings.ToLower(keyFile.Address)
	return
}

func IsValidatorBlockProposer(ctx sdk.Context, rootDir string) (result bool) {
	validatorIdentity, err := GetValidatorCometBFTIdentity(ctx, rootDir)
	if err != nil {
		GetAppLogger().Error(ctx, err, errormsg.CouldNotGetValidatorIdentity)
		return
	}
	hexProposerAddress := hex.EncodeToString(ctx.BlockHeader().ProposerAddress)
	result = hexProposerAddress == validatorIdentity
	return
}

func IsValidAddress(address string) (valid bool, err error) {
	// Attempt to decode the address
	_, err = sdk.AccAddressFromBech32(address)
	if err != nil {
		return
	}
	if !strings.Contains(address, "plmnt") {
		valid = false
		return
	}
	valid = true
	return
}
