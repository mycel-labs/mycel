package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/app/params"
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

func (topLevelDomain TopLevelDomain) IsEditable(sender string) (isEditable bool, err error) {
	if topLevelDomain.AccessControl[sender] == DomainRole_NO_ROLE {
		err = errorsmod.Wrapf(ErrTopLevelDomainNotEditable, "%s", sender)
	}
	isEditable = topLevelDomain.AccessControl[sender] == DomainRole_EDITOR || topLevelDomain.AccessControl[sender] == DomainRole_OWNER
	return isEditable, err
}

func (topLevelDomain *TopLevelDomain) ExtendExpirationDate(from time.Time, extensionPeriodInYear uint64) (expirationDate time.Time) {
	expirationDate = from.AddDate(0, 0, params.OneYearInDays*int(extensionPeriodInYear))
	topLevelDomain.ExpirationDate = expirationDate

	return expirationDate
}

func (topLevelDomain *TopLevelDomain) GetRole(address string) (role DomainRole) {
	role = topLevelDomain.AccessControl[address]
	return role
}
