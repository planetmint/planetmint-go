package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"planetmint-go/x/asset/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NotarizeAsset(goCtx context.Context, msg *types.MsgNotarizeAsset) (*types.MsgNotarizeAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.machineKeeper.GetMachineIndex(ctx, msg.PubKey)

	if !found {
		return &types.MsgNotarizeAssetResponse{}, errors.New("machine not found")
	}

	valid := ValidateSignature(msg.Hash, msg.Signature, msg.PubKey)
	if !valid {
		return &types.MsgNotarizeAssetResponse{}, errors.New("invalid signature")
	}

	var asset = types.Asset{
		Hash:      msg.Hash,
		Signature: msg.Signature,
		Pubkey:    msg.PubKey,
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
	pubKey := &secp256k1.PubKey{Key: publicKeyBytes}

	// Verify the signature
	isValid := pubKey.VerifySignature(hash[:], signatureBytes)

	return isValid
}
