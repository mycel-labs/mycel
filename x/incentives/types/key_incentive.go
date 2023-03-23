package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// IncentiveKeyPrefix is the prefix to retrieve all Incentive
	IncentiveKeyPrefix = "Incentive/value/"
)

// IncentiveKey returns the store key to retrieve a Incentive from the index fields
func IncentiveKey(
	epoch uint64,
) []byte {
	var key []byte

	epochBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(epochBytes, epoch)
	key = append(key, epochBytes...)
	key = append(key, []byte("/")...)

	return key
}
