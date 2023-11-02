package keeper

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
	furnacetypes "github.com/mycel-domain/mycel/x/furnace/types"
)

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

// Pay TLD registration fee
func (k Keeper) PayTLDRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.TopLevelDomain, registrationPeriodInYear uint64) (registrationFee types.TopLevelDomainRegistrationFee, err error) {
	// TODO: Support other denoms
	denom := params.DefaultBondDenom

	// Get base fee
	baseFeeInUsd := k.GetParams(ctx).TopLevelDomainBaseFeeInUsd
	if baseFeeInUsd == 0 {
		panic("base fee is not set")
	}

	// Get Registration fee (=X)
	fee, err := domain.GetRegistrationFeeAmountInDenom(denom, baseFeeInUsd, registrationPeriodInYear)
	if err != nil {
		return types.TopLevelDomainRegistrationFee{}, err
	}

	// Get burn weight (=W)
	weight, err := k.GetBurnWeight(ctx)
	if err != nil {
		return types.TopLevelDomainRegistrationFee{}, err
	}
	registrationFee.BurnWeight = weight

	// Get price (=P)
	price, err := types.GetMycelPrice(denom)
	if err != nil {
		return types.TopLevelDomainRegistrationFee{}, err
	}

	// Calc burn amount (=WX/P)
	amountToBurn := weight.Mul(math.LegacyNewDecFromBigInt(fee.BigInt())).Quo(math.LegacyNewDecFromBigInt(price.BigInt())).TruncateInt()
	amountToTreasury := fee.Sub(amountToBurn)

	registrationFee.RegistrationFeeToBurn = sdk.NewCoin(denom, amountToBurn)
	registrationFee.RegistrationFeeToTreasury = sdk.NewCoin(denom, amountToTreasury)

	// Send coins to treasury
	err = k.distributionKeeper.FundCommunityPool(ctx, sdk.NewCoins(registrationFee.RegistrationFeeToTreasury), payer)
	if err != nil {
		return types.TopLevelDomainRegistrationFee{}, err
	}

	// Send coins to furnace module
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, payer, furnacetypes.ModuleName, sdk.NewCoins(registrationFee.RegistrationFeeToBurn))
	if err != nil {
		return types.TopLevelDomainRegistrationFee{}, err
	}
	// Store burn amount
	_, err = k.furnaceKeeper.AddRegistrationFeeToBurnAmounts(ctx, registrationPeriodInYear, registrationFee.RegistrationFeeToBurn)
	if err != nil {
		return types.TopLevelDomainRegistrationFee{}, err
	}

	// Set total registration fee
	if registrationFee.RegistrationFeeToBurn.Denom == registrationFee.RegistrationFeeToTreasury.Denom {
		registrationFee.TotalRegistrationFee = sdk.NewCoins(registrationFee.RegistrationFeeToBurn.Add(registrationFee.RegistrationFeeToTreasury))
	} else {
		registrationFee.TotalRegistrationFee = sdk.NewCoins(registrationFee.RegistrationFeeToBurn, registrationFee.RegistrationFeeToTreasury)
	}

	return registrationFee, nil
}

func (k Keeper) RegisterTopLevelDomain(ctx sdk.Context, domain types.TopLevelDomain, owner sdk.AccAddress, registrationPeriodIYear uint64) (err error) {
	// Validate domain
	err = k.ValidateTopLevelDomain(ctx, domain)
	if err != nil {
		return err
	}

	// Pay TLD registration fee
	topLevelDomainRegistrationFee, err := k.PayTLDRegstrationFee(ctx, owner, domain, registrationPeriodIYear)
	if err != nil {
		return err
	}

	// Set domain
	k.SetTopLevelDomain(ctx, domain)

	// Emit event
	EmitRegisterTopLevelDomainEvent(ctx, domain, topLevelDomainRegistrationFee)

	return err
}
