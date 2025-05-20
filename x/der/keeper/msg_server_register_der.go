package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/der/types"
)

func (k msgServer) RegisterDER(goCtx context.Context, msg *types.MsgRegisterDER) (*types.MsgRegisterDERResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterDERResponse{}, nil
}
