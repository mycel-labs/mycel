package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SecondLevelDomainKeyPrefix is the prefix to retrieve all Domain
	SecondLevelDomainKeyPrefix = "SecondLevelDomain/value/"
)

// SecondLevelDomainKey returns the store key to retrieve a Domain from the index fields
func SecondLevelDomainKey(
	name string,
	parent string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	parentBytes := []byte(parent)
	key = append(key, parentBytes...)
	key = append(key, []byte("/")...)

	return key
}
