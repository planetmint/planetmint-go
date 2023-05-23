package keeper

import (
	"context"

	"planetmint-go/x/asset/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NotarizeAsset(goCtx context.Context, msg *types.MsgNotarizeAsset) (*types.MsgNotarizeAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// CHECK IF MSG CREATOR (pub_key) IS ATTESTED MACHINE
	var asset = types.Asset{
		Hash:      msg.CidHash,
		Signature: msg.Sign,
		Pubkey:    msg.Creator,
	}

	// CHECK LOCATION FOR NODE

	// STORE CID_HASH SIGNATURE PUBLIC KEY
	k.StoreAsset(ctx, asset)

	return &types.MsgNotarizeAssetResponse{}, nil
}
