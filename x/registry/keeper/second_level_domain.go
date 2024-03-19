package keeper

import (
	"context"
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

// SetSecondLevelDomain set a specific secondLevelDomain in the store from its index
func (k Keeper) SetSecondLevelDomain(ctx context.Context, secondLevelDomain types.SecondLevelDomain) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SecondLevelDomainKeyPrefix))
	b := k.cdc.MustMarshal(&secondLevelDomain)
	store.Set(types.SecondLevelDomainKey(
		secondLevelDomain.Name,
		secondLevelDomain.Parent,
	), b)
}

// GetSecondLevelDomain returns a secondLevelDomain from its index
func (k Keeper) GetSecondLevelDomain(
	ctx context.Context,
	name string,
	parent string,
) (val types.SecondLevelDomain, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SecondLevelDomainKeyPrefix))

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

// RemoveSecondLevelDomain removes a secondLevelDomain from the store
func (k Keeper) RemoveSecondLevelDomain(
	ctx context.Context,
	name string,
	parent string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SecondLevelDomainKeyPrefix))
	store.Delete(types.SecondLevelDomainKey(
		name,
		parent,
	))
}

// GetAllSecondLevelDomain returns all secondLevelDomain
func (k Keeper) GetAllSecondLevelDomain(ctx context.Context) (list []types.SecondLevelDomain) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SecondLevelDomainKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecondLevelDomain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Get is second-level-domain already taken
func (k Keeper) GetIsSecondLevelDomainAlreadyTaken(goCtx context.Context, domain types.SecondLevelDomain) (isDomainAlreadyTaken bool) {
	_, isDomainAlreadyTaken = k.GetSecondLevelDomain(goCtx, domain.Name, domain.Parent)
	return isDomainAlreadyTaken
}

// Get valid second-level-domain
func (k Keeper) GetValidSecondLevelDomain(goCtx context.Context, name string, parent string) (secondLevelDomain types.SecondLevelDomain, err error) {
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
	_, err = k.GetValidTopLevelDomain(goCtx, parent)
	if err != nil {
		return types.SecondLevelDomain{}, err
	}

	// Get second-level-domain
	secondLevelDomain, isFound := k.GetSecondLevelDomain(goCtx, name, parent)
	if !isFound {
		return types.SecondLevelDomain{}, errorsmod.Wrapf(types.ErrSecondLevelDomainNotFound, "%s.%s", name, parent)
	}

	// Check if second-level-domain is not expired
	ctx := sdk.UnwrapSDKContext(goCtx)
	if ctx.BlockTime().After(secondLevelDomain.ExpirationDate) && secondLevelDomain.ExpirationDate != (time.Time{}) {
		return types.SecondLevelDomain{}, errorsmod.Wrapf(types.ErrSecondLevelDomainExpired, "%s", name)
	}

	return secondLevelDomain, nil
}

// Get parent domain
func (k Keeper) GetSecondLevelDomainParent(goCtx context.Context, domain types.SecondLevelDomain) (parentDomain types.TopLevelDomain, found bool) {
	// Get parent domain
	parent := domain.ParseParent()
	parentDomain, found = k.GetTopLevelDomain(goCtx, parent)
	return parentDomain, found
}

// Get parent domain's subdomain config
func (k Keeper) GetSecondLevelDomainParentsSubdomainConfig(goCtx context.Context, domain types.SecondLevelDomain) types.SubdomainConfig {
	// Get parent domain
	parentDomain, found := k.GetSecondLevelDomainParent(goCtx, domain)
	if !found || parentDomain.SubdomainConfig == nil {
		panic("parent domain or config not found")
	}
	return *parentDomain.SubdomainConfig
}

// Get Role of the second-level domain
func (k Keeper) GetSecondLevelDomainRole(goCtx context.Context, name, parent, address string) (role types.DomainRole, found bool) {
	sld, found := k.GetSecondLevelDomain(goCtx, name, parent)
	if !found {
		return types.DomainRole_NO_ROLE, false
	}
	role = sld.GetRole(address)
	return role, true
}

// Increment parents subdomain count
func (k Keeper) IncrementParentsSubdomainCount(goCtx context.Context, domain types.SecondLevelDomain) {
	// Increment parent's subdomain count
	parent := domain.ParseParent()
	parentDomain, found := k.GetTopLevelDomain(goCtx, parent)
	if !found {
		panic("parent not found")
	}
	parentDomain.SubdomainCount++
	k.SetTopLevelDomain(goCtx, parentDomain)
}

// Pay SLD registration fee
func (k Keeper) PaySecondLevelDomainRegstrationFee(goCtx context.Context, payer sdk.AccAddress, domain types.SecondLevelDomain, registrationPeriodInYear uint64) (fee sdk.Coin, err error) {
	config := k.GetSecondLevelDomainParentsSubdomainConfig(goCtx, domain)

	fee, err = config.GetRegistrationFee(domain.Name, registrationPeriodInYear)
	if err != nil {
		return fee, err
	}

	// Send coins from payer to module account
	err = k.bankKeeper.SendCoinsFromAccountToModule(goCtx, payer, types.ModuleName, sdk.NewCoins(fee))
	if err != nil {
		return fee, err
	}

	// Update store
	parent, found := k.GetTopLevelDomain(goCtx, domain.Parent)
	if !found {
		panic("parent not found")
	}
	parent.TotalWithdrawalAmount = parent.TotalWithdrawalAmount.Add(fee)
	k.SetTopLevelDomain(goCtx, parent)

	return fee, err
}

// Validate second-level-domain is registrable
func (k Keeper) ValidateSecondLevelDomainIsRegistrable(goCtx context.Context, secondLevelDomain types.SecondLevelDomain, sldOwner sdk.AccAddress) error {
	// Validate second-level-domain
	err := secondLevelDomain.Validate()
	if err != nil {
		return err
	}
	// Check if second-level-domain is already taken
	isTaken := k.GetIsSecondLevelDomainAlreadyTaken(goCtx, secondLevelDomain)
	if isTaken {
		return errorsmod.Wrapf(types.ErrSecondLevelDomainAlreadyTaken, "%s.%s", secondLevelDomain.Name, secondLevelDomain.Parent)
	}

	// Get parent domain of second-level-domain
	parentDomain, found := k.GetSecondLevelDomainParent(goCtx, secondLevelDomain)
	if !found {
		return errorsmod.Wrapf(types.ErrSecondLevelDomainParentDoesNotExist, "%s", secondLevelDomain.Parent)
	}

	// Check if the registering domain is allowed or not
	isPrivate := parentDomain.SubdomainConfig.RegistrationPolicy == types.RegistrationPolicyType_PRIVATE
	isOwner := parentDomain.GetRole(sldOwner.String()) == types.DomainRole_OWNER

	if isPrivate && !isOwner {
		return errorsmod.Wrapf(types.ErrNotAllowedRegisterDomain, "%s", parentDomain.Name)
	}

	// Check if parent domain has subdomain registration config
	if parentDomain.SubdomainConfig.MaxSubdomainRegistrations <= parentDomain.SubdomainCount {
		return errorsmod.Wrapf(types.ErrTopLevelDomainMaxSubdomainCountReached, "%d", parentDomain.SubdomainCount)
	}

	return nil
}

// Register second level domain
func (k Keeper) RegisterSecondLevelDomain(goCtx context.Context, secondLevelDomain types.SecondLevelDomain, owner sdk.AccAddress, registrationPeriodIYear uint64) (err error) {
	// Validate second-level-domain is registrable
	err = k.ValidateSecondLevelDomainIsRegistrable(goCtx, secondLevelDomain, owner)
	if err != nil {
		return err
	}

	// Increment parents subdomain SubdomainCount
	k.IncrementParentsSubdomainCount(goCtx, secondLevelDomain)

	// Pay SLD registration fee
	fee, err := k.PaySecondLevelDomainRegstrationFee(goCtx, owner, secondLevelDomain, registrationPeriodIYear)
	if err != nil {
		return err
	}

	// Append to owned domain
	k.AppendToOwnedDomain(goCtx, owner.String(), secondLevelDomain.Name, secondLevelDomain.Parent)

	// Set domain
	k.SetSecondLevelDomain(goCtx, secondLevelDomain)

	// Emit event
	EmitRegisterSecondLevelDomainEvent(goCtx, secondLevelDomain, fee)

	return err
}
