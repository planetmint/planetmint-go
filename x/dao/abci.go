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

	if isPoPHeight(req.Header.GetHeight()) && util.IsValidatorBlockProposer(ctx, proposerAddress) {
		// TODO: implement PoP trigger
		fmt.Println("TODO: implement PoP trigger")
		hexProposerAddress := hex.EncodeToString(proposerAddress)
		err := util.InitRDDLReissuanceProcess(ctx, hexProposerAddress, req.Header.GetHeight())
		if err != nil {
			logger.Error("error while issuing RDDL", err)
		}
	}
}

func isPoPHeight(height int64) bool {
	cfg := config.GetConfig()
	return height%int64(cfg.PoPEpochs) == 0
}

func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {
	k.DistributeCollectedFees(ctx)
}
