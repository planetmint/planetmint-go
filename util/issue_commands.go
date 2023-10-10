package util

import (
	"bytes"
	"errors"
	"os/exec"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
)

func InitRDDLReissuanceProcess(ctx sdk.Context, proposerAddress string, tx_unsigned string, blk_height int64) error {
	//get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	// Construct the command
	sending_validator_address := config.GetConfig().ValidatorAddress
	keyring := config.GetConfig().PlanetmintKeyring
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-proposal",
		"--from", sending_validator_address, "-y",
		proposerAddress, tx_unsigned, strconv.FormatInt(blk_height, 10))
	if keyring != "" {
		cmd = exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-proposal",
			"--from", sending_validator_address, "-y", "--keyring-backend ", keyring,
			proposerAddress, tx_unsigned, strconv.FormatInt(blk_height, 10))
	}
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
	keyring := config.GetConfig().PlanetmintKeyring
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-result",
		"--from", sending_validator_address, "-y",
		proposerAddress, txID, strconv.FormatUint(blk_height, 10))
	if keyring != "" {
		cmd = exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-result",
			"--from", sending_validator_address, "-y", "--keyring-backend ", keyring,
			proposerAddress, txID, strconv.FormatUint(blk_height, 10))
	}
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
