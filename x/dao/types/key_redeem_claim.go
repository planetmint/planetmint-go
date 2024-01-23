package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// RedeemClaimKeyPrefix is the prefix to retrieve all RedeemClaim
	RedeemClaimKeyPrefix                 = "RedeemClaim/value/"
	RedeemClaimBeneficiaryCountKeyPrefix = "RedeemClaim/beneficiary/count/"
	RedeemClaimLiquidTXKeyPrefix         = "RedeemClaim/liquidTX/value/"
)

// RedeemClaimKey returns the store key to retrieve a RedeemClaim from the index fields
func RedeemClaimKey(
	beneficiary string,
	id uint64,
) []byte {
	var key []byte

	beneficiaryBytes := []byte(beneficiary)
	key = append(key, beneficiaryBytes...)
	key = append(key, []byte("/")...)

	idBytes := SerializeUint64(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}

func SerializeUint64(value uint64) []byte {
	buf := make([]byte, 8)
	// Adding 1 because 0 will be interpreted as nil, which is an invalid key
	binary.BigEndian.PutUint64(buf, value+1)
	return buf
}

func DeserializeUint64(value []byte) uint64 {
	integer := binary.BigEndian.Uint64(value)
	// Subtract 1 because addition in serialization
	return integer - 1
}
