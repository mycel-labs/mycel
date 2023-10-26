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
	boundedRatio := k.mintKeeper.BondedRatio(ctx)

	// TODO: Get alpha from params
	stakingInflationRatio := k.GetParams(ctx).StakingInflationRatio
	alpha := math.LegacyMustNewDecFromStr(fmt.Sprintf("%f", stakingInflationRatio))

	w1 := alpha.Mul(boundedRatio)
	w2 := inflation.Mul(math.LegacyMustNewDecFromStr("1").Sub(alpha))
	weight = w1.Add(w2)
	return weight, nil
}

// Pay TLD registration fee
func (k Keeper) PayTLDRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.TopLevelDomain, registrationPeriodInYear uint64) (err error) {
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
		return err
	}

	// Get burn weight (=W)
	weight, err := k.GetBurnWeight(ctx)
	if err != nil {
		return err
	}

	// Get price (=P)
	price, err := types.GetMycelPrice(denom)
	if err != nil {
		return err
	}

	// Calc burn amount (=WX/P)
	amountToBurn := weight.Mul(math.LegacyNewDecFromBigInt(fee.BigInt())).Quo(math.LegacyNewDecFromBigInt(price.BigInt())).TruncateInt()
	amountToTreasury := fee.Sub(amountToBurn)

	coinToBurn := sdk.NewCoin(denom, amountToBurn)
	coinToTreasury := sdk.NewCoin(denom, amountToTreasury)

	// Send coins to treasury
	err = k.distributionKeeper.FundCommunityPool(ctx, sdk.NewCoins(coinToTreasury), payer)
	if err != nil {
		return err
	}

	// Send coins to furnace module
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, payer, furnacetypes.ModuleName, sdk.NewCoins(coinToBurn))
	if err != nil {
		return err
	}
	// Store burn amount
	_, err = k.furnaceKeeper.AddRegistrationFeeToBurnAmounts(ctx, registrationPeriodInYear, coinToBurn)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) RegisterTopLevelDomain(ctx sdk.Context, domain types.TopLevelDomain, owner sdk.AccAddress, registrationPeriodIYear uint64) (err error) {
	// Validate domain
	err = k.ValidateTopLevelDomain(ctx, domain)
	if err != nil {
		return err
	}

	// Pay TLD registration fee
	err = k.PayTLDRegstrationFee(ctx, owner, domain, registrationPeriodIYear)
	if err != nil {
		return err
	}

	// Set domain
	k.SetTopLevelDomain(ctx, domain)

	// Emit event
	EmitRegisterTopLevelDomainEvent(ctx, domain)


	return err
}
