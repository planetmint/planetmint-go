package util

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
	"sigs.k8s.io/yaml"
)

func buildSignBroadcastTx(goCtx context.Context, loggingContext string, sendingValidatorAddress string, msg sdk.Msg) {
	go func() {
		ctx := sdk.UnwrapSDKContext(goCtx)
		addr := sdk.MustAccAddressFromBech32(sendingValidatorAddress)
		txJSON, err := lib.BuildUnsignedTx(addr, msg)
		if err != nil {
			GetAppLogger().Error(ctx, loggingContext+" build unsigned tx failed: "+err.Error())
			return
		}
		GetAppLogger().Debug(ctx, loggingContext+" unsigned tx: "+txJSON)
		out, err := lib.BroadcastTxWithFileLock(addr, msg)
		if err != nil {
			GetAppLogger().Error(ctx, loggingContext+" broadcast tx failed: "+err.Error())
			return
		}
		txResponse, err := lib.GetTxResponseFromOut(out)
		if err != nil {
			GetAppLogger().Error(ctx, loggingContext+" getting tx response from out failed: "+err.Error())
			return
		}
		txResponseJSON, err := yaml.YAMLToJSON([]byte(txResponse.String()))
		if err != nil {
			GetAppLogger().Error(ctx, loggingContext+" converting tx response from yaml to json failed: "+err.Error())
			return
		}
		GetAppLogger().Info(ctx, loggingContext+" broadcast tx succeeded: "+string(txResponseJSON))
	}()
}

func SendInitReissuance(goCtx context.Context, proposerAddress string, txUnsigned string, blockHeight int64,
	firstIncludedPop int64, lastIncludedPop int64) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgReissueRDDLProposal(sendingValidatorAddress, proposerAddress, txUnsigned, blockHeight,
		firstIncludedPop, lastIncludedPop)
	GetAppLogger().Info(ctx, "create re-issuance proposal: "+msg.String())
	buildSignBroadcastTx(goCtx, "initializing RDDL re-issuance", sendingValidatorAddress, msg)
}

func SendReissuanceResult(goCtx context.Context, proposerAddress string, txID string, blockHeight int64) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgReissueRDDLResult(sendingValidatorAddress, proposerAddress, txID, blockHeight)
	GetAppLogger().Info(ctx, "create re-issuance result: "+msg.String())
	buildSignBroadcastTx(goCtx, "sending the re-issuance result", sendingValidatorAddress, msg)
}

func SendDistributionRequest(goCtx context.Context, distribution daotypes.DistributionOrder) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgDistributionRequest(sendingValidatorAddress, &distribution)
	GetAppLogger().Info(ctx, "create Distribution Request: "+msg.String())
	buildSignBroadcastTx(goCtx, "sending the distribution request", sendingValidatorAddress, msg)
}

func SendDistributionResult(goCtx context.Context, lastPoP int64, daoTxID string, invTxID string, popTxID string) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgDistributionResult(sendingValidatorAddress, lastPoP, daoTxID, invTxID, popTxID)
	GetAppLogger().Info(ctx, "create Distribution Result: "+msg.String())
	buildSignBroadcastTx(goCtx, "send distribution result", sendingValidatorAddress, msg)
}

func SendLiquidAssetRegistration(goCtx context.Context, notarizedAsset machinetypes.LiquidAsset) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := machinetypes.NewMsgNotarizeLiquidAsset(sendingValidatorAddress, &notarizedAsset)
	GetAppLogger().Info(ctx, "create Liquid Asset Registration: "+msg.String())
	buildSignBroadcastTx(goCtx, "Liquid Asset Registration:", sendingValidatorAddress, msg)
}

func SendInitPoP(goCtx context.Context, proposer string, challenger string, challengee string, blockHeight int64) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgInitPop(sendingValidatorAddress, proposer, challenger, challengee, blockHeight)
	GetAppLogger().Info(ctx, "create Init PoP message: "+msg.String())
	buildSignBroadcastTx(goCtx, "Init PoP:", sendingValidatorAddress, msg)
}
