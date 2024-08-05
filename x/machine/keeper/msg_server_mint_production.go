package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

func (k msgServer) MintProduction(goCtx context.Context, msg *types.MsgMintProduction) (*types.MsgMintProductionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMintProductionResponse{}, nil
}
