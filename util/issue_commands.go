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
	ctx := sdk.UnwrapSDKContext(goCtx)
	GetAppLogger().Info(ctx, loggingContext+": "+msg.String())
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
		if txResponse.Code == 0 {
			GetAppLogger().Info(ctx, loggingContext+" broadcast tx succeeded")
			return
		}
		txResponseJSON, err := yaml.YAMLToJSON([]byte(txResponse.String()))
		if err != nil {
			GetAppLogger().Error(ctx, loggingContext+" converting tx response from yaml to json failed: "+err.Error())
			return
		}
		GetAppLogger().Error(ctx, loggingContext+" broadcast tx failed: "+string(txResponseJSON))
	}()
}

func SendInitReissuance(goCtx context.Context, proposerAddress string, txUnsigned string, blockHeight int64,
	firstIncludedPop int64, lastIncludedPop int64) {
	// get_last_PoPBlockHeight() // TODO: to be read form the upcoming PoP-store
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgReissueRDDLProposal(sendingValidatorAddress, proposerAddress, txUnsigned, blockHeight,
		firstIncludedPop, lastIncludedPop)
	loggingContext := "reissuance proposal"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendReissuanceResult(goCtx context.Context, proposerAddress string, txID string, blockHeight int64) {
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgReissueRDDLResult(sendingValidatorAddress, proposerAddress, txID, blockHeight)
	loggingContext := "reissuance result"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendDistributionRequest(goCtx context.Context, distribution daotypes.DistributionOrder) {
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgDistributionRequest(sendingValidatorAddress, &distribution)
	loggingContext := "distribution request"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendDistributionResult(goCtx context.Context, lastPoP int64, daoTxID string, invTxID string, popTxID string) {
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgDistributionResult(sendingValidatorAddress, lastPoP, daoTxID, invTxID, popTxID)
	loggingContext := "distribution result"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendLiquidAssetRegistration(goCtx context.Context, notarizedAsset machinetypes.LiquidAsset) {
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := machinetypes.NewMsgNotarizeLiquidAsset(sendingValidatorAddress, &notarizedAsset)
	loggingContext := "notarize liquid asset"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendInitPoP(goCtx context.Context, proposer string, challenger string, challengee string, blockHeight int64) {
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgInitPop(sendingValidatorAddress, proposer, challenger, challengee, blockHeight)
	loggingContext := "PoP"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendUpdateRedeemClaim(goCtx context.Context, beneficiary string, id uint64, txID string) {
	sendingValidatorAddress := config.GetConfig().ValidatorAddress
	msg := daotypes.NewMsgUpdateRedeemClaim(sendingValidatorAddress, beneficiary, txID, id)
	loggingContext := "redeem claim"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}
