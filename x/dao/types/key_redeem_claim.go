package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RedeemClaimKeyPrefix is the prefix to retrieve all RedeemClaim
	RedeemClaimKeyPrefix = "RedeemClaim/value/"
)

// RedeemClaimKey returns the store key to retrieve a RedeemClaim from the index fields
func RedeemClaimKey(
	beneficiary string,
	liquidTxHash string,
) []byte {
	var key []byte

	beneficiaryBytes := []byte(beneficiary)
	key = append(key, beneficiaryBytes...)
	key = append(key, []byte("/")...)

	liquidTxHashBytes := []byte(liquidTxHash)
	key = append(key, liquidTxHashBytes...)
	key = append(key, []byte("/")...)

	return key
}
