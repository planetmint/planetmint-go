package dao

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/crgimenes/go-osc"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/keeper"

	abci "github.com/cometbft/cometbft/abci/types"
	cometcfg "github.com/cometbft/cometbft/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	logger := ctx.Logger()
	proposerAddress := req.Header.GetProposerAddress()

	// Check if node is block proposer
	if isPoPHeight(req.Header.GetHeight()) && isValidatorBlockProposer(ctx, proposerAddress) {
		// TODO: implement PoP trigger
		fmt.Println("TODO: implement PoP trigger")
		err := issueRDDL()
		if err != nil {
			logger.Error("error while issuing RDDL", err)
		}
	}
}

// TODO: define final message
func issueRDDL() error {
	cfg := config.GetConfig()
	client := osc.NewClient(cfg.IssuanceEndpoint, cfg.IssuancePort)

	msg := osc.NewMessage("/rddl/token")
	err := client.Send(msg)

	return err
}

func isPoPHeight(height int64) bool {
	cfg := config.GetConfig()
	return height%int64(cfg.PoPEpochs) == 0
}

func isValidatorBlockProposer(ctx sdk.Context, proposerAddress []byte) bool {
	logger := ctx.Logger()
	conf := config.GetConfig()

	cfg := cometcfg.DefaultConfig()
	jsonFilePath := filepath.Join(conf.ConfigRootDir, cfg.PrivValidatorKey)

	hexProposerAddress := hex.EncodeToString(proposerAddress)

	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		logger.Error("error while opening config", err)
	}
	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		logger.Error("error while reading file", err)
	}

	var keyFile KeyFile
	err = json.Unmarshal(jsonBytes, &keyFile)
	if err != nil {
		logger.Error("error while unmarshaling key file", err)
	}

	return hexProposerAddress == strings.ToLower(keyFile.Address)
}

func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {
	k.DistributeCollectedFees(ctx)
}
