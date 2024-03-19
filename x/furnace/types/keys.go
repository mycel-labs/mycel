package types

const (
	// ModuleName defines the module name
	ModuleName = "furnace"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_furnace"
)

var ParamsKey = []byte("p_furnace")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	EpochBurnConfigKey = "EpochBurnConfig/value/"
)
