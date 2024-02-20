package cmd

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/app/params"
)

// RegisterDenoms registers token denoms.
func RegisterDenoms() {
	err := sdk.RegisterDenom(params.HumanCoinUnit, math.LegacyOneDec())
	if err != nil {
		panic(err)
	}
	err = sdk.RegisterDenom(params.BaseCoinUnit, math.LegacyNewDecWithPrec(1, params.MycelExponent))
	if err != nil {
		panic(err)
	}
}

func initSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := params.Bech32PrefixAccAddr + "pub"
	validatorAddressPrefix := params.Bech32PrefixAccAddr + "valoper"
	validatorPubKeyPrefix := params.Bech32PrefixAccAddr + "valoperpub"
	consNodeAddressPrefix := params.Bech32PrefixAccAddr + "valcons"
	consNodePubKeyPrefix := params.Bech32PrefixAccAddr + "valconspub"

	// Set Denom
	RegisterDenoms()

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(params.Bech32PrefixAccAddr, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}
