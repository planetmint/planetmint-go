package dao

import (
	"fmt"

	"github.com/crgimenes/go-osc"
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

func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {
	k.DistributeCollectedFees(ctx)
}
