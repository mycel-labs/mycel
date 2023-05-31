package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DomainOwnershipKeyPrefix is the prefix to retrieve all DomainOwnership
	DomainOwnershipKeyPrefix = "DomainOwnership/value/"
)

// DomainOwnershipKey returns the store key to retrieve a DomainOwnership from the index fields
func DomainOwnershipKey(
	owner string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
