package util

import (
	"fmt"
	"os/exec"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitRDDLReissuanceProcess(ctx sdk.Context, proposerAddress string, blk_height int64) error {
	tx_unsigned, err := GetUnsignedReissuanceTransaction()
	//blk_height := 0 //get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store

	// Construct the command
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-proposal", proposerAddress, tx_unsigned, strconv.FormatInt(blk_height, 10))

	// Start the command in a non-blocking way
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting command: %s\n", err)
	} else {
		fmt.Println("Command started in background")
	}
	return err
}

func SendRDDLReissuanceResult(ctx sdk.Context, proposerAddress string, txID string, blk_height uint64) error {
	// Construct the command
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-result", proposerAddress, txID, strconv.FormatUint(blk_height, 10))

	// Start the command in a non-blocking way
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error starting command: %s\n", err)
	} else {
		fmt.Println("Command started in background")
	}
	return err
}
