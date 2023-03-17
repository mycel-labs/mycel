package types

const (
	// ModuleName defines the module name
	ModuleName = "mycel"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mycel"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func WalletRecordFormats() map[string]string {
	return map[string]string{
		"ETHEREUM_MAINNET": "ETHEREUM",
		"ETHEREUM_GOERLI":  "ETHEREUM",
		"POLYGON_MAINNET":  "ETHEREUM",
		"POLYGON_MUMBAI":   "ETHEREUM",
	}
}

func WalletAddressRegex() map[string]string {
	return map[string]string{
		"ETHEREUM": "^0x[a-fA-F0-9]{40}$",
	}
}
