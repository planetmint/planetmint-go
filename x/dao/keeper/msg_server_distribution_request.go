package keeper

import (
	"context"

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

		}
		pop_tx, err := util.DistributeAsset(msg.Distribution.PopAddr, msg.Distribution.PopAmount)
		if err != nil {

		}
		dao_tx, err := util.DistributeAsset(msg.Distribution.DaoAddr, msg.Distribution.DaoAmount)
		if err != nil {

		}
		//TODO: should be handled via a Result value distribution
		msg.Distribution.InvestorTxid = investor_tx
		msg.Distribution.PopTxid = pop_tx
		msg.Distribution.DaoTxid = dao_tx
	}
	k.StoreDistributionOrder(ctx, *msg.GetDistribution())

	return &types.MsgDistributionRequestResponse{}, nil
}
