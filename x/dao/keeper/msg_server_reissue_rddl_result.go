package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReissueRDDLResult(goCtx context.Context, msg *types.MsgReissueRDDLResult) (*types.MsgReissueRDDLResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgReissueRDDLResultResponse{}, nil
}
