package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
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
	} else {
		return nil, errorsmod.Wrapf(types.ErrDistributionNotFound, " for provided block height %s", strconv.FormatUint(msg.GetLastPop(), 10))
	}

	return &types.MsgDistributionResultResponse{}, nil
}
