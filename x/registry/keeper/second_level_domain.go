package keeper

import (
	"github.com/mycel-domain/mycel/x/registry/types"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
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

// Get valid second level domain
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
	// Get second level domain
	secondLevelDomain, isFound := k.GetSecondLevelDomain(ctx, name, parent)
	if !isFound {
		return secondLevelDomain, errorsmod.Wrapf(types.ErrDomainNotFound, "%s.%s", name, parent)
	}

	// Check if domain is not expired
	expirationDate := time.Unix(0, secondLevelDomain.ExpirationDate)
	if ctx.BlockTime().After(expirationDate) && secondLevelDomain.ExpirationDate != 0 {
		return secondLevelDomain, errorsmod.Wrapf(types.ErrDomainExpired, "%s", name)
	}

	return secondLevelDomain, nil
}

// Get parent domain
func (k Keeper) GetParentDomain(ctx sdk.Context, domain types.SecondLevelDomain) (parentDomain types.TopLevelDomain, found bool) {
	// Get parent domain
	parent := domain.ParseParent()
	parentDomain, found = k.GetTopLevelDomain(ctx, parent)
	return parentDomain, found
}

// Get parent domain's subdomain config
func (k Keeper) GetParentsSubdomainConfig(ctx sdk.Context, domain types.SecondLevelDomain) types.SubdomainConfig {
	// Get parent domain
	parentDomain, found := k.GetParentDomain(ctx, domain)
	if !found || parentDomain.SubdomainConfig == nil {
		panic("parent domain or config not found")
	}
	return *parentDomain.SubdomainConfig
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
func (k Keeper) PaySLDRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.SecondLevelDomain, registrationPeriodInYear uint64) (fee *sdk.Coin, err error) {
	config := k.GetParentsSubdomainConfig(ctx, domain)

	fee, err = config.GetRegistrationFee(domain.Name, registrationPeriodInYear)
	if err != nil {
		return fee, err
	}

	// Send coins from payer to module account
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, payer, types.ModuleName, sdk.NewCoins(*fee))

	// Update store
	parent, found := k.GetTopLevelDomain(ctx, domain.Parent)
	if !found {
		panic("parent not found")
	}
	parent.TotalWithdrawalAmount = parent.TotalWithdrawalAmount.Add(*fee)
	k.SetTopLevelDomain(ctx, parent)

	return fee, err
}

// Append to owned domain
func (k Keeper) AppendToOwnedDomain(ctx sdk.Context, owner string, name string, parent string) {
	domainOwnership, found := k.GetDomainOwnership(ctx, owner)
	if found {
		domainOwnership.Domains = append(domainOwnership.Domains, &types.OwnedDomain{Name: name, Parent: parent})
		k.SetDomainOwnership(ctx, domainOwnership)
	} else {
		k.SetDomainOwnership(ctx, types.DomainOwnership{Owner: owner, Domains: []*types.OwnedDomain{{Name: name, Parent: parent}}})
	}
}

// Register second level domain
func (k Keeper) RegisterSecondLevelDomain(ctx sdk.Context, domain types.SecondLevelDomain, owner sdk.AccAddress, registrationPeriodIYear uint64) (err error) {
	// Validate domain
	err = k.ValidateSecondLevelDomain(ctx, domain)
	if err != nil {
		return err
	}

	// Get parent domain of second level domain
	parentDomain, found := k.GetParentDomain(ctx, domain)

	if !found {
		panic("parent not found")
	}

	// Check if parent domain has subdomain registration config
	if parentDomain.SubdomainConfig.MaxSubdomainRegistrations <= parentDomain.SubdomainCount {
		err = errorsmod.Wrapf(types.ErrMaxSubdomainCountReached, "%d", parentDomain.SubdomainCount)
		return err
	}

	// Increment parents subdomain SubdomainCount
	k.IncrementParentsSubdomainCount(ctx, domain)

	// Pay SLD registration fee
	fee, err := k.PaySLDRegstrationFee(ctx, owner, domain, registrationPeriodIYear)
	if err != nil {
		return err
	}

	// Append to owned domain
	k.AppendToOwnedDomain(ctx, owner.String(), domain.Name, domain.Parent)

	// Set domain
	k.SetSecondLevelDomain(ctx, domain)

	// Emit event
	EmitRegisterSecondLevelDomainEvent(ctx, domain, *fee)

	return err
}
