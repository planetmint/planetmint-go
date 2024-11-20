package dao

import (
	"encoding/hex"

	"github.com/planetmint/planetmint-go/monitor"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/keeper"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	proposerAddress := ctx.BlockHeader().ProposerAddress

	// Check if node is block proposer
	// take the following actions only once, that's why we filter for the Block Proposer
	if !util.IsValidatorBlockProposer(ctx, k.RootDir) {
		return
	}
	currentBlockHeight := req.Header.GetHeight()

	hexProposerAddress := hex.EncodeToString(proposerAddress)
	if isPopHeight(ctx, k, currentBlockHeight) {
		// select PoP participants
		challenger, challengee, err := monitor.SelectPoPParticipantsOutOfActiveActors()
		if err != nil {
			util.GetAppLogger().Error(ctx, err, "error during PoP Participant selection")
		}
		if err != nil || challenger == "" || challengee == "" {
			challenger = ""
			challengee = ""
		}

		// Init PoP - independent from challenger and challengee
		// The keeper will send the MQTT initializing message to challenger && challengee
		util.SendInitPoP(ctx, challenger, challengee, currentBlockHeight)
	}

	if isReissuanceHeight(ctx, k, currentBlockHeight) {
		reissuance, err := k.CreateNextReissuanceObject(ctx, currentBlockHeight)
		if err == nil {
			util.SendInitReissuance(ctx, hexProposerAddress, reissuance.GetCommand(), currentBlockHeight,
				reissuance.GetFirstIncludedPop(), reissuance.GetLastIncludedPop())
		} else {
			util.GetAppLogger().Error(ctx, err, "error while computing the RDDL reissuance")
		}
	}

	if isDistributionHeight(ctx, k, currentBlockHeight) {
		distribution, err := k.GetDistributionForReissuedTokens(ctx, currentBlockHeight)
		if err != nil {
			util.GetAppLogger().Error(ctx, err, "error while computing the RDDL distribution")
		}
		distribution.Proposer = hexProposerAddress
		util.SendDistributionRequest(ctx, distribution)
	}
}

func isPopHeight(ctx sdk.Context, k keeper.Keeper, height int64) bool {
	return height%k.GetParams(ctx).PopEpochs == 0
}

func isReissuanceHeight(ctx sdk.Context, k keeper.Keeper, height int64) bool {
	// e.g. 483840 % 17280 = 0
	return height%k.GetParams(ctx).ReissuanceEpochs == 0
}

func isDistributionHeight(ctx sdk.Context, k keeper.Keeper, height int64) bool {
	// e.g. 360 % 17280 = 360
	if height <= k.GetParams(ctx).ReissuanceEpochs {
		return false
	}
	// e.g. 484200 % 17280 = 360
	return height%k.GetParams(ctx).ReissuanceEpochs == k.GetParams(ctx).DistributionOffset
}

func EndBlocker(_ sdk.Context, _ abci.RequestEndBlock, _ keeper.Keeper) {
	// EndBlocker is currently not implemented and used by planetmint
}
