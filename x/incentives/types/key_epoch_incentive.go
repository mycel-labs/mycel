package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EpochIncentiveKeyPrefix is the prefix to retrieve all EpochIncentive
	EpochIncentiveKeyPrefix = "EpochIncentive/value/"
)

// EpochIncentiveKey returns the store key to retrieve a EpochIncentive from the index fields
func EpochIncentiveKey(
	epoch int64,
) []byte {
	var key []byte

	epochBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(epochBytes, uint64(epoch))
	key = append(key, epochBytes...)
	key = append(key, []byte("/")...)

	return key
}
