package keeper

import (
	"context"
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/app/params"
	furnacetypes "github.com/mycel-domain/mycel/x/furnace/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// SetTopLevelDomain set a specific topLevelDomain in the store from its index
func (k Keeper) SetTopLevelDomain(goCtx context.Context, topLevelDomain types.TopLevelDomain) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

	b := k.cdc.MustMarshal(&topLevelDomain)
	store.Set(types.TopLevelDomainKey(
		topLevelDomain.Name,
	), b)
}

// GetTopLevelDomain returns a topLevelDomain from its index
func (k Keeper) GetTopLevelDomain(
	goCtx context.Context,
	name string,
) (val types.TopLevelDomain, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

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
	goCtx context.Context,
	name string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

	store.Delete(types.TopLevelDomainKey(
		name,
	))
}

// GetAllTopLevelDomain returns all topLevelDomain
func (k Keeper) GetAllTopLevelDomain(goCtx context.Context) (list []types.TopLevelDomain) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TopLevelDomain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Get is top-level-domain already taken
func (k Keeper) GetIsTopLevelDomainAlreadyTaken(goCtx context.Context, domain types.TopLevelDomain) (isDomainAlreadyTaken bool) {
	_, isDomainAlreadyTaken = k.GetTopLevelDomain(goCtx, domain.Name)
	return isDomainAlreadyTaken
}

// Get valid-top-level domain
func (k Keeper) GetValidTopLevelDomain(goCtx context.Context, name string) (topLevelDomain types.TopLevelDomain, err error) {
	// Regex validation
	err = types.ValidateTopLevelDomainName(name)
	if err != nil {
		return topLevelDomain, err
	}

	// Get top level domain
	topLevelDomain, found := k.GetTopLevelDomain(goCtx, name)
	if !found {
		return topLevelDomain, errorsmod.Wrapf(types.ErrTopLevelDomainNotFound, "%s", name)
	}

	// Check if domain is not expired
	ctx := sdk.UnwrapSDKContext(goCtx)
	if ctx.BlockTime().After(topLevelDomain.ExpirationDate) && topLevelDomain.ExpirationDate != (time.Time{}) {
		return topLevelDomain, errorsmod.Wrapf(types.ErrTopLevelDomainExpired, "%s", name)
	}

	return topLevelDomain, nil
}

// Get Role of the domain
func (k Keeper) GetTopLevelDomainRole(goCtx context.Context, name, address string) (role types.DomainRole, found bool) {
	tld, found := k.GetTopLevelDomain(goCtx, name)
	if !found {
		return types.DomainRole_NO_ROLE, false
	}
	role = tld.GetRole(address)
	return role, true
}

// Pay top-level-domain registration fee
func (k Keeper) PayTopLevelDomainFee(goCtx context.Context, payer sdk.AccAddress, domain types.TopLevelDomain, registrationPeriodInYear uint64) (registrationFee types.TopLevelDomainFee, err error) {
	// Get registration fee
	registrationFee, err = k.GetTopLevelDomainFee(goCtx, domain, registrationPeriodInYear)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Send coins to treasury
	err = k.distributionKeeper.FundCommunityPool(goCtx, sdk.NewCoins(registrationFee.FeeToTreasury), payer)
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}

	// Send coins to furnace module
	err = k.bankKeeper.SendCoinsFromAccountToModule(goCtx, payer, furnacetypes.ModuleName, sdk.NewCoins(registrationFee.FeeToBurn))
	if err != nil {
		return types.TopLevelDomainFee{}, err
	}
	// Store burn amount
	_, err = k.furnaceKeeper.AddRegistrationFeeToBurnAmounts(goCtx, registrationPeriodInYear, registrationFee.FeeToBurn)
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

func (k Keeper) ValidateTopLevelDomainIsRegistrable(goCtx context.Context, topLevelDomain types.TopLevelDomain) error {
	// Validate top-level-domain
	err := topLevelDomain.Validate()
	if err != nil {
		return err
	}
	// Check if top-level-domain is already taken
	isTaken := k.GetIsTopLevelDomainAlreadyTaken(goCtx, topLevelDomain)
	if isTaken {
		return errorsmod.Wrapf(types.ErrTopLevelDomainAlreadyTaken, "%s", topLevelDomain.Name)
	}

	return nil
}

// Register top-level-domain
func (k Keeper) RegisterTopLevelDomain(goCtx context.Context, creator string, domainName string, registrationPeriodInYear uint64) (topLevelDomain types.TopLevelDomain, fee types.TopLevelDomainFee, err error) {
	// Create top-level-domain
	ctx := sdk.UnwrapSDKContext(goCtx)
	currentTime := ctx.BlockTime()
	expirationDate := currentTime.AddDate(0, 0, params.OneYearInDays*int(registrationPeriodInYear))
	accessControl := types.AccessControl{
		Address: creator,
		Role:    types.DomainRole_OWNER,
	}
	defaultRegistrationConfig := types.GetDefaultSubdomainConfig(303)
	topLevelDomain = types.TopLevelDomain{
		Name:                  domainName,
		ExpirationDate:        expirationDate,
		SubdomainConfig:       &defaultRegistrationConfig,
		AccessControl:         []*types.AccessControl{&accessControl},
		TotalWithdrawalAmount: sdk.NewCoins(),
	}

	// Validate top-level-domain is registrable
	err = k.ValidateTopLevelDomainIsRegistrable(goCtx, topLevelDomain)
	if err != nil {
		return types.TopLevelDomain{}, types.TopLevelDomainFee{}, err
	}

	// Pay TLD registration fee
	creatorAddress, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return types.TopLevelDomain{}, types.TopLevelDomainFee{}, err
	}
	fee, err = k.PayTopLevelDomainFee(goCtx, creatorAddress, topLevelDomain, registrationPeriodInYear)
	if err != nil {
		return types.TopLevelDomain{}, types.TopLevelDomainFee{}, err
	}

	// Set domain
	k.SetTopLevelDomain(goCtx, topLevelDomain)

	// Append to owned domain
	k.AppendToOwnedDomain(goCtx, creator, topLevelDomain.Name, "")

	// Emit event
	EmitRegisterTopLevelDomainEvent(goCtx, topLevelDomain, fee)

	return topLevelDomain, fee, nil
}

// Extend expiration date
func (k Keeper) ExtendTopLevelDomainExpirationDate(goCtx context.Context, creator string, domainName string, extensionPeriodInYear uint64) (topLevelDomain types.TopLevelDomain, fee types.TopLevelDomainFee, err error) {
	// Get domain
	topLevelDomain, found := k.GetTopLevelDomain(goCtx, domainName)
	if !found {
		return types.TopLevelDomain{}, types.TopLevelDomainFee{}, errorsmod.Wrapf(types.ErrTopLevelDomainNotFound, "%s", domainName)
	}

	// Check if the domain is editable
	_, err = topLevelDomain.IsEditable(creator)
	if err != nil {
		return types.TopLevelDomain{}, types.TopLevelDomainFee{}, err
	}

	creatorAddress, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return types.TopLevelDomain{}, types.TopLevelDomainFee{}, err
	}
	// Check if the domain is editable
	_, err = topLevelDomain.IsEditable(creator)
	if err != nil {
		return types.TopLevelDomain{}, types.TopLevelDomainFee{}, err
	}

	// Pay TLD extend fee
	fee, err = k.PayTopLevelDomainFee(goCtx, creatorAddress, topLevelDomain, extensionPeriodInYear)
	if err != nil {
		return types.TopLevelDomain{}, types.TopLevelDomainFee{}, err
	}

	// Update domain store
	topLevelDomain.ExtendExpirationDate(topLevelDomain.ExpirationDate, extensionPeriodInYear)
	k.SetTopLevelDomain(goCtx, topLevelDomain)

	// Emit event
	EmitExtendTopLevelDomainExpirationDateEvent(goCtx, topLevelDomain, fee)

	return topLevelDomain, fee, err
}

func (k Keeper) UpdateTopLevelDomainRegistrationPolicy(goCtx context.Context, creator string, domainName string, registrationPolicy string) (err error) {
	// Get domain
	topLevelDomain, found := k.GetTopLevelDomain(goCtx, domainName)
	if !found {
		return errorsmod.Wrapf(types.ErrTopLevelDomainNotFound, "%s", domainName)
	}

	// Check if the domain is editable
	_, err = topLevelDomain.IsEditable(creator)
	if err != nil {
		return err
	}

	// Validate registrationPolicy
	rp, err := topLevelDomain.ValidateTopLevelDomainRegistrationPolicy(registrationPolicy)
	if err != nil {
		return err
	}

	// Update domain store
	topLevelDomain.UpdateRegistrationPolicy(rp)
	k.SetTopLevelDomain(goCtx, topLevelDomain)

	// Emit event
	EmitUpdateTopLevelDomainRegistrationPolicyEvent(goCtx, topLevelDomain)

	return nil
}
