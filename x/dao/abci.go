package dao

import (
	"encoding/hex"
	"fmt"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/keeper"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	logger := ctx.Logger()
	proposerAddress := req.Header.GetProposerAddress()

	// Check if node is block proposer
	// take the following actions only once, that's why we filter for the Block Proposer
	if util.IsValidatorBlockProposer(ctx, proposerAddress) {
		blockHeight := req.Header.GetHeight()
		if isPoPHeight(blockHeight) {
			// TODO: implement PoP trigger
			fmt.Println("TODO: implement PoP trigger")
			hexProposerAddress := hex.EncodeToString(proposerAddress)
			conf := config.GetConfig()
			tx_unsigned := GetReissuanceCommand(conf.ReissuanceAsset, blockHeight)
			err := util.InitRDDLReissuanceProcess(ctx, hexProposerAddress, tx_unsigned, blockHeight)
			if err != nil {
				logger.Error("error while initializing RDDL issuance", err)
			}
		}
		if isDistributionHeight(blockHeight) {
			// initialize the distribution message
		}
	}

}

func isPoPHeight(height int64) bool {
	cfg := config.GetConfig()
	return height%int64(cfg.PoPEpochs) == 0
}

func isDistributionHeight(height int64) bool {
	cfg := config.GetConfig()
	return height%int64(cfg.DistributionEpochs) == 0
}

func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {
	k.DistributeCollectedFees(ctx)
}
