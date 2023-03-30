package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EpochInfoKeyPrefix is the prefix to retrieve all EpochInfo
	EpochInfoKeyPrefix = "EpochInfo/value/"
)

// EpochInfoKey returns the store key to retrieve a EpochInfo from the index fields
func EpochInfoKey(
	identifier string,
) []byte {
	var key []byte

	identifierBytes := []byte(identifier)
	key = append(key, identifierBytes...)
	key = append(key, []byte("/")...)

	return key
}
