package util

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

func buildSignBroadcastTx(goCtx context.Context, sendingValidatorAddress string, msg sdk.Msg) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
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

func InitRDDLReissuanceProcess(goCtx context.Context, proposerAddress string, txUnsigned string, blockHeight int64) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	logger := ctx.Logger()
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Reissuance Proposal")
	msg := daotypes.NewMsgReissueRDDLProposal(sendingValidatorAddress, proposerAddress, txUnsigned, blockHeight)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendRDDLReissuanceResult(goCtx context.Context, proposerAddress string, txID string, blockHeight int64) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := ctx.Logger()
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Reissuance Result")
	msg := daotypes.NewMsgReissueRDDLResult(sendingValidatorAddress, proposerAddress, txID, blockHeight)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendRDDLDistributionRequest(goCtx context.Context, distribution daotypes.DistributionOrder) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := ctx.Logger()
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Distribution Request")
	msg := daotypes.NewMsgDistributionRequest(sendingValidatorAddress, &distribution)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendRDDLDistributionResult(goCtx context.Context, lastPoP string, daoTxID string, invTxID string, popTxID string) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := ctx.Logger()
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	logger.Debug("REISSUE: create Distribution Result")
	iLastPoP, err := strconv.ParseInt(lastPoP, 10, 64)
	if err != nil {
		return
	}
	msg := daotypes.NewMsgDistributionResult(sendingValidatorAddress, iLastPoP, daoTxID, invTxID, popTxID)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}
