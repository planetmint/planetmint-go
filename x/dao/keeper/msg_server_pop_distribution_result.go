package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) PopDistributionResult(goCtx context.Context, msg *types.MsgPopDistributionResult) (*types.MsgPopDistributionResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	distribution, found := k.LookupDistributionOrder(ctx, msg.GetLastPop())
	if found {
		distribution.DaoTxid = msg.DaoTx
		distribution.PopTxid = msg.PopTx
		distribution.InvestorTxid = msg.InvestorTx
		k.StoreDistributionOrder(ctx, distribution)
	}

	return &types.MsgPopDistributionResultResponse{}, nil
}
