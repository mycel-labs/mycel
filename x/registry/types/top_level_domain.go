package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (domain *TopLevelDomain) SetRegistrationFees(reqFeeByName []ReqRegistrationFeeByName, reqFeeByLength []ReqRegistrationFeeByLength, reqDefaultFee sdk.Coin) (err error) {
	for _, v := range reqFeeByName {
		domain.SubdomainConfig.SubdomainRegistrationFees.FeeByName[v.Name] = &Fee{
			IsRegistrable: v.IsRegistrable,
			Fee:           v.Fee,
		}
	}

	for _, v := range reqFeeByLength {
		domain.SubdomainConfig.SubdomainRegistrationFees.FeeByLength[v.Length] = &Fee{
			IsRegistrable: v.IsRegistrable,
			Fee:           v.Fee,
		}
	}

	domain.SubdomainConfig.SubdomainRegistrationFees.DefaultFee = &reqDefaultFee

	return err
}
