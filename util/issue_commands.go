package util

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
)

func buildSignBroadcastTx(goCtx context.Context, sendingValidatorAddress string, msg sdk.Msg) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	addr := sdk.MustAccAddressFromBech32(sendingValidatorAddress)
	txJSON, err := lib.BuildUnsignedTx(goCtx, addr, msg)
	if err != nil {
		GetAppLogger().Error(ctx, "broadcast tx: failed")
		return
	}
	GetAppLogger().Info(ctx, "broadcast tx: "+txJSON)
	_, err = lib.BroadcastTxWithFileLock(goCtx, addr, msg)
	return
}

// TODO check if we can convert this calls parameter to a SDKContext
func InitRDDLReissuanceProcess(goCtx context.Context, proposerAddress string, txUnsigned string, blockHeight int64) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "reissuance: create Proposal")
	msg := daotypes.NewMsgReissueRDDLProposal(sendingValidatorAddress, proposerAddress, txUnsigned, blockHeight)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendRDDLReissuanceResult(goCtx context.Context, proposerAddress string, txID string, blockHeight int64) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "reissuance: create Result")
	msg := daotypes.NewMsgReissueRDDLResult(sendingValidatorAddress, proposerAddress, txID, blockHeight)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

// TODO check if we can convert this calls parameter to a SDKContext
func SendRDDLDistributionRequest(goCtx context.Context, distribution daotypes.DistributionOrder) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "distribution: create request")
	msg := daotypes.NewMsgDistributionRequest(sendingValidatorAddress, &distribution)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendRDDLDistributionResult(goCtx context.Context, lastPoP string, daoTxID string, invTxID string, popTxID string) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	iLastPoP, err := strconv.ParseInt(lastPoP, 10, 64)
	if err != nil {
		GetAppLogger().Error(ctx, "distribution: result conversion failed")
		return
	}
	GetAppLogger().Info(ctx, "distribution: create result")
	msg := daotypes.NewMsgDistributionResult(sendingValidatorAddress, iLastPoP, daoTxID, invTxID, popTxID)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendLiquidAssetRegistration(goCtx context.Context, notarizedAsset machinetypes.LiquidAsset) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "liquid asset registration: notarize result")
	msg := machinetypes.NewMsgNotarizeLiquidAsset(sendingValidatorAddress, &notarizedAsset)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}
