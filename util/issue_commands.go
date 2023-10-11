package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
)

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

	cmd := exec.Command("bash", "-c", cmdArgStr)

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

	cmd := exec.Command("bash", "-c", cmdArgStr)

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
