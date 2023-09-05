package keeper

import (
	"context"
	"encoding/hex"
	"errors"

	"planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterTrustAnchor(goCtx context.Context, msg *types.MsgRegisterTrustAnchor) (*types.MsgRegisterTrustAnchorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isValidTrustAnchorPubkey := validatePublicKey(msg.TrustAnchor.Pubkey)
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

func validatePublicKey(pubkey string) bool {
	pubkeyBytes, err := hex.DecodeString(pubkey)
	if err != nil {
		return false
	}

	// Check if byte slice has correct length
	if len(pubkeyBytes) != 33 {
		return false
	}

	// Check if byte slice starts with 0x02 or 0x03
	firstByte := pubkeyBytes[0]
	if firstByte != 0x02 && firstByte != 0x03 {
		return false
	}

	return true
}
