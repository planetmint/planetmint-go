package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) DistributionRequest(goCtx context.Context, msg *types.MsgDistributionRequest) (*types.MsgDistributionRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lastReissuance, found := k.GetLastReIssuance(ctx)
	if !found {
		return nil, errorsmod.Wrap(types.ErrReissuanceNotFound, "for last reissuance height")
	}

	if lastReissuance.TxID == "" {
		return nil, errorsmod.Wrap(types.ErrReissuanceTxIDMissing, "for last reissuance height")
	}

	validatorIdentity, validResult := util.GetValidatorCometBFTIdentity(ctx)
	if validResult && msg.Distribution.GetProposer() == validatorIdentity {
		util.GetAppLogger().Info(ctx, "distribution request: Entering Asset Distribution Mode")
		// issue three distributions:
		investorTx, err := util.DistributeAsset(msg.Distribution.InvestorAddr, msg.Distribution.InvestorAmount)
		if err != nil {
			util.GetAppLogger().Error(ctx, "Distribution Request: could not distribute asset to Investors: ", err.Error())
		}
		popTx, err := util.DistributeAsset(msg.Distribution.PopAddr, msg.Distribution.PopAmount)
		if err != nil {
			util.GetAppLogger().Error(ctx, "Distribution Request: could not distribute asset to PoP: ", err.Error())
		}
		daoTx, err := util.DistributeAsset(msg.Distribution.DaoAddr, msg.Distribution.DaoAmount)
		if err != nil {
			util.GetAppLogger().Error(ctx, "Distribution Request: could not distribute asset to DAO: ", err.Error())
		}

		msg.Distribution.InvestorTxID = investorTx
		msg.Distribution.PopTxID = popTx
		msg.Distribution.DaoTxID = daoTx
		util.SendDistributionResult(goCtx, msg.Distribution.LastPop, daoTx, investorTx, popTx)
	}
	util.GetAppLogger().Info(ctx, "distribution request: storing distribution")
	k.StoreDistributionOrder(ctx, *msg.GetDistribution())

	return &types.MsgDistributionRequestResponse{}, nil
}
