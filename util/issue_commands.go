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
		return
	}
	GetAppLogger().Info(ctx, "broadcast tx: "+txJSON)
	_, err = lib.BroadcastTxWithFileLock(goCtx, addr, msg)
	return
}

func InitRDDLReissuanceProcess(goCtx context.Context, proposerAddress string, txUnsigned string, blockHeight int64) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Reissuance Proposal")
	msg := daotypes.NewMsgReissueRDDLProposal(sendingValidatorAddress, proposerAddress, txUnsigned, blockHeight)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendRDDLReissuanceResult(goCtx context.Context, proposerAddress string, txID string, blockHeight int64) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Reissuance Result")
	msg := daotypes.NewMsgReissueRDDLResult(sendingValidatorAddress, proposerAddress, txID, blockHeight)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendRDDLDistributionRequest(goCtx context.Context, distribution daotypes.DistributionOrder) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Distribution Request")
	msg := daotypes.NewMsgDistributionRequest(sendingValidatorAddress, &distribution)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendRDDLDistributionResult(goCtx context.Context, lastPoP string, daoTxID string, invTxID string, popTxID string) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Distribution Result")
	iLastPoP, err := strconv.ParseInt(lastPoP, 10, 64)
	if err != nil {
		return
	}
	msg := daotypes.NewMsgDistributionResult(sendingValidatorAddress, iLastPoP, daoTxID, invTxID, popTxID)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}

func SendLiquidAssetRegistration(goCtx context.Context, notarizedAsset machinetypes.LiquidAsset) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Liquid Asset Registration")
	msg := machinetypes.NewMsgNotarizeLiquidAsset(sendingValidatorAddress, &notarizedAsset)
	err = buildSignBroadcastTx(goCtx, sendingValidatorAddress, msg)
	return
}
