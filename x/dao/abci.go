package dao

import (
	"encoding/hex"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/keeper"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	proposerAddress := req.Header.GetProposerAddress()

	// Check if node is block proposer
	// take the following actions only once, that's why we filter for the Block Proposer
	if !util.IsValidatorBlockProposer(ctx, proposerAddress) {
		return
	}
	currentBlockHeight := req.Header.GetHeight()

	hexProposerAddress := hex.EncodeToString(proposerAddress)
	if isPopHeight(req.Header.GetHeight()) {
		// select PoP participants
		challenger, challengee := k.SelectPopParticipants(ctx)

		if challenger != "" && challengee != "" {
			// Issue PoP
			util.SendInitPoP(ctx, hexProposerAddress, challenger, challengee, currentBlockHeight)
			// TODO send MQTT message to challenger && challengee
		}
	}

	if isReIssuanceHeight(currentBlockHeight) {
		reIssuance, err := k.CreateNextReIssuanceObject(ctx, currentBlockHeight)
		if err == nil {
			util.SendInitReissuance(ctx, hexProposerAddress, reIssuance.GetRawTx(), currentBlockHeight,
				reIssuance.GetFirstIncludedPop(), reIssuance.GetLastIncludedPop())
		} else {
			util.GetAppLogger().Error(ctx, "error while computing the RDDL re-issuance ", err)
		}
	}

	if isDistributionHeight(currentBlockHeight) {
		distribution, err := k.GetDistributionForReissuedTokens(ctx, currentBlockHeight)
		if err != nil {
			util.GetAppLogger().Error(ctx, "error while computing the RDDL distribution ", err)
		}
		distribution.Proposer = hexProposerAddress
		util.SendDistributionRequest(ctx, distribution)
	}
}

func isPopHeight(height int64) bool {
	cfg := config.GetConfig()
	return height%int64(cfg.PopEpochs) == 0
}

func isReIssuanceHeight(height int64) bool {
	cfg := config.GetConfig()
	return height%int64(cfg.ReIssuanceEpochs) == 0
}

func isDistributionHeight(height int64) bool {
	cfg := config.GetConfig()
	return height%int64(cfg.DistributionEpochs) == 0
}

func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, k keeper.Keeper) {
	k.DistributeCollectedFees(ctx)
}
