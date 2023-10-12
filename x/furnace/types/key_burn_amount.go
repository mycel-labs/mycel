package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BurnAmountKeyPrefix is the prefix to retrieve all BurnAmount
	BurnAmountKeyPrefix = "BurnAmount/value/"
)

// BurnAmountKey returns the store key to retrieve a BurnAmount from the index fields
func BurnAmountKey(
	index uint64,
) []byte {
	var key []byte

	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
