package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
)

func GetMycelPrice(denom string) (price sdk.Int, err error) {
	if denom == params.DefaultBondDenom {
		price = sdk.NewInt(1 * params.MycelExponent)
	} else {
		// TODO: Get price from oracle
		return sdk.NewInt(0), ErrInvalidDenom
	}
	return price, nil
}

func (topLevelDommain TopLevelDomain) GetRegistrationFeeByDenom(denom string, registrationPeriodInYear uint64) (fee sdk.Coin, err error) {
	fee = sdk.NewCoin(denom, sdk.NewInt(int64(registrationPeriodInYear)*int64(topLevelDommain.SubdomainConfig.MaxSubdomainRegistrations)))
	return fee, nil
}
