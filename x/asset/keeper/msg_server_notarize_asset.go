package keeper

import (
	"context"

	"github.com/planetmint/planetmint-go/x/asset/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NotarizeAsset(goCtx context.Context, msg *types.MsgNotarizeAsset) (*types.MsgNotarizeAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//	var asset = types.Asset{Cid: msg.Hash}

	k.StoreAsset(ctx, *msg)

	return &types.MsgNotarizeAssetResponse{}, nil
}
