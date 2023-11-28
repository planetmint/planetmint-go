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
	PubKey  Key    `json:"pub-key"`
	PrivKey Key    `json:"priv-key"`
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
