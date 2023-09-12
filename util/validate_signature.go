package util

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func ValidateSignature(message string, signature string, publicKey string) bool {
	// Convert the message, signature, and public key from hex to bytes
	messageBytes, _ := hex.DecodeString(message)
	signatureBytes, _ := hex.DecodeString(signature)
	publicKeyBytes, _ := hex.DecodeString(publicKey)

	// Create a secp256k1 public key object
	pubKey := &secp256k1.PubKey{Key: publicKeyBytes}

	// Verify the signature
	isValid := pubKey.VerifySignature(messageBytes, signatureBytes)

	return isValid
}

func GetHexPubKey(ext_pub_key string) (string, error) {
	xpubKey, err := hdkeychain.NewKeyFromString(ext_pub_key)
	if err != nil {
		return "", err
	}
	pubKey, err := xpubKey.ECPubKey()
	if err != nil {
		return "", err
	}
	byte_key := pubKey.SerializeCompressed()
	return hex.EncodeToString(byte_key), nil
}
