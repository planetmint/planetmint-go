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

func setConfig(goCtx context.Context) {
	rpcConf := lib.GetConfig()
	ctx := sdk.UnwrapSDKContext(goCtx)
	rpcConf.SetChainID(ctx.ChainID())
}
func buildSignBroadcastTx(goCtx context.Context, loggingContext string, sendingValidatorAddress string, msg sdk.Msg) {
	go func() {
		setConfig(goCtx)
		ctx := sdk.UnwrapSDKContext(goCtx)
		addr := sdk.MustAccAddressFromBech32(sendingValidatorAddress)
		txJSON, err := lib.BuildUnsignedTx(addr, msg)
		if err != nil {
			GetAppLogger().Error(ctx, loggingContext+" build unsigned tx failed: "+err.Error())
			return
		}
		GetAppLogger().Info(ctx, loggingContext+" broadcast tx: "+txJSON)
		_, err = lib.BroadcastTxWithFileLock(addr, msg)
		if err != nil {
			GetAppLogger().Error(ctx, loggingContext+" broadcast tx failed: "+err.Error())
		}
	}()
}

func InitRDDLReissuanceProcess(goCtx context.Context, proposerAddress string, txUnsigned string, blockHeight int64) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create re-issuance proposal")
	msg := daotypes.NewMsgReissueRDDLProposal(sendingValidatorAddress, proposerAddress, txUnsigned, blockHeight)
	buildSignBroadcastTx(goCtx, "initializing RDDL re-issuance", sendingValidatorAddress, msg)
}

func SendRDDLReissuanceResult(goCtx context.Context, proposerAddress string, txID string, blockHeight int64) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create re-issuance result")
	msg := daotypes.NewMsgReissueRDDLResult(sendingValidatorAddress, proposerAddress, txID, blockHeight)
	buildSignBroadcastTx(goCtx, "sending the re-issuance result", sendingValidatorAddress, msg)
}

func SendRDDLDistributionRequest(goCtx context.Context, distribution daotypes.DistributionOrder) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Distribution Request")
	msg := daotypes.NewMsgDistributionRequest(sendingValidatorAddress, &distribution)
	buildSignBroadcastTx(goCtx, "sending the distribution request", sendingValidatorAddress, msg)
}

func SendRDDLDistributionResult(goCtx context.Context, lastPoP string, daoTxID string, invTxID string, popTxID string) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Distribution Result")
	iLastPoP, err := strconv.ParseInt(lastPoP, 10, 64)
	if err != nil {
		ctx.Logger().Error("Distribution Result: preparation failed ", err.Error())
		return
	}
	msg := daotypes.NewMsgDistributionResult(sendingValidatorAddress, iLastPoP, daoTxID, invTxID, popTxID)
	buildSignBroadcastTx(goCtx, "send distribution result", sendingValidatorAddress, msg)
}

func SendLiquidAssetRegistration(goCtx context.Context, notarizedAsset machinetypes.LiquidAsset) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Liquid Asset Registration")
	msg := machinetypes.NewMsgNotarizeLiquidAsset(sendingValidatorAddress, &notarizedAsset)
	buildSignBroadcastTx(goCtx, "Liquid Asset Registration:", sendingValidatorAddress, msg)
}
