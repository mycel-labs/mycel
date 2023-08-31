package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	math "math"
)

func GetDefaultSubdomainConfig(baseFee int64) SubdomainConfig {
	defaultFee := sdk.NewCoin(MycelDenom, sdk.NewInt(baseFee))
	fees := GetFeeByNameLength(10, int(baseFee))

	return SubdomainConfig{
		MaxSubdomainRegistrations: math.MaxUint64,
		SubdomainRegistrationFees: &SubdomainRegistrationFees{
			FeeByLength: fees,
			DefaultFee:  &defaultFee,
		},
	}
}

func GetFeeByNameLength(base int, baseFee int) map[uint32]*Fee {
	fees := make(map[uint32]*Fee)
	for i := uint32(1); i < 5; i++ {
		amount := baseFee * int(math.Pow(float64(base), float64((5-i))))
		fee := sdk.NewCoin(MycelDenom, sdk.NewInt(int64(amount)))
		fees[i] = &Fee{
			IsRegistrable: true,
			Fee:           &fee,
		}
	}
	return fees
}

func (config *SubdomainConfig) GetRegistrationFee(name string, registrationPeriodInYear uint64) (amount *sdk.Coin, err error) {
	amount = config.SubdomainRegistrationFees.DefaultFee

	// Set amout if bylength found
	if config.SubdomainRegistrationFees.FeeByName[name] != nil {
		if config.SubdomainRegistrationFees.FeeByName[name].IsRegistrable {
			amount = config.SubdomainRegistrationFees.FeeByName[name].Fee
		} else {
			err = sdkerrors.Wrap(errors.New(name), ErrDomainNotRegistrable.Error())
		}
	}

	// Set amout if byname found
	if config.SubdomainRegistrationFees.FeeByLength[uint32(len(name))] != nil {
		if config.SubdomainRegistrationFees.FeeByLength[uint32(len(name))].IsRegistrable {
			amount = config.SubdomainRegistrationFees.FeeByLength[uint32(len(name))].Fee
		} else {
			err = sdkerrors.Wrap(errors.New(name), ErrDomainNotRegistrable.Error())
		}
	}

	amount.Amount = amount.Amount.Mul(sdk.NewInt(int64(registrationPeriodInYear)))

	return amount, err
}
