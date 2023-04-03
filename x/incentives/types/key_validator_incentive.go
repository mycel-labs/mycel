package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorIncentiveKeyPrefix is the prefix to retrieve all ValidatorIncentive
	ValidatorIncentiveKeyPrefix = "ValidatorIncentive/value/"
)

// ValidatorIncentiveKey returns the store key to retrieve a ValidatorIncentive from the index fields
func ValidatorIncentiveKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
