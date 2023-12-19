package keeper

import (
	"fmt"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/app/params"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// Get burn weight
func (k Keeper) GetBurnWeight(ctx sdk.Context) (weight math.LegacyDec, err error) {
	inflation := k.mintKeeper.GetMinter(ctx).Inflation
	bondedRatio := k.mintKeeper.BondedRatio(ctx)

	// TODO: Get alpha from params
	stakingInflationRatio := k.GetParams(ctx).StakingInflationRatio
	alpha := math.LegacyMustNewDecFromStr(fmt.Sprintf("%f", stakingInflationRatio))

	w1 := alpha.Mul(bondedRatio)
	w2 := inflation.Mul(math.LegacyMustNewDecFromStr("1").Sub(alpha))
	weight = w1.Add(w2)
	return weight, nil
}

// Get top-level-domain fee
func (k Keeper) GetTopLevelDomainFee(ctx sdk.Context, topLevelDomain types.TopLevelDomain, registrationPeriodInYear uint64) (topLevelDomainFee types.TopLevelDomainFee, err error) {
	// TODO: Support other denoms
	denom := params.DefaultBondDenom

	// Get base fee
	baseFeeInUsd := k.GetParams(ctx).TopLevelDomainBaseFeeInUsd
	if baseFeeInUsd == 0 {
		panic("base fee is not set")
	}

	// Get Registration fee (=X)
	fee, err := topLevelDomain.GetRegistrationFeeAmountInDenom(denom, baseFeeInUsd, registrationPeriodInYear)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}
	topLevelDomainFee.TotalFee = sdk.NewCoins(sdk.NewCoin(denom, fee))

	// Get burn weight (=W)
	weight, err := k.GetBurnWeight(ctx)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}
	topLevelDomainFee.BurnWeight = weight.String()

	// Get price (=P)
	price, err := types.GetMycelPrice(denom)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Calc burn amount (=WX/P)
	amountToBurn := weight.Mul(math.LegacyNewDecFromBigInt(fee.BigInt())).Quo(math.LegacyNewDecFromBigInt(price.BigInt())).TruncateInt()
	amountToTreasury := fee.Sub(amountToBurn)
	topLevelDomainFee.FeeToBurn = sdk.NewCoin(denom, amountToBurn)
	topLevelDomainFee.FeeToTreasury = sdk.NewCoin(denom, amountToTreasury)

	return topLevelDomainFee, nil
}
