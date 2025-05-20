package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/der/types"
)

func (k msgServer) RegisterDER(goCtx context.Context, msg *types.MsgRegisterDER) (*types.MsgRegisterDERResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.StoreDerAttest(ctx, *msg.Der)

	//TODO: init NFT creation and storag of NFT to DER associations

	return &types.MsgRegisterDERResponse{}, nil
}
