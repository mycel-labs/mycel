package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DomainKeyPrefix is the prefix to retrieve all Domain
	DomainKeyPrefix = "Domain/value/"
)

// DomainKey returns the store key to retrieve a Domain from the index fields
func DomainKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
