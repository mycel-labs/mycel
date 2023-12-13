package types

import (
	"errors"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/app/params"
)

func GetDefaultSubdomainConfig(baseFee int64) SubdomainConfig {
	defaultFee := sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(baseFee))
	return SubdomainConfig{
		MaxSubdomainRegistrations: 100_000,
		SubdomainRegistrationFees: &SubdomainRegistrationFees{
			DefaultFee: &defaultFee,
		},
	}
}

func (config SubdomainConfig) GetRegistrationFee(name string, registrationPeriodInYear uint64) (amount sdk.Coin, err error) {
    amount = *config.SubdomainRegistrationFees.DefaultFee
    nameLength := uint32(len(name))

    // Check if a fee by name exists
    for _, feeByName := range config.SubdomainRegistrationFees.FeeByName {
        if feeByName.Name == name {
            if feeByName.IsRegistrable {
                amount = *feeByName.Fee
            } else {
                err = errorsmod.Wrap(errors.New(name), ErrSecondLevelDomainNotRegistrable.Error())
                return amount, err
            }
            break
        }
    }

    // Check if a fee by length exists if no specific name fee was found
    if err == nil {
        for _, feeByLength := range config.SubdomainRegistrationFees.FeeByLength {
            if feeByLength.Length == nameLength {
                if feeByLength.IsRegistrable {
                    amount = *feeByLength.Fee
                } else {
                    err = errorsmod.Wrap(errors.New(name), ErrSecondLevelDomainNotRegistrable.Error())
                }
                break
            }
        }
    }

    amount.Amount = amount.Amount.Mul(sdk.NewInt(int64(registrationPeriodInYear)))
    return amount, err
}
