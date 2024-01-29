package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

var (
	distributionRequestTag = "distribution request: "
)

func (k msgServer) DistributionRequest(goCtx context.Context, msg *types.MsgDistributionRequest) (*types.MsgDistributionRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lastReissuance, found := k.GetLastReissuance(ctx)
	if !found {
		return nil, errorsmod.Wrap(types.ErrReissuanceNotFound, "for last reissuance height")
	}

	if lastReissuance.TxID == "" {
		return nil, errorsmod.Wrap(types.ErrReissuanceTxIDMissing, "for last reissuance height")
	}

	util.GetAppLogger().Info(ctx, distributionRequestTag+"storing distribution: "+msg.GetDistribution().String())
	k.StoreDistributionOrder(ctx, *msg.GetDistribution())

	validatorIdentity, validResult := util.GetValidatorCometBFTIdentity(ctx)
	if validResult && msg.Distribution.GetProposer() == validatorIdentity {
		reissuanceAsset := k.GetParams(ctx).ReissuanceAsset
		util.GetAppLogger().Info(ctx, distributionRequestTag+"entering asset distribution mode")
		// issue three distributions:
		investorTx, err := util.DistributeAsset(msg.Distribution.InvestorAddr, msg.Distribution.InvestorAmount, reissuanceAsset)
		if err != nil {
			util.GetAppLogger().Error(ctx, distributionRequestTag+"could not distribute asset to investors: "+err.Error())
		}
		popTx, err := util.DistributeAsset(msg.Distribution.PopAddr, msg.Distribution.PopAmount, reissuanceAsset)
		if err != nil {
			util.GetAppLogger().Error(ctx, distributionRequestTag+"could not distribute asset to PoP: "+err.Error())
		}
		daoTx, err := util.DistributeAsset(msg.Distribution.DaoAddr, msg.Distribution.DaoAmount, reissuanceAsset)
		if err != nil {
			util.GetAppLogger().Error(ctx, distributionRequestTag+"could not distribute asset to DAO: "+err.Error())
		}

		msg.Distribution.InvestorTxID = investorTx
		msg.Distribution.PopTxID = popTx
		msg.Distribution.DaoTxID = daoTx
		util.SendDistributionResult(goCtx, msg.Distribution.LastPop, daoTx, investorTx, popTx)
	} else {
		util.GetAppLogger().Error(ctx, distributionRequestTag+"failed. valid result: %v proposer: %s validator identity: %s", validResult, msg.Distribution.GetProposer(), validatorIdentity)
	}

	return &types.MsgDistributionRequestResponse{}, nil
}
