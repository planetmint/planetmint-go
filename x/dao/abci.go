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

	if isReissuanceHeight(currentBlockHeight) {
		reissuance, err := k.CreateNextReissuanceObject(ctx, currentBlockHeight)
		if err == nil {
			util.SendInitReissuance(ctx, hexProposerAddress, reissuance.GetCommand(), currentBlockHeight,
				reissuance.GetFirstIncludedPop(), reissuance.GetLastIncludedPop())
		} else {
			util.GetAppLogger().Error(ctx, "error while computing the RDDL reissuance ", err)
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
	conf := config.GetConfig()
	return height%int64(conf.PopEpochs) == 0
}

func isReissuanceHeight(height int64) bool {
	conf := config.GetConfig()
	// e.g. 483840 % 17280 = 0
	return height%int64(conf.ReissuanceEpochs) == 0
}

func isDistributionHeight(height int64) bool {
	conf := config.GetConfig()
	// e.g. 360 % 17280 = 360
	if height <= int64(conf.ReissuanceEpochs) {
		return false
	}
	// e.g. 484200 % 17280 = 360
	return height%int64(conf.ReissuanceEpochs) == int64(conf.DistributionOffset)
}

func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, k keeper.Keeper) {
	k.DistributeCollectedFees(ctx)
}
