package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) DistributionRequest(goCtx context.Context, msg *types.MsgDistributionRequest) (*types.MsgDistributionRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validatorIdentity, validResult := util.GetValidatorCometBFTIdentity(ctx)
	if validResult && msg.Distribution.GetProposer() == validatorIdentity {
		// issue three distributions:
		investorTx, err := util.DistributeAsset(msg.Distribution.InvestorAddr, msg.Distribution.InvestorAmount)
		if err != nil {
			ctx.Logger().Error("Distribution Request: could not distribute asset to Investors: ", err.Error())
		}
		popTx, err := util.DistributeAsset(msg.Distribution.PopAddr, msg.Distribution.PopAmount)
		if err != nil {
			ctx.Logger().Error("Distribution Request: could not distribute asset to PoP: ", err.Error())
		}
		daoTx, err := util.DistributeAsset(msg.Distribution.DaoAddr, msg.Distribution.DaoAmount)
		if err != nil {
			ctx.Logger().Error("Distribution Request: could not distribute asset to DAO: ", err.Error())
		}

		msg.Distribution.InvestorTxid = investorTx
		msg.Distribution.PopTxid = popTx
		msg.Distribution.DaoTxid = daoTx
		lastPopString := strconv.FormatInt(msg.Distribution.LastPop, 10)
		err = util.SendRDDLDistributionResult(ctx, lastPopString, daoTx, investorTx, popTx)
		if err != nil {
			ctx.Logger().Error("Distribution Request: could not send distribution result ", err.Error())
		}
	}
	k.StoreDistributionOrder(ctx, *msg.GetDistribution())

	return &types.MsgDistributionRequestResponse{}, nil
}
