package keeper

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/app/params"
	furnacetypes "github.com/mycel-domain/mycel/x/furnace/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// SetTopLevelDomain set a specific topLevelDomain in the store from its index
func (k Keeper) SetTopLevelDomain(ctx sdk.Context, topLevelDomain types.TopLevelDomain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopLevelDomainKeyPrefix))
	b := k.cdc.MustMarshal(&topLevelDomain)
	store.Set(types.TopLevelDomainKey(
		topLevelDomain.Name,
	), b)
}

// GetTopLevelDomain returns a topLevelDomain from its index
func (k Keeper) GetTopLevelDomain(
	ctx sdk.Context,
	name string,

) (val types.TopLevelDomain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopLevelDomainKeyPrefix))

	b := store.Get(types.TopLevelDomainKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTopLevelDomain removes a topLevelDomain from the store
func (k Keeper) RemoveTopLevelDomain(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopLevelDomainKeyPrefix))
	store.Delete(types.TopLevelDomainKey(
		name,
	))
}

// GetAllTopLevelDomain returns all topLevelDomain
func (k Keeper) GetAllTopLevelDomain(ctx sdk.Context) (list []types.TopLevelDomain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopLevelDomainKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TopLevelDomain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Get valid top level domain
func (k Keeper) GetValidTopLevelDomain(ctx sdk.Context, name string) (topLevelDomain types.TopLevelDomain, err error) {
	// Regex validation
	err = types.ValidateTopLevelDomainName(name)
	if err != nil {
		return topLevelDomain, err
	}

	// Get top level domain
	topLevelDomain, isFound := k.GetTopLevelDomain(ctx, name)
	if !isFound {
		return topLevelDomain, errorsmod.Wrapf(types.ErrDomainNotFound, "%s", name)
	}

	// Check if domain is not expired
	expirationDate := time.Unix(0, topLevelDomain.ExpirationDateInUnixNano)
	if ctx.BlockTime().After(expirationDate) && topLevelDomain.ExpirationDateInUnixNano != 0 {
		return topLevelDomain, errorsmod.Wrapf(types.ErrDomainExpired, "%s", name)
	}

	return topLevelDomain, nil
}

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

// Pay TLD registration fee
func (k Keeper) PayTopLevelDomainFee(ctx sdk.Context, payer sdk.AccAddress, domain types.TopLevelDomain, registrationPeriodInYear uint64) (registrationFee types.TopLevelDomainFee, err error) {
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
		return types.TopLevelDomainFee{}, err
	}

	// Get burn weight (=W)
	weight, err := k.GetBurnWeight(ctx)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}
	registrationFee.BurnWeight = weight.String()

	// Get price (=P)
	price, err := types.GetMycelPrice(denom)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Calc burn amount (=WX/P)
	amountToBurn := weight.Mul(math.LegacyNewDecFromBigInt(fee.BigInt())).Quo(math.LegacyNewDecFromBigInt(price.BigInt())).TruncateInt()
	amountToTreasury := fee.Sub(amountToBurn)

	registrationFee.FeeToBurn = sdk.NewCoin(denom, amountToBurn)
	registrationFee.FeeToTreasury = sdk.NewCoin(denom, amountToTreasury)

	// Send coins to treasury
	err = k.distributionKeeper.FundCommunityPool(ctx, sdk.NewCoins(registrationFee.FeeToTreasury), payer)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Send coins to furnace module
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, payer, furnacetypes.ModuleName, sdk.NewCoins(registrationFee.FeeToBurn))
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}
	// Store burn amount
	_, err = k.furnaceKeeper.AddRegistrationFeeToBurnAmounts(ctx, registrationPeriodInYear, registrationFee.FeeToBurn)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Set total registration fee
	if registrationFee.FeeToBurn.Denom == registrationFee.FeeToTreasury.Denom {
		registrationFee.TotalFee = sdk.NewCoins(registrationFee.FeeToBurn.Add(registrationFee.FeeToTreasury))
	} else {
		registrationFee.TotalFee = sdk.NewCoins(registrationFee.FeeToBurn, registrationFee.FeeToTreasury)
	}

	return registrationFee, nil
}

// Register top level domain
func (k Keeper) RegisterTopLevelDomain(ctx sdk.Context, domain types.TopLevelDomain, creator string, registrationPeriodInYear uint64) (fee types.TopLevelDomainFee, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Validate domain
	err = k.ValidateTopLevelDomain(ctx, domain)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Pay TLD registration fee
	fee, err = k.PayTopLevelDomainFee(ctx, creatorAddress, domain, registrationPeriodInYear)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Set domain
	k.SetTopLevelDomain(ctx, domain)

	// Emit event
	EmitRegisterTopLevelDomainEvent(ctx, domain, fee)

	return fee, nil
}

// Extend top level domain expirationDate
func (k Keeper) ExtendTopLevelDomainExpirationDate(ctx sdk.Context, creator string, domain *types.TopLevelDomain, extensionPeriodInYear uint64) (fee types.TopLevelDomainFee, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}
	// Check if the domain is editable
	_, err = domain.IsEditable(creator)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Pay TLD extend fee
	fee, err = k.PayTopLevelDomainFee(ctx, creatorAddress, *domain, extensionPeriodInYear)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Update domain store
	currentExpirationDate := time.Unix(0, domain.ExpirationDateInUnixNano)
	domain.ExtendExpirationDate(currentExpirationDate, extensionPeriodInYear)
	k.SetTopLevelDomain(ctx, *domain)

	// Emit event
	EmitExtendTopLevelDomainExpirationDateEvent(ctx, *domain, fee)

	return fee, err
}
