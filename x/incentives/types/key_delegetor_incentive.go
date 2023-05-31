package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DelegetorIncentiveKeyPrefix is the prefix to retrieve all DelegetorIncentive
	DelegetorIncentiveKeyPrefix = "DelegetorIncentive/value/"
)

// DelegetorIncentiveKey returns the store key to retrieve a DelegetorIncentive from the index fields
func DelegetorIncentiveKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
