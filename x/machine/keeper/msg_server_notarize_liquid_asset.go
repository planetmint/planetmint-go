package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

func (k msgServer) NotarizeLiquidAsset(goCtx context.Context, msg *types.MsgNotarizeLiquidAsset) (*types.MsgNotarizeLiquidAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.StoreLiquidAttest(ctx, *msg.GetNotarization())

	return &types.MsgNotarizeLiquidAssetResponse{}, nil
}
