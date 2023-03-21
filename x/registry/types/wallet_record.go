package types

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
