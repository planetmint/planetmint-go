package keeper

import (
	"context"
	"errors"

	"planetmint-go/util"
	"planetmint-go/x/asset/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NotarizeAsset(goCtx context.Context, msg *types.MsgNotarizeAsset) (*types.MsgNotarizeAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.machineKeeper.GetMachineIndex(ctx, msg.PubKey)

	if !found {
		return nil, errors.New("machine not found")
	}
	hex_pub_key, err := util.GetHexPubKey(msg.PubKey)
	if err != nil {
		return nil, errors.New("could not convert xpub key to hex pub key")
	}
	valid := util.ValidateSignature(msg.Hash, msg.Signature, hex_pub_key)
	if !valid {
		return nil, errors.New("invalid signature")
	}

	var asset = types.Asset{
		Hash:      msg.Hash,
		Signature: msg.Signature,
		Pubkey:    msg.PubKey,
	}

	k.StoreAsset(ctx, asset)

	return &types.MsgNotarizeAssetResponse{}, nil
}
