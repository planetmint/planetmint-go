package util

import (
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func ValidateSignature(message string, signature string, publicKey string) (bool, error) {
	// Convert the message, signature, and public key from hex to bytes
	messageBytes, err := hex.DecodeString(message)
	if err != nil {
		return false, errors.New("invalid message hex string")
	}
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false, errors.New("invalid signature hex string")
	}
	publicKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return false, errors.New("invalid public key hex string")
	}

	// Create a secp256k1 public key object
	pubKey := &secp256k1.PubKey{Key: publicKeyBytes}

	// Verify the signature
	isValid := pubKey.VerifySignature(messageBytes, signatureBytes)
	if !isValid {
		return false, errors.New("invalid signature")
	} else {
		return isValid, nil
	}
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
