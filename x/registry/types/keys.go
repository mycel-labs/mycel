package types

const (
	// ModuleName defines the module name
	ModuleName = "registry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_registry"
)

var ParamsKey = []byte("p_registry")

func KeyPrefix(p string) []byte {
	return []byte(p)
}
