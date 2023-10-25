package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
)

const (
	TopLevelDomainBaseFeeInUSD = 0.05 * 10e6
)

func GetMycelPrice(denom string) (price math.Int, err error) {
	switch denom {
	case params.DefaultBondDenom:
		price = sdk.NewInt(1)
		// TODO: Get price from oracle
	default:
		return sdk.NewInt(0), ErrInvalidDenom
	}
	return price, nil
}

func GetBeseFeeAmountInDenom(denom string) (amount math.Int, err error) {
	switch denom {
	// TODO: Get price from oracle
	case params.DefaultBondDenom:
		// 1USD = 10e6 umycel
		amount = sdk.NewInt(TopLevelDomainBaseFeeInUSD)
	case "uusdc":
		// 1USD = 10e6 uusdc
		amount = sdk.NewInt(TopLevelDomainBaseFeeInUSD)
	default:
		return amount, ErrInvalidDenom
	}
	return amount, nil
}

func (topLevelDommain TopLevelDomain) GetRegistrationFeeAmountInDenom(denom string, registrationPeriodInYear uint64) (amount math.Int, err error) {
	baseFeeAmount, err := GetBeseFeeAmountInDenom(denom)
	if err != nil {
		return amount, err
	}
	amount = sdk.NewInt(int64(registrationPeriodInYear) * int64(topLevelDommain.SubdomainConfig.MaxSubdomainRegistrations)).Mul(baseFeeAmount)
	return amount, nil
}
