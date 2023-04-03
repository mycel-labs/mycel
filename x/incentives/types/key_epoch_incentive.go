package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EpochIncentiveKeyPrefix is the prefix to retrieve all EpochIncentive
	EpochIncentiveKeyPrefix = "EpochIncentive/value/"
)

// EpochIncentiveKey returns the store key to retrieve a EpochIncentive from the index fields
func EpochIncentiveKey(
	epoch int32,
) []byte {
	var key []byte

	epochBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(epochBytes, uint32(epoch))
	key = append(key, epochBytes...)
	key = append(key, []byte("/")...)

	return key
}
