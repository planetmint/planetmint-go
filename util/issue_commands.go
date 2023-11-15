package util

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

func buildSignBroadcastTx(ctx sdk.Context, sendingValidatorAddress string, msg sdk.Msg) (err error) {
	logger := ctx.Logger()
	addr := sdk.MustAccAddressFromBech32(sendingValidatorAddress)
	txBytes, txJSON, err := lib.BuildAndSignTx(addr, msg)
	if err != nil {
		return
	}
	logger.Debug("REISSUE: tx: " + txJSON)
	_, err = lib.BroadcastTx(txBytes)
	return
}

func InitRDDLReissuanceProcess(ctx sdk.Context, proposerAddress string, txUnsigned string, blockHeight int64) (err error) {
	//get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	logger := ctx.Logger()
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Reissuance Proposal")
	msg := daotypes.NewMsgReissueRDDLProposal(sendingValidatorAddress, proposerAddress, txUnsigned, blockHeight)
	err = buildSignBroadcastTx(ctx, sendingValidatorAddress, msg)
	return
}

func SendRDDLReissuanceResult(ctx sdk.Context, proposerAddress string, txID string, blockHeight int64) (err error) {
	logger := ctx.Logger()
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Reissuance Result")
	msg := daotypes.NewMsgReissueRDDLResult(sendingValidatorAddress, proposerAddress, txID, blockHeight)
	err = buildSignBroadcastTx(ctx, sendingValidatorAddress, msg)
	return
}

func SendRDDLDistributionRequest(ctx sdk.Context, distribution daotypes.DistributionOrder) (err error) {
	logger := ctx.Logger()
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Distribution Request")
	msg := daotypes.NewMsgDistributionRequest(sendingValidatorAddress, &distribution)
	err = buildSignBroadcastTx(ctx, sendingValidatorAddress, msg)
	return
}

func SendRDDLDistributionResult(ctx sdk.Context, lastPoP string, daoTxID string, invTxID string, popTxID string) (err error) {
	logger := ctx.Logger()
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Distribution Result")
	iLastPoP, err := strconv.ParseInt(lastPoP, 10, 64)
	if err != nil {
		return
	}
	msg := daotypes.NewMsgDistributionResult(sendingValidatorAddress, iLastPoP, daoTxID, invTxID, popTxID)
	err = buildSignBroadcastTx(ctx, sendingValidatorAddress, msg)
	return
}
