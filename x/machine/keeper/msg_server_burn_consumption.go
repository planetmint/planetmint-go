package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

func (k msgServer) BurnConsumption(goCtx context.Context, msg *types.MsgBurnConsumption) (*types.MsgBurnConsumptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBurnConsumptionResponse{}, nil
}
