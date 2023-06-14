package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"planetmint-go/x/asset/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NotarizeAsset(goCtx context.Context, msg *types.MsgNotarizeAsset) (*types.MsgNotarizeAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	machine, found := k.machineKeeper.GetMachine(ctx, msg.Creator)

	if !found {
		return &types.MsgNotarizeAssetResponse{}, errors.New("machine not found")
	}

	valid := ValidateSignature(msg.CidHash, msg.Sign, msg.Creator)
	if !valid {
		return &types.MsgNotarizeAssetResponse{}, errors.New("invalid signature")
	}

	var asset = types.Asset{
		Hash:      msg.CidHash,
		Signature: msg.Sign,
		Pubkey:    machine.IssuerPlanetmint,
	}

	k.StoreAsset(ctx, asset)

	return &types.MsgNotarizeAssetResponse{}, nil
}

func ValidateSignature(message string, signature string, publicKey string) bool {
	// Convert the message, signature, and public key from hex to bytes
	messageBytes, _ := hex.DecodeString(message)
	signatureBytes, _ := hex.DecodeString(signature)
	publicKeyBytes, _ := hex.DecodeString(publicKey)

	// Hash the message
	hash := sha256.Sum256(messageBytes)

	// Create a secp256k1 public key object
	pubKey := &ed25519.PubKey{Key: publicKeyBytes}

	// Verify the signature
	isValid := pubKey.VerifySignature(hash[:], signatureBytes)

	return isValid
}
