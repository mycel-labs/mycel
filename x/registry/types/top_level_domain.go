package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"

	"github.com/mycel-domain/mycel/app/params"
)

func GetMycelPrice(denom string) (price math.Int, err error) {
	switch denom {
	case params.DefaultBondDenom:
		price = math.NewInt(1)
		// TODO: Get price from oracle
	default:
		return math.NewInt(0), ErrInvalidDenom
	}
	return price, nil
}

func GetBeseFeeAmountInDenom(denom string, baseFeeInUsd int64) (amount math.Int, err error) {
	switch denom {
	// TODO: Get price from oracle
	case params.DefaultBondDenom:
		// 1USD = 10e6 umycel
		amount = math.NewInt(baseFeeInUsd)
	case "uusdc":
		// 1USD = 10e6 uusdc
		amount = math.NewInt(baseFeeInUsd)
	default:
		return amount, ErrInvalidDenom
	}
	return amount, nil
}

func (topLevelDomain TopLevelDomain) GetRegistrationFeeAmountInDenom(denom string, baseFeeInUsd int64, registrationPeriodInYear uint64) (amount math.Int, err error) {
	baseFeeAmount, err := GetBeseFeeAmountInDenom(denom, baseFeeInUsd)
	if err != nil {
		return amount, err
	}
	amount = math.NewInt(int64(registrationPeriodInYear) * int64(topLevelDomain.SubdomainConfig.MaxSubdomainRegistrations)).Mul(baseFeeAmount)
	return amount, nil
}

func (topLevelDomain *TopLevelDomain) ExtendExpirationDate(from time.Time, extensionPeriodInYear uint64) (expirationDate time.Time) {
	expirationDate = from.AddDate(0, 0, params.OneYearInDays*int(extensionPeriodInYear))
	topLevelDomain.ExpirationDate = expirationDate

	return expirationDate
}

func (topLevelDomain *TopLevelDomain) UpdateRegistrationPolicy(rp RegistrationPolicyType) {
	topLevelDomain.SubdomainConfig.RegistrationPolicy = rp
}

func (topLevelDomain TopLevelDomain) IsEditable(sender string) (isEditable bool, err error) {
	role := topLevelDomain.GetRole(sender)
	if role == DomainRole_NO_ROLE {
		err = errorsmod.Wrapf(ErrTopLevelDomainNotEditable, "%s", sender)
	}
	isEditable = role == DomainRole_EDITOR || role == DomainRole_OWNER
	return isEditable, err
}

func (topLevelDomain *TopLevelDomain) GetRole(address string) (role DomainRole) {
	for _, accessControl := range topLevelDomain.AccessControl {
		if accessControl.Address == address {
			return accessControl.Role
		}
	}
	return DomainRole_NO_ROLE
}
