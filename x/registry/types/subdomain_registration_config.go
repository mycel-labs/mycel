package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	math "math"
)

func GetDefaultSubdomainRegistrationConfig(baseFee int64) SubdomainRegistrationConfig {
	defaultFee := sdk.NewCoin(MycelDenom, sdk.NewInt(baseFee))
	fees := GetFeeByNameLength(10, int(baseFee))

	return SubdomainRegistrationConfig{
		MaxSubdomainRegistrations: 100,
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
