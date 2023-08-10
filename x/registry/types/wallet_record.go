package types

import (
	"errors"
	fmt "fmt"
	"regexp"

	"filippo.io/edwards25519"
	"github.com/btcsuite/btcutil/base58"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func WalletRecordFormats() map[string]string {
	return map[string]string{
		// Bitcoin
		"BITCOIN_MAINNET": "BITCOIN",
		"BITCOIN_TESTNET": "BITCOIN",
		// EVM
		"ETHEREUM_MAINNET": "ETHEREUM",
		"ETHEREUM_GOERLI":  "ETHEREUM",
		"ETHEREUM_SEPOLIA": "ETHEREUM",
		"POLYGON_MAINNET":  "ETHEREUM",
		"POLYGON_MUMBAI":   "ETHEREUM",
		"BNB_MAINNET":      "ETHEREUM",
		"BNB_TESTNET":      "ETHEREUM",
		"AVALANCHE_CCHAIN": "ETHEREUM",
		"AVALANCHE_FUJI":   "ETHEREUM",
		"GNOSIS_MAINNET":   "ETHEREUM",
		"GNOSIS_CHIADO":    "ETHEREUM",
		"OPTIMISM_MAINNET": "ETHEREUM",
		"OPTIMISM_GOERLI":  "ETHEREUM",
		"ARBITRUM_MAINNET": "ETHEREUM",
		"ARBITRUM_GOERLI":  "ETHEREUM",
		// 		"SHARDEUM_MAINNET":  "ETHEREUM",
		// 		"SHARDEUM_TESTNET":  "ETHEREUM",
		"SHARDEUM_BETANET": "ETHEREUM",
		// 		"ZETA_MAINNET":     "ETHEREUM",
		"ZETA_TESTNET": "ETHEREUM",

		// Move
		"APTOS_MAINNET": "MOVE",
		"APTOS_TESTNET": "MOVE",
		"SUI_MAINNET":   "MOVE",
		"SUI_TESTNET":   "MOVE",

		// Solana
		"SOLANA_MAINNET": "SOLANA",
		"SOLANA_TESTNET": "SOLANA",
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
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s %s", walletAddressFormat, address)), ErrInvalidWalletAddress.Error())
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
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s %s", walletAddressFormat, address)), ErrInvalidWalletAddress.Error())
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
