package types

import (
	"errors"
	"fmt"

	"cosmossdk.io/math"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/app/params"
)

func GetDefaultSubdomainConfig(baseFee int64) *SubdomainConfig {
	defaultFee := sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(baseFee))

	return &SubdomainConfig{
		MaxSubdomainRegistrations: 100_000,
		SubdomainRegistrationFees: &SubdomainRegistrationFees{
			DefaultFee:  &defaultFee,
		},
	}
}

// TODO: This function will cause consensus failure
func GetFeeByNameLength(base int64, baseFee int64, step int64) map[uint32]*Fee {
	fees := make(map[uint32]*Fee, step)
	baseDec, err := math.LegacyNewDecFromStr(fmt.Sprintf("%d", base))
	if err != nil {
		panic(err)
	}
	baseFeeDec, err := math.LegacyNewDecFromStr(fmt.Sprintf("%d", baseFee))
	if err != nil {
		panic(err)
	}
	for i := uint32(1); i <= uint32(step); i++ {
		exponent := uint64(step+1) - uint64(i)
		amount := baseFeeDec.Mul(baseDec.Power(exponent)).RoundInt()
		fee := sdk.NewCoin(params.DefaultBondDenom, amount)
		fees[i] = &Fee{
			IsRegistrable: true,
			Fee:           &fee,
		}
	}
	return fees
}

func (config SubdomainConfig) GetRegistrationFee(name string, registrationPeriodInYear uint64) (amount sdk.Coin, err error) {
	amount = *config.SubdomainRegistrationFees.DefaultFee

	// Set amount if bylength found
	if config.SubdomainRegistrationFees.FeeByName[name] != nil {
		if config.SubdomainRegistrationFees.FeeByName[name].IsRegistrable {
			amount = *config.SubdomainRegistrationFees.FeeByName[name].Fee
		} else {
			err = errorsmod.Wrap(errors.New(name), ErrSecondLevelDomainNotRegistrable.Error())
		}
	}

	// Set amount if byname found
	if config.SubdomainRegistrationFees.FeeByLength[uint32(len(name))] != nil {
		if config.SubdomainRegistrationFees.FeeByLength[uint32(len(name))].IsRegistrable {
			amount = *config.SubdomainRegistrationFees.FeeByLength[uint32(len(name))].Fee
		} else {
			err = errorsmod.Wrap(errors.New(name), ErrSecondLevelDomainNotRegistrable.Error())
		}
	}

	amount.Amount = amount.Amount.Mul(sdk.NewInt(int64(registrationPeriodInYear)))

	return amount, err
}
