package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TopLevelDomainKeyPrefix is the prefix to retrieve all TopLevelDomain
	TopLevelDomainKeyPrefix = "TopLevelDomain/value/"
)

// TopLevelDomainKey returns the store key to retrieve a TopLevelDomain from the index fields
func TopLevelDomainKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
