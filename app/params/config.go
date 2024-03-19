package params

import (
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	HumanCoinUnit = "mycel"
	BaseCoinUnit  = "umycel"
	MycelExponent = 6

	DefaultBondDenom = BaseCoinUnit

	Bech32PrefixAccAddr = "mycel"

	OneYearInDays = 365
)

// RegisterDenoms registers token denoms.
func RegisterDenoms() {
	err := sdk.RegisterDenom(HumanCoinUnit, math.LegacyOneDec())
	if err != nil {
		panic(err)
	}
	err = sdk.RegisterDenom(BaseCoinUnit, math.LegacyNewDecWithPrec(1, MycelExponent))
	if err != nil {
		panic(err)
	}
}
