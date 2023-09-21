package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BurnAmountKeyPrefix is the prefix to retrieve all BurnAmount
	BurnAmountKeyPrefix = "BurnAmount/value/"
)

// BurnAmountKey returns the store key to retrieve a BurnAmount from the index fields
func BurnAmountKey(
	identifier uint64,
) []byte {
	var key []byte

	identifierBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(identifierBytes, identifier)
	key = append(key, identifierBytes...)
	key = append(key, []byte("/")...)

	return key
}
