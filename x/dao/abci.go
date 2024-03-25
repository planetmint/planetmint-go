package dao

import (
	"encoding/hex"

	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	proposerAddress := req.Header.GetProposerAddress()

	// Check if node is block proposer
	// take the following actions only once, that's why we filter for the Block Proposer
	if !util.IsValidatorBlockProposer(ctx, proposerAddress, k.RootDir) {
		return
	}
	currentBlockHeight := req.Header.GetHeight()

	params := k.GetParams(ctx)
	hexProposerAddress := hex.EncodeToString(proposerAddress)
	go func() {
		if isPopHeight(params, currentBlockHeight) {
			// select PoP participants
			challenger, challengee := k.SelectPopParticipants(ctx)

			// Init PoP - independent from challenger and challengee
			// The keeper will send the MQTT initializing message to challenger && challengee
			util.SendInitPoP(ctx, hexProposerAddress, challenger, challengee, currentBlockHeight)
		}
	}()

	if isReissuanceHeight(params, currentBlockHeight) {
		reissuance, err := k.CreateNextReissuanceObject(ctx, currentBlockHeight)
		if err == nil {
			util.SendInitReissuance(ctx, hexProposerAddress, reissuance.GetCommand(), currentBlockHeight,
				reissuance.GetFirstIncludedPop(), reissuance.GetLastIncludedPop())
		} else {
			util.GetAppLogger().Error(ctx, "error while computing the RDDL reissuance ", err)
		}
	}
	if isDistributionHeight(params, currentBlockHeight) {
		distribution, err := k.GetDistributionForReissuedTokens(ctx, currentBlockHeight)
		if err != nil {
			util.GetAppLogger().Error(ctx, "error while computing the RDDL distribution ", err)
		}
		distribution.Proposer = hexProposerAddress
		util.SendDistributionRequest(ctx, distribution)
	}
}

func isPopHeight(params daotypes.Params, height int64) bool {
	return height%params.PopEpochs == 0
}

func isReissuanceHeight(params daotypes.Params, height int64) bool {
	// e.g. 483840 % 17280 = 0
	return height%params.ReissuanceEpochs == 0
}

func isDistributionHeight(params daotypes.Params, height int64) bool {
	// e.g. 360 % 17280 = 360
	if height <= params.ReissuanceEpochs {
		return false
	}
	// e.g. 484200 % 17280 = 360
	return height%params.ReissuanceEpochs == params.DistributionOffset
}

func EndBlocker(_ sdk.Context, _ abci.RequestEndBlock, _ keeper.Keeper) {
	// EndBlocker is currently not implemented and used by planetmint
}
