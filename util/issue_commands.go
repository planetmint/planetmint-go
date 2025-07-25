package util

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	dertypes "github.com/planetmint/planetmint-go/x/der/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
	"sigs.k8s.io/yaml"
)

func buildSignBroadcastTx(goCtx context.Context, loggingContext string, sendingValidatorAddress string, msg sdk.Msg) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	GetAppLogger().Info(ctx, loggingContext+": "+msg.String())
	TerminationWaitGroup.Add(1)
	go func() {
		defer TerminationWaitGroup.Done()
		ctx := sdk.UnwrapSDKContext(goCtx)
		addr := sdk.MustAccAddressFromBech32(sendingValidatorAddress)
		txJSON, err := lib.BuildUnsignedTx(addr, msg)
		if err != nil {
			GetAppLogger().Error(ctx, err, loggingContext+" build unsigned tx failed for: %v, %v", addr, msg)
			return
		}
		GetAppLogger().Debug(ctx, loggingContext+" unsigned tx: "+txJSON)
		out, err := lib.BroadcastTxWithFileLock(addr, msg)
		if err != nil {
			GetAppLogger().Error(ctx, err, loggingContext+" broadcast tx failed: %v, %v", addr, msg)
			return
		}
		txResponse, err := lib.GetTxResponseFromOut(out)
		if err != nil {
			GetAppLogger().Error(ctx, err, loggingContext+" getting tx response from out failed: %v", out)
			return
		}
		if txResponse.Code == 0 {
			GetAppLogger().Info(ctx, loggingContext+" broadcast tx succeeded")
			return
		}
		txResponseJSON, err := yaml.YAMLToJSON([]byte(txResponse.String()))
		if err != nil {
			GetAppLogger().Error(ctx, err, loggingContext+" converting tx response from yaml to json failed: %v", txResponse)
			return
		}
		GetAppLogger().Info(ctx, loggingContext+" broadcast tx failed: "+string(txResponseJSON))
	}()
}

func SendInitReissuance(goCtx context.Context, proposerAddress string, txUnsigned string, blockHeight int64,
	firstIncludedPop int64, lastIncludedPop int64) {
	sendingValidatorAddress := config.GetConfig().GetNodeAddress()
	msg := daotypes.NewMsgReissueRDDLProposal(sendingValidatorAddress, proposerAddress, txUnsigned, blockHeight,
		firstIncludedPop, lastIncludedPop)
	loggingContext := "reissuance proposal"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendReissuanceResult(goCtx context.Context, proposerAddress string, txID string, blockHeight int64) {
	sendingValidatorAddress := config.GetConfig().GetNodeAddress()
	msg := daotypes.NewMsgReissueRDDLResult(sendingValidatorAddress, proposerAddress, txID, blockHeight)
	loggingContext := "reissuance result"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendDistributionRequest(goCtx context.Context, distribution daotypes.DistributionOrder) {
	sendingValidatorAddress := config.GetConfig().GetNodeAddress()
	msg := daotypes.NewMsgDistributionRequest(sendingValidatorAddress, &distribution)
	loggingContext := "distribution request"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendDistributionResult(goCtx context.Context, lastPoP int64, daoTxID string, invTxID string,
	popTxID string, earlyInvestorTxID string, strategicTxID string) {
	sendingValidatorAddress := config.GetConfig().GetNodeAddress()
	msg := daotypes.NewMsgDistributionResult(sendingValidatorAddress, lastPoP, daoTxID, invTxID, popTxID, earlyInvestorTxID, strategicTxID)
	loggingContext := "distribution result"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendLiquidAssetRegistration(goCtx context.Context, notarizedAsset machinetypes.LiquidAsset) {
	sendingValidatorAddress := config.GetConfig().GetNodeAddress()
	msg := machinetypes.NewMsgNotarizeLiquidAsset(sendingValidatorAddress, &notarizedAsset)
	loggingContext := "notarize liquid asset"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendLiquidDerAssetRegistration(goCtx context.Context, notarizedAsset dertypes.LiquidDerAsset) {
	sendingValidatorAddress := config.GetConfig().GetNodeAddress()
	msg := dertypes.NewMsgNotarizeLiquidDerAsset(sendingValidatorAddress, &notarizedAsset)
	loggingContext := "notarize liquid DER asset"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendInitPoP(goCtx context.Context, challenger string, challengee string, blockHeight int64) {
	sendingValidatorAddress := config.GetConfig().GetNodeAddress()
	msg := daotypes.NewMsgInitPop(sendingValidatorAddress, sendingValidatorAddress, challenger, challengee, blockHeight)
	loggingContext := "PoP"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendUpdateRedeemClaim(goCtx context.Context, beneficiary string, id uint64, txID string) {
	sendingValidatorAddress := config.GetConfig().GetNodeAddress()
	msg := daotypes.NewMsgUpdateRedeemClaim(sendingValidatorAddress, beneficiary, txID, id)
	loggingContext := "redeem claim"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}

func SendTokens(goCtx context.Context, beneficiary sdk.AccAddress, amount uint64, denominator string) {
	sendingValidatorAddress := config.GetConfig().GetNodeAddress()

	coin := sdk.NewCoin(denominator, sdk.NewIntFromUint64(amount))
	coins := sdk.NewCoins(coin)
	orgAddr := sdk.MustAccAddressFromBech32(sendingValidatorAddress)
	msg := banktypes.NewMsgSend(orgAddr, beneficiary, coins)

	loggingContext := "sending " + denominator + " tokens"
	buildSignBroadcastTx(goCtx, loggingContext, sendingValidatorAddress, msg)
}
