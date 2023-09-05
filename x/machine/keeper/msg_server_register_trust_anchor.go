package keeper

import (
	"context"
	"errors"

	config "planetmint-go/config"
	"planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterTrustAnchor(goCtx context.Context, msg *types.MsgRegisterTrustAnchor) (*types.MsgRegisterTrustAnchorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isValidTrustAnchorPubkey := validateExtendedPublicKey(msg.TrustAnchor.Pubkey, config.LiquidNetParams)
	if !isValidTrustAnchorPubkey {
		return nil, errors.New("invalid trust anchor pubkey")
	}

	_, _, found := k.GetTrustAnchor(ctx, msg.TrustAnchor.Pubkey)
	if found {
		return nil, errors.New("trust anchor is already registered")
	}

	k.StoreTrustAnchor(ctx, *msg.TrustAnchor, false)

	return &types.MsgRegisterTrustAnchorResponse{}, nil
}
