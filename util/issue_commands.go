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
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-proposal",
		"--from", sending_validator_address, "-y",
		proposerAddress, tx_unsigned, strconv.FormatInt(blk_height, 10))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Start the command in a non-blocking way
	err := cmd.Start()
	outstr := stdout.String()
	errstr := stderr.String()
	if err != nil || len(errstr) > 0 {
		fmt.Printf("Error starting command: s\n", errstr)
		if err == nil {
			err = errors.New(errstr)
		}
	} else {
		fmt.Println("Command started in background %s\n", outstr)
	}
	return err
}

func SendRDDLReissuanceResult(ctx sdk.Context, proposerAddress string, txID string, blk_height uint64) error {
	// Construct the command
	sending_validator_address := config.GetConfig().ValidatorAddress
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-result",
		"--from", sending_validator_address, "-y",
		proposerAddress, txID, strconv.FormatUint(blk_height, 10))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	// Start the command in a non-blocking way
	err := cmd.Start()
	outstr := stdout.String()
	errstr := stderr.String()

	if err != nil || len(errstr) > 0 {
		fmt.Printf("Error starting command: s\n", errstr)
		if err == nil {
			err = errors.New(errstr)
		}
	} else {
		fmt.Println("Command started in background %s\n", outstr)
	}
	return err
}
