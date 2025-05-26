package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/der/types"
)

func (k msgServer) NotarizeLiquidDerAsset(goCtx context.Context, msg *types.MsgNotarizeLiquidDerAsset) (*types.MsgNotarizeLiquidDerAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.StoreLiquidDerAsset(ctx, *msg.DerAsset)

	return &types.MsgNotarizeLiquidDerAssetResponse{}, nil
}
