package util

import (
	"os/exec"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func InitRDDLReissuanceProcess(ctx sdk.Context, proposerAddress string, txUnsigned string, blockHeight int64) error {
	//get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	logger := ctx.Logger()
	// Construct the command
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Proposal")
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-proposal",
		"--from", sendingValidatorAddress, "-y",
		proposerAddress, txUnsigned, strconv.FormatInt(blockHeight, 10))

	logger.Debug("REISSUE: create Proposal")
	return cmd.Start()
}

func SendRDDLReissuanceResult(ctx sdk.Context, proposerAddress string, txID string, blockHeight int64) error {
	logger := ctx.Logger()
	// Construct the command
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Result")
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-result",
		"--from", sendingValidatorAddress, "-y",
		proposerAddress, txID, strconv.FormatInt(blockHeight, 10))
	logger.Debug("REISSUE: create Result")
	return cmd.Start()
}

func SendRDDLDistributionRequest(ctx sdk.Context, distribution types.DistributionOrder) error {
	logger := ctx.Logger()
	// Construct the command
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	cmd := exec.Command("planetmint-god", "tx", "dao", "distribution-request",
		"--from", sendingValidatorAddress, "-y", "'"+distribution.String()+"'")
	logger.Debug("REISSUE: create Result")
	return cmd.Start()
}

func SendRDDLDistributionResult(ctx sdk.Context, lastPoP string, daoTxid string, invTxid string, popTxid string) error {
	logger := ctx.Logger()
	// Construct the command
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Result")
	cmd := exec.Command("planetmint-god", "tx", "dao", "distribution-result",
		"--from", sendingValidatorAddress, "-y",
		lastPoP, daoTxid, invTxid, popTxid)
	logger.Debug("REISSUE: create Result")
	return cmd.Start()
}
