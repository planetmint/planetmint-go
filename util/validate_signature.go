package util

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func ValidateSignature(message string, signature string, publicKey string) bool {
	// Convert the message, signature, and public key from hex to bytes
	messageBytes := []byte(message)
	signatureBytes, _ := hex.DecodeString(signature)
	publicKeyBytes, _ := hex.DecodeString(publicKey)

	// Create a secp256k1 public key object
	pubKey := &secp256k1.PubKey{Key: publicKeyBytes}

	// Verify the signature
	isValid := pubKey.VerifySignature(messageBytes, signatureBytes)

	return isValid
}
