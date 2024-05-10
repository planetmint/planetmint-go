package keeper

import (
	"context"
	"encoding/hex"

	"github.com/planetmint/planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterTrustAnchor(goCtx context.Context, msg *types.MsgRegisterTrustAnchor) (*types.MsgRegisterTrustAnchorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isValidTrustAnchorPubkey := validatePublicKey(msg.TrustAnchor.Pubkey)
	if !isValidTrustAnchorPubkey {
		return nil, types.ErrInvalidTrustAnchorKey
	}

	_, _, found := k.GetTrustAnchor(ctx, msg.TrustAnchor.Pubkey)
	if found {
		return nil, types.ErrTrustAnchorAlreadyRegistered
	}

	err := k.StoreTrustAnchor(ctx, *msg.TrustAnchor, false)
	if err != nil {
		return nil, err
	}
	return &types.MsgRegisterTrustAnchorResponse{}, err
}

func validatePublicKey(pubkey string) bool {
	pubkeyBytes, err := hex.DecodeString(pubkey)
	if err != nil {
		return false
	}

	//uncompressed key
	if len(pubkeyBytes) == 64 {
		return true
	}

	// compressed key
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
