package util

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
)

func setRPCConfig(goCtx context.Context) {
	rpcConf := lib.GetConfig()
	ctx := sdk.UnwrapSDKContext(goCtx)
	rpcConf.SetChainID(ctx.ChainID())
}
func buildSignBroadcastTx(goCtx context.Context, loggingContext string, sendingValidatorAddress string, msg sdk.Msg) {
	go func() {
		setRPCConfig(goCtx)
		ctx := sdk.UnwrapSDKContext(goCtx)
		addr := sdk.MustAccAddressFromBech32(sendingValidatorAddress)
		txJSON, err := lib.BuildUnsignedTx(addr, msg)
		if err != nil {
			GetAppLogger().Error(ctx, loggingContext+" build unsigned tx failed: "+err.Error())
			return
		}
		GetAppLogger().Info(ctx, loggingContext+" unsigned tx: "+txJSON)
		_, err = lib.BroadcastTxWithFileLock(addr, msg)
		if err != nil {
			GetAppLogger().Error(ctx, loggingContext+" broadcast tx failed: "+err.Error())
			return
		}
		GetAppLogger().Info(ctx, loggingContext+" broadcast tx succeeded")
	}()
}

func SendInitReissuance(goCtx context.Context, proposerAddress string, txUnsigned string, blockHeight int64,
	firstIncludedPop int64, lastIncludedPop int64) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create re-issuance proposal")
	msg := daotypes.NewMsgReissueRDDLProposal(sendingValidatorAddress, proposerAddress, txUnsigned, blockHeight,
		firstIncludedPop, lastIncludedPop)
	buildSignBroadcastTx(goCtx, "initializing RDDL re-issuance", sendingValidatorAddress, msg)
}

func SendReissuanceResult(goCtx context.Context, proposerAddress string, txID string, blockHeight int64) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create re-issuance result")
	msg := daotypes.NewMsgReissueRDDLResult(sendingValidatorAddress, proposerAddress, txID, blockHeight)
	buildSignBroadcastTx(goCtx, "sending the re-issuance result", sendingValidatorAddress, msg)
}

func SendDistributionRequest(goCtx context.Context, distribution daotypes.DistributionOrder) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Distribution Request")
	msg := daotypes.NewMsgDistributionRequest(sendingValidatorAddress, &distribution)
	buildSignBroadcastTx(goCtx, "sending the distribution request", sendingValidatorAddress, msg)
}

func SendDistributionResult(goCtx context.Context, lastPoP int64, daoTxID string, invTxID string, popTxID string) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Distribution Result")
	msg := daotypes.NewMsgDistributionResult(sendingValidatorAddress, lastPoP, daoTxID, invTxID, popTxID)
	buildSignBroadcastTx(goCtx, "send distribution result", sendingValidatorAddress, msg)
}

func SendLiquidAssetRegistration(goCtx context.Context, notarizedAsset machinetypes.LiquidAsset) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Liquid Asset Registration")
	msg := machinetypes.NewMsgNotarizeLiquidAsset(sendingValidatorAddress, &notarizedAsset)
	buildSignBroadcastTx(goCtx, "Liquid Asset Registration:", sendingValidatorAddress, msg)
}

func SendInitPoP(goCtx context.Context, proposer string, challenger string, challengee string, blockHeight int64) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	GetAppLogger().Info(ctx, "create Init PoP message")
	msg := daotypes.NewMsgInitPop(sendingValidatorAddress, proposer, challenger, challengee, blockHeight)
	buildSignBroadcastTx(goCtx, "Init PoP:", sendingValidatorAddress, msg)
}
