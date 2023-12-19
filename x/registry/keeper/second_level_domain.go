package keeper

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

// SetSecondLevelDomain set a specific second-level-domain in the store from its index
func (k Keeper) SetSecondLevelDomain(ctx sdk.Context, domain types.SecondLevelDomain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecondLevelDomainKeyPrefix))
	b := k.cdc.MustMarshal(&domain)
	store.Set(types.SecondLevelDomainKey(
		domain.Name,
		domain.Parent,
	), b)
}

// GetSecondLevelDomain returns a second-level-domain from its index
func (k Keeper) GetSecondLevelDomain(
	ctx sdk.Context,
	name string,
	parent string,
) (val types.SecondLevelDomain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecondLevelDomainKeyPrefix))

	b := store.Get(types.SecondLevelDomainKey(
		name,
		parent,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecondLevelDomain removes a second-level-domain from the store
func (k Keeper) RemoveSecondLevelDomain(
	ctx sdk.Context,
	name string,
	parent string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecondLevelDomainKeyPrefix))
	store.Delete(types.SecondLevelDomainKey(
		name,
		parent,
	))
}

// GetAllSecondLevelDomain returns all second-level-domain
func (k Keeper) GetAllSecondLevelDomain(ctx sdk.Context) (list []types.SecondLevelDomain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecondLevelDomainKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecondLevelDomain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Get is second-level-domain already taken
func (k Keeper) GetIsSecondLevelDomainAlreadyTaken(ctx sdk.Context, domain types.SecondLevelDomain) (isDomainAlreadyTaken bool) {
	_, isDomainAlreadyTaken = k.GetSecondLevelDomain(ctx, domain.Name, domain.Parent)
	return isDomainAlreadyTaken
}

// Get valid second-level-domain
func (k Keeper) GetValidSecondLevelDomain(ctx sdk.Context, name string, parent string) (secondLevelDomain types.SecondLevelDomain, err error) {
	// Regex validation
	err = types.ValidateSecondLevelDomainName(name)
	if err != nil {
		return secondLevelDomain, err
	}
	err = types.ValidateSecondLevelDomainParent(parent)
	if err != nil {
		return secondLevelDomain, err
	}

	// Get parent domain
	_, err = k.GetValidTopLevelDomain(ctx, parent)
	if err != nil {
		return types.SecondLevelDomain{}, err
	}

	// Get second-level-domain
	secondLevelDomain, isFound := k.GetSecondLevelDomain(ctx, name, parent)
	if !isFound {
		return types.SecondLevelDomain{}, errorsmod.Wrapf(types.ErrSecondLevelDomainNotFound, "%s.%s", name, parent)
	}

	// Check if second-level-domain is not expired
	if ctx.BlockTime().After(secondLevelDomain.ExpirationDate) && secondLevelDomain.ExpirationDate != (time.Time{}) {
		return types.SecondLevelDomain{}, errorsmod.Wrapf(types.ErrSecondLevelDomainExpired, "%s", name)
	}

	return secondLevelDomain, nil
}

// Get parent domain
func (k Keeper) GetSecondLevelDomainParent(ctx sdk.Context, domain types.SecondLevelDomain) (parentDomain types.TopLevelDomain, found bool) {
	// Get parent domain
	parent := domain.ParseParent()
	parentDomain, found = k.GetTopLevelDomain(ctx, parent)
	return parentDomain, found
}

// Get parent domain's subdomain config
func (k Keeper) GetSecondLevelDomainParentsSubdomainConfig(ctx sdk.Context, domain types.SecondLevelDomain) types.SubdomainConfig {
	// Get parent domain
	parentDomain, found := k.GetSecondLevelDomainParent(ctx, domain)
	if !found || parentDomain.SubdomainConfig == nil {
		panic("parent domain or config not found")
	}
	return *parentDomain.SubdomainConfig
}

// Get Role of the second-level domain
func (k Keeper) GetSecondLevelDomainRole(ctx sdk.Context, name, parent, address string) (role types.DomainRole, found bool) {
	sld, found := k.GetSecondLevelDomain(ctx, name, parent)
	if !found {
		return types.DomainRole_NO_ROLE, false
	}
	role = sld.GetRole(address)
	return role, true
}

// Increment parents subdomain count
func (k Keeper) IncrementParentsSubdomainCount(ctx sdk.Context, domain types.SecondLevelDomain) {
	// Increment parent's subdomain count
	parent := domain.ParseParent()
	parentDomain, found := k.GetTopLevelDomain(ctx, parent)
	if !found {
		panic("parent not found")
	}
	parentDomain.SubdomainCount++
	k.SetTopLevelDomain(ctx, parentDomain)
}

// Pay SLD registration fee
func (k Keeper) PaySecondLevelDomainRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.SecondLevelDomain, registrationPeriodInYear uint64) (fee sdk.Coin, err error) {
	config := k.GetSecondLevelDomainParentsSubdomainConfig(ctx, domain)

	fee, err = config.GetRegistrationFee(domain.Name, registrationPeriodInYear)
	if err != nil {
		return fee, err
	}

	// Send coins from payer to module account
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, payer, types.ModuleName, sdk.NewCoins(fee))
	if err != nil {
		return fee, err
	}

	// Update store
	parent, found := k.GetTopLevelDomain(ctx, domain.Parent)
	if !found {
		panic("parent not found")
	}
	parent.TotalWithdrawalAmount = parent.TotalWithdrawalAmount.Add(fee)
	k.SetTopLevelDomain(ctx, parent)

	return fee, err
}

// Validate second-level-domain is registrable
func (k Keeper) ValidateSecondLevelDomainIsRegistrable(ctx sdk.Context, secondLevelDomain types.SecondLevelDomain) error {
	// Validate second-level-domain
	err := secondLevelDomain.Validate()
	if err != nil {
		return err
	}
	// Check if second-level-domain is already taken
	isTaken := k.GetIsSecondLevelDomainAlreadyTaken(ctx, secondLevelDomain)
	if isTaken {
		return errorsmod.Wrapf(types.ErrSecondLevelDomainAlreadyTaken, "%s.%s", secondLevelDomain.Name, secondLevelDomain.Parent)
	}

	// Get parent domain of second-level-domain
	parentDomain, found := k.GetSecondLevelDomainParent(ctx, secondLevelDomain)
	if !found {
		return errorsmod.Wrapf(types.ErrSecondLevelDomainParentDoesNotExist, "%s", secondLevelDomain.Parent)
	}

	// Check if parent domain has subdomain registration config
	if parentDomain.SubdomainConfig.MaxSubdomainRegistrations <= parentDomain.SubdomainCount {
		return errorsmod.Wrapf(types.ErrTopLevelDomainMaxSubdomainCountReached, "%d", parentDomain.SubdomainCount)
	}

	return nil
}

// Register second level domain
func (k Keeper) RegisterSecondLevelDomain(ctx sdk.Context, secondLevelDomain types.SecondLevelDomain, owner sdk.AccAddress, registrationPeriodIYear uint64) (err error) {
	// Validate second-level-domain is registrable
	err = k.ValidateSecondLevelDomainIsRegistrable(ctx, secondLevelDomain)
	if err != nil {
		return err
	}

	// Increment parents subdomain SubdomainCount
	k.IncrementParentsSubdomainCount(ctx, secondLevelDomain)

	// Pay SLD registration fee
	fee, err := k.PaySecondLevelDomainRegstrationFee(ctx, owner, secondLevelDomain, registrationPeriodIYear)
	if err != nil {
		return err
	}

	// Append to owned domain
	k.AppendToOwnedDomain(ctx, owner.String(), secondLevelDomain.Name, secondLevelDomain.Parent)

	// Set domain
	k.SetSecondLevelDomain(ctx, secondLevelDomain)

	// Emit event
	EmitRegisterSecondLevelDomainEvent(ctx, secondLevelDomain, fee)

	return err
}
