package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
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

func SendRDDLReissuanceResult(ctx sdk.Context, proposerAddress string, txID string, blockHeight int64) error {
	logger := ctx.Logger()
	// Construct the command
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Result")
	cmd := exec.Command("planetmint-god", "tx", "dao", "reissue-rddl-result",
		"--from", sendingValidatorAddress, "-y",
		proposerAddress, txID, strconv.FormatInt(blockHeight, 10))

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

func SendLiquidAssetRegistration(ctx sdk.Context, notarizedAsset machinetypes.LiquidAsset) error {
	logger := ctx.Logger()
	// Construct the command
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Result")
	obj := fmt.Sprintf("{ \"MachineID\": \"%s\", \"MachineAddress\": \"%s\", \"AssetID\": \"%s\", \"Registered\": %t }",
		notarizedAsset.MachineID, notarizedAsset.MachineAddress, notarizedAsset.AssetID, notarizedAsset.GetRegistered())
	cmd := exec.Command("planetmint-god", "tx", "machine", "notarize-liquid-asset",
		"--from", sendingValidatorAddress, "-y", "--fees", "1plmnt", obj)

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
