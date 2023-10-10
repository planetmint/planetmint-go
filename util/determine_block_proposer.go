package util

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
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
	PubKey  Key    `json:"pub_key"`
	PrivKey Key    `json:"priv_key"`
}

func GetValidatorCometBFTIdentity(ctx sdk.Context) (string, bool) {
	logger := ctx.Logger()
	conf := config.GetConfig()

	cfg := cometcfg.DefaultConfig()
	jsonFilePath := filepath.Join(conf.ConfigRootDir, cfg.PrivValidatorKey)

	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		logger.Error("error while opening config", err)
		return "", false
	}
	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		logger.Error("error while reading file", err)
		return "", false
	}

	var keyFile KeyFile
	err = json.Unmarshal(jsonBytes, &keyFile)
	if err != nil {
		logger.Error("error while unmarshaling key file", err)
		return "", false
	}
	return strings.ToLower(keyFile.Address), true
}

func IsValidatorBlockProposer(ctx sdk.Context, proposerAddress []byte) bool {
	validator_identity, valid_result := GetValidatorCometBFTIdentity(ctx)
	if !valid_result {
		return false
	}
	hexProposerAddress := hex.EncodeToString(proposerAddress)
	return hexProposerAddress == validator_identity
}
