package keeper

import (
	"context"
	"errors"

	"planetmint-go/x/asset/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NotarizeAsset(goCtx context.Context, msg *types.MsgNotarizeAsset) (*types.MsgNotarizeAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	machine, found := k.machineKeeper.GetMachine(ctx, msg.Creator)

	if !found {
		return &types.MsgNotarizeAssetResponse{}, errors.New("machine not found")
	}

	// TODO: validate signature

	var asset = types.Asset{
		Hash:      msg.CidHash,
		Signature: msg.Sign,
		Pubkey:    machine.IssuerPlanetmint,
	}

	k.StoreAsset(ctx, asset)

	return &types.MsgNotarizeAssetResponse{}, nil
}
