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

	validator_identity, valid_result := util.GetValidatorCometBFTIdentity(ctx)
	if valid_result && msg.Distribution.GetProposer() == validator_identity {
		// issue three distributions:
		investor_tx, err := util.DistributeAsset(msg.Distribution.InvestorAddr, msg.Distribution.InvestorAmount)
		if err != nil {
			ctx.Logger().Error("Distribution Request: could not distribute asset to Investors: ", err.Error())
		}
		pop_tx, err := util.DistributeAsset(msg.Distribution.PopAddr, msg.Distribution.PopAmount)
		if err != nil {
			ctx.Logger().Error("Distribution Request: could not distribute asset to PoP: ", err.Error())
		}
		dao_tx, err := util.DistributeAsset(msg.Distribution.DaoAddr, msg.Distribution.DaoAmount)
		if err != nil {
			ctx.Logger().Error("Distribution Request: could not distribute asset to DAO: ", err.Error())
		}

		msg.Distribution.InvestorTxid = investor_tx
		msg.Distribution.PopTxid = pop_tx
		msg.Distribution.DaoTxid = dao_tx
		last_pop_str := strconv.FormatUint(msg.Distribution.LastPop, 10)
		err = util.SendRDDLDistributionResult(ctx, last_pop_str, dao_tx, investor_tx, pop_tx)
		if err != nil {
			ctx.Logger().Error("Distribution Request: could not send distribution result ", err.Error())
		}
	}
	k.StoreDistributionOrder(ctx, *msg.GetDistribution())

	return &types.MsgDistributionRequestResponse{}, nil
}
