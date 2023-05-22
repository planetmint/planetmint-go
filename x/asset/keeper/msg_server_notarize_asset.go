package keeper

import (
	"context"

	"planetmint-go/x/asset/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NotarizeAsset(goCtx context.Context, msg *types.MsgNotarizeAsset) (*types.MsgNotarizeAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// CHECK IF MSG CREATOR (pub_key) IS ATTESTED MACHINE

	// CHECK SHORTENED URL FOR NODE

	// STORE CID_HASH SIGNATURE PUBLIC KEY

	return &types.MsgNotarizeAssetResponse{}, nil
}
