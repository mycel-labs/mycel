package keeper

import (
	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetParentDomain(ctx sdk.Context, domain types.SecondLevelDomain) (parentDomain types.TopLevelDomain, found bool) {
	// Get parent domain
	parent := domain.ParseParent()
	parentDomain, found = k.GetTopLevelDomain(ctx, parent)
	return parentDomain, found
}

func (k Keeper) GetParentsSubdomainConfig(ctx sdk.Context, domain types.SecondLevelDomain) types.SubdomainConfig {
	// Get parent domain
	parentDomain, found := k.GetParentDomain(ctx, domain)
	if !found || parentDomain.SubdomainConfig == nil {
		panic("parent domain or config not found")
	}
	return *parentDomain.SubdomainConfig
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
	parent.RegistrationFee = parent.RegistrationFee.Add(*fee)
	k.SetTopLevelDomain(ctx, parent)

	return fee, err
}

func (k Keeper) AppendToOwnedDomain(ctx sdk.Context, owner string, name string, parent string) {
	domainOwnership, found := k.GetDomainOwnership(ctx, owner)
	if found {
		domainOwnership.Domains = append(domainOwnership.Domains, &types.OwnedDomain{Name: name, Parent: parent})
		k.SetDomainOwnership(ctx, domainOwnership)
	} else {
		k.SetDomainOwnership(ctx, types.DomainOwnership{Owner: owner, Domains: []*types.OwnedDomain{{Name: name, Parent: parent}}})
	}
}

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

func (k Keeper) RegisterSecondLevelDomain(ctx sdk.Context, domain types.SecondLevelDomain, owner sdk.AccAddress, registrationPeriodIYear uint64) (err error) {
	// Validate domain
	err = k.ValidateSecondLevelDomain(ctx, domain)
	if err != nil {
		return err
	}

	// Pay registration fee
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
