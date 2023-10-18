package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) DistributionResult(goCtx context.Context, msg *types.MsgDistributionResult) (*types.MsgDistributionResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	distribution, found := k.LookupDistributionOrder(ctx, msg.GetLastPop())
	if found {
		distribution.DaoTxid = msg.DaoTxid
		distribution.PopTxid = msg.PopTxid
		distribution.InvestorTxid = msg.InvestorTxid
		k.StoreDistributionOrder(ctx, distribution)
	}

	return &types.MsgDistributionResultResponse{}, nil
}
