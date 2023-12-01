package util

import (
	"encoding/binary"
	"encoding/hex"
)

func SerializeInt64(value int64) []byte {
	// Adding 1 because 0 will be interpreted as nil, which is an invalid key
	buf := make([]byte, 8)
	// Use binary.BigEndian to write the int64 into the byte slice
	binary.BigEndian.PutUint64(buf, uint64(value+1))
	return buf
}

func SerializeString(value string) []byte {
	byteArray := []byte(value)
	return byteArray
}

func SerializeHexString(value string) ([]byte, error) {
	return hex.DecodeString(value)
}
