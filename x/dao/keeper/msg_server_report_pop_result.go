package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReportPopResult(goCtx context.Context, msg *types.MsgReportPopResult) (*types.MsgReportPopResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgReportPopResultResponse{}, nil
}
