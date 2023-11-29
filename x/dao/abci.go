package dao

import (
	"encoding/hex"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/keeper"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, _ keeper.Keeper) {
	proposerAddress := req.Header.GetProposerAddress()

	// Check if node is block proposer
	// take the following actions only once, that's why we filter for the Block Proposer
	if !util.IsValidatorBlockProposer(ctx, proposerAddress) {
		return
	}
	blockHeight := req.Header.GetHeight()
	if isPoPHeight(req.Header.GetHeight()) {
		hexProposerAddress := hex.EncodeToString(proposerAddress)
		// select PoP participants
		challenger := ""
		challengee := ""

		// Issue PoP
		util.SendInitPoP(ctx, hexProposerAddress, challenger, challengee, blockHeight)
		// TODO send MQTT message to challenger && challengee
	}
	// TODO will be reintegrated with by merging branch 184-implement-staged-claim
	// if isDistributionHeight(blockHeight) {
	// // reissue 1st

	// conf := config.GetConfig()
	// txUnsigned := keeper.GetReissuanceCommand(conf.ReissuanceAsset, blockHeight)
	// util.SendInitReissuance(ctx, hexProposerAddress, txUnsigned, blockHeight)

	// // distribute thereafter
	//// initialize the distribution message
	// distribution, err := k.GetDistributionForReissuedTokens(ctx, blockHeight)
	// if err != nil {
	// util.GetAppLogger().Error(ctx, "error while computing the RDDL distribution ", err)
	// }
	// util.SendDistributionRequest(ctx, distribution)
	// }
}

func isPoPHeight(height int64) bool {
	cfg := config.GetConfig()
	return height%int64(cfg.PoPEpochs) == 0
}

// TODO will be reintegrated with by merging branch 184-implement-staged-claim
// func isDistributionHeight(height int64) bool {
// 	cfg := config.GetConfig()
// 	return height%int64(cfg.DistributionEpochs) == 0
// }

func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, k keeper.Keeper) {
	k.DistributeCollectedFees(ctx)
}
