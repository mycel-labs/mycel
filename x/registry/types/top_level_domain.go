package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
)

type TopLevelDomainRegistrationFee struct {
	TotalRegistrationFee      sdk.Coins
	BurnWeight                math.LegacyDec
	RegistrationFeeToBurn     sdk.Coin
	RegistrationFeeToTreasury sdk.Coin
}

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

func GetBeseFeeAmountInDenom(denom string, baseFeeInUsd int64) (amount math.Int, err error) {
	switch denom {
	// TODO: Get price from oracle
	case params.DefaultBondDenom:
		// 1USD = 10e6 umycel
		amount = sdk.NewInt(baseFeeInUsd)
	case "uusdc":
		// 1USD = 10e6 uusdc
		amount = sdk.NewInt(baseFeeInUsd)
	default:
		return amount, ErrInvalidDenom
	}
	return amount, nil
}

func (topLevelDommain TopLevelDomain) GetRegistrationFeeAmountInDenom(denom string, baseFeeInUsd int64, registrationPeriodInYear uint64) (amount math.Int, err error) {
	baseFeeAmount, err := GetBeseFeeAmountInDenom(denom, baseFeeInUsd)
	if err != nil {
		return amount, err
	}
	amount = sdk.NewInt(int64(registrationPeriodInYear) * int64(topLevelDommain.SubdomainConfig.MaxSubdomainRegistrations)).Mul(baseFeeAmount)
	return amount, nil
}
