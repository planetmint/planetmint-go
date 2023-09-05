package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"planetmint-go/x/machine/types"
)

func (k msgServer) RegisterTrustAnchor(goCtx context.Context, msg *types.MsgRegisterTrustAnchor) (*types.MsgRegisterTrustAnchorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterTrustAnchorResponse{}, nil
}
