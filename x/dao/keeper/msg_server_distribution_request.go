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

	validatorIdentity, err := util.GetValidatorCometBFTIdentity(ctx, k.RootDir)
	if err != nil {
		return nil, err
	}
	if msg.Distribution.GetProposer() != validatorIdentity {
		util.GetAppLogger().Info(ctx, distributionRequestTag+"Not the proposer. proposer: %s validator identity: %s", msg.Distribution.GetProposer(), validatorIdentity)
		return &types.MsgDistributionRequestResponse{}, nil
	}

	reissuanceAsset := k.GetParams(ctx).ReissuanceAsset
	util.GetAppLogger().Info(ctx, distributionRequestTag+"entering asset distribution mode")
	// issue 5 distributions:
	earlyInvestorTx, err := util.DistributeAsset(msg.Distribution.EarlyInvAddr, msg.Distribution.EarlyInvAmount, reissuanceAsset)
	if err != nil {
		util.GetAppLogger().Error(ctx, distributionRequestTag+"could not distribute asset to early investors: "+err.Error())
	}
	investorTx, err := util.DistributeAsset(msg.Distribution.InvestorAddr, msg.Distribution.InvestorAmount, reissuanceAsset)
	if err != nil {
		util.GetAppLogger().Error(ctx, distributionRequestTag+"could not distribute asset to investors: "+err.Error())
	}
	strategicTx, err := util.DistributeAsset(msg.Distribution.StrategicAddr, msg.Distribution.StrategicAmount, reissuanceAsset)
	if err != nil {
		util.GetAppLogger().Error(ctx, distributionRequestTag+"could not distribute asset to strategic investments: "+err.Error())
	}
	popTx, err := util.DistributeAsset(msg.Distribution.PopAddr, msg.Distribution.PopAmount, reissuanceAsset)
	if err != nil {
		util.GetAppLogger().Error(ctx, distributionRequestTag+"could not distribute asset to PoP: "+err.Error())
	}
	daoTx, err := util.DistributeAsset(msg.Distribution.DaoAddr, msg.Distribution.DaoAmount, reissuanceAsset)
	if err != nil {
		util.GetAppLogger().Error(ctx, distributionRequestTag+"could not distribute asset to DAO: "+err.Error())
	}

	util.SendDistributionResult(goCtx, msg.Distribution.LastPop, daoTx, investorTx, popTx, earlyInvestorTx, strategicTx)

	return &types.MsgDistributionRequestResponse{}, nil
}
