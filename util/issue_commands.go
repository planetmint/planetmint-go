package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func execCommand(cmdArgsStr string) error {

	cmd := exec.Command("bash", "-c", cmdArgsStr)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Start the command in a non-blocking way
	err := cmd.Start()
	errstr := stderr.String()
	if err != nil || len(errstr) > 0 {
		if err == nil {
			err = errors.New(errstr)
		}
	}
	return err
}

func InitRDDLReissuanceProcess(ctx sdk.Context, proposerAddress string, tx_unsigned string, blk_height int64) error {
	//get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	// Construct the command
	sending_validator_address := config.GetConfig().ValidatorAddress
	planetmintKeyring := config.GetConfig().PlanetmintKeyring

	cmdArgStr := fmt.Sprintf("planetmint-god tx dao reissue-rddl-proposal %s '%s' %s --from %s -y",
		proposerAddress, tx_unsigned, strconv.FormatInt(blk_height, 10),
		sending_validator_address)
	if planetmintKeyring != "" {
		cmdArgStr = fmt.Sprintf("%s --keyring-backend %s", cmdArgStr, planetmintKeyring)
	}
	ctx.Logger().Debug("REISSUE: create Proposal")
	return execCommand(cmdArgStr)
}

func SendRDDLReissuanceResult(ctx sdk.Context, proposerAddress string, txID string, blk_height uint64) error {
	// Construct the command
	sending_validator_address := config.GetConfig().ValidatorAddress
	planetmintKeyring := config.GetConfig().PlanetmintKeyring

	cmdArgStr := fmt.Sprintf("planetmint-god tx dao reissue-rddl-result %s '%s' %s --from %s -y",
		proposerAddress, txID, strconv.FormatUint(blk_height, 10),
		sending_validator_address)
	if planetmintKeyring != "" {
		cmdArgStr = fmt.Sprintf("%s --keyring-backend %s", cmdArgStr, planetmintKeyring)
	}
	ctx.Logger().Debug("REISSUE: create Result")
	return execCommand(cmdArgStr)
}

func SendRDDLDistributionRequest(ctx sdk.Context, distribution types.DistributionOrder) error {
	// Construct the command
	sending_validator_address := config.GetConfig().ValidatorAddress
	planetmintKeyring := config.GetConfig().PlanetmintKeyring

	cmdArgStr := fmt.Sprintf("planetmint-god tx dao distribution-request '%s' --from %s -y",
		distribution.String(), sending_validator_address)
	if planetmintKeyring != "" {
		cmdArgStr = fmt.Sprintf("%s --keyring-backend %s", cmdArgStr, planetmintKeyring)
	}
	ctx.Logger().Debug("REISSUE: create Result")
	return execCommand(cmdArgStr)
}

func SendRDDLDistributionResult(ctx sdk.Context, last_pop string, dao_txid string, inv_txid string, pop_txid string) error {
	// Construct the command
	sending_validator_address := config.GetConfig().ValidatorAddress
	planetmintKeyring := config.GetConfig().PlanetmintKeyring

	cmdArgStr := fmt.Sprintf("planetmint-god tx dao distribution-result %s %s %s %s --from %s -y",
		last_pop, dao_txid, inv_txid, pop_txid,
		sending_validator_address)
	if planetmintKeyring != "" {
		cmdArgStr = fmt.Sprintf("%s --keyring-backend %s", cmdArgStr, planetmintKeyring)
	}
	ctx.Logger().Debug("REISSUE: create Result")
	return execCommand(cmdArgStr)
}
