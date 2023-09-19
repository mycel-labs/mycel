package types

import (
	"errors"
	fmt "fmt"
	"regexp"

	errorsmod "cosmossdk.io/errors"
	"filippo.io/edwards25519"
	"github.com/btcsuite/btcutil/base58"
)

func WalletRecordFormats() map[string]string {
	return map[string]string{
		// Bitcoin
		"BITCOIN_MAINNET_MAINNET": "BITCOIN",
		"BITCOIN_TESTNET_TESTNET": "BITCOIN",

		// EVM
		"ETHEREUM_MAINNET_MAINNET": "ETHEREUM",
		"ETHEREUM_TESTNET_GOERLI":  "ETHEREUM",
		"ETHEREUM_TESTNET_SEPOLIA": "ETHEREUM",
		// Polygon
		"POLYGON_MAINNET_MAINNET": "ETHEREUM",
		"POLYGON_TESTNET_MUMBAI":  "ETHEREUM",
		// BNB
		"BNB_MAINNET_MAINNET": "ETHEREUM",
		"BNB_TESTNET_TESTNET": "ETHEREUM",
		// Avalanche
		"AVALANCHE_MAINNET_CCHAIN": "ETHEREUM",
		"AVALANCHE_TESTNET_FUJI":   "ETHEREUM",
		// Gnosis
		"GNOSIS_MAINNET_MAINNET": "ETHEREUM",
		"GNOSIS_TESTNET_CHIADO":  "ETHEREUM",
		// Optimism
		"OPTIMISM_MAINNET_MAINNET": "ETHEREUM",
		"OPTIMISM_TESTNET_GOERLI":  "ETHEREUM",
		// Arbitrum
		"ARBITRUM_MAINNET_MAINNET": "ETHEREUM",
		"ARBITRUM_TESTNET_GOERLI":  "ETHEREUM",
		// Shardeum
		"SHARDEUM_BETANET_SPHINX": "ETHEREUM",
		// ZetaChain
		"ZETA_TESTNET_ATHENS": "ETHEREUM",

		// Move
		"APTOS_MAINNET_MAINNET": "MOVE",
		"APTOS_TESTNET_TESTNET": "MOVE",
		"SUI_MAINNET_MAINNET":   "MOVE",
		"SUI_TESTNET_TESTNET":   "MOVE",

		// Solana
		"SOLANA_MAINNET_MAINNET": "SOLANA",
		"SOLANA_TESTNET_TESTNET": "SOLANA",
	}
}

func WalletAddressRegex() map[string]string {
	return map[string]string{
		"BITCOIN":  "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",
		"ETHEREUM": "^0x[a-fA-F0-9]{40}$",
		"MOVE":     "^0x[a-fA-F0-9]{64}$",
	}
}

func ValidateEd25519PublicKey(walletAddressFormat string, address string) (err error) {
	decodedBytes := base58.Decode(address)
	_, err = new(edwards25519.Point).SetBytes(decodedBytes)
	if err != nil {
		err = errorsmod.Wrapf(errors.New(fmt.Sprintf("%s %s", walletAddressFormat, address)), ErrInvalidWalletAddress.Error())
	}

	return err
}

func ValidateWalletAddressWithRegex(walletAddressFormat string, address string) (err error) {
	walletAddressRegex, isFound := WalletAddressRegex()[walletAddressFormat]
	if !isFound {
		panic(fmt.Sprintf("Wallet address format %s is not found in WalletAddressRegex", walletAddressFormat))
	}

	regex := regexp.MustCompile(walletAddressRegex)
	if !regex.MatchString(address) {
		err = errorsmod.Wrapf(errors.New(fmt.Sprintf("%s %s", walletAddressFormat, address)), ErrInvalidWalletAddress.Error())
	}
	return err

}

func ValidateWalletAddress(walletAddressFormat string, address string) (err error) {
	switch walletAddressFormat {
	case "BITCOIN", "ETHEREUM", "MOVE":
		err = ValidateWalletAddressWithRegex(walletAddressFormat, address)
	case "SOLANA":
		err = ValidateEd25519PublicKey(walletAddressFormat, address)
	default:
		panic(fmt.Sprintf("Wallet address format %s is not found in WalletAddressRegex", walletAddressFormat))
	}
	return err
}
