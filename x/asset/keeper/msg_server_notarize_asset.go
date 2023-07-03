package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"planetmint-go/x/asset/types"
)

func (k msgServer) NotarizeAsset(goCtx context.Context, msg *types.MsgNotarizeAsset) (*types.MsgNotarizeAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgNotarizeAssetResponse{}, nil
}
