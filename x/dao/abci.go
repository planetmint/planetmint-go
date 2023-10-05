package dao

import (
	"encoding/hex"
	"fmt"
	"os/exec"
	"strconv"

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
		err := initRDDLReissuanceProcess(ctx, proposerAddress, req.Header.GetHeight())
		if err != nil {
			logger.Error("error while issuing RDDL", err)
		}
	}
}

func initRDDLReissuanceProcess(ctx sdk.Context, proposerAddress []byte, blk_height int64) error {
	tx_unsigned, err := util.GetUnsignedReissuanceTransaction()
	//blk_height := 0 //get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	hexProposerAddress := hex.EncodeToString(proposerAddress)

	// Construct the command
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-proposal", hexProposerAddress, tx_unsigned, strconv.Itoa(blk_height))

	// Start the command in a non-blocking way
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting command: %s\n", err)
	} else {
		fmt.Println("Command started in background")
	}
	return err
}

func isPoPHeight(height int64) bool {
	cfg := config.GetConfig()
	return height%int64(cfg.PoPEpochs) == 0
}

func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {
	k.DistributeCollectedFees(ctx)
}
