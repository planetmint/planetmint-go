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
	return ValidateSignatureByteMsg(messageBytes, signature, publicKey)
}

func ValidateSignatureByteMsg(message []byte, signature string, publicKey string) (bool, error) {
	// Convert  signature, and public key from hex to bytes
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
	isValid := pubKey.VerifySignature(message, signatureBytes)
	if !isValid {
		return false, errors.New("invalid signature")
	}
	return isValid, nil
}

func GetHexPubKey(extPubKey string) (string, error) {
	xpubKey, err := hdkeychain.NewKeyFromString(extPubKey)
	if err != nil {
		return "", err
	}
	pubKey, err := xpubKey.ECPubKey()
	if err != nil {
		return "", err
	}
	byteKey := pubKey.SerializeCompressed()
	return hex.EncodeToString(byteKey), nil
}
