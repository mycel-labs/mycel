package types

const (
	// ModuleName defines the module name
	ModuleName = "epochs"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_epochs"
)

var ParamsKey = []byte("p_epochs")

func KeyPrefix(p string) []byte {
	return []byte(p)
}
