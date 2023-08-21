package keeper

import (
	"errors"
	"fmt"
	"github.com/mycel-domain/mycel/x/registry/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetParentDomain(ctx sdk.Context, domain types.Domain) (parentDomain types.Domain, found bool) {
	// Check if domain is not TLD
	level := domain.GetDomainLevel()
	if level == 1 {
		found = false
	} else {
		// Get parent domain
		parentsName, parentsParent := domain.ParseParent()
		parentDomain, found = k.GetDomain(ctx, parentsName, parentsParent)
	}
	return parentDomain, found
}

func (k Keeper) GetParentsSubdomainRegistraionConfig(ctx sdk.Context, domain types.Domain) types.SubdomainRegistrationConfig {
	// Get parent domain
	parentDomain, found := k.GetParentDomain(ctx, domain)
	if !found || parentDomain.SubdomainRegistrationConfig == nil {
		panic("parent domain or config not found")
	}
	return *parentDomain.SubdomainRegistrationConfig
}

// Pay SLD registration fee
func (k Keeper) PaySLDRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.Domain, registrationPeriodInYear uint64) (err error) {
	config := k.GetParentsSubdomainRegistraionConfig(ctx, domain)

	fee, err := config.GetRegistrationFee(domain.Name, registrationPeriodInYear)
	if err != nil {
		return err
	}

	// TODO: Pay fee
	fee = fee
	return err
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

func (k Keeper) IncrementParentsSubdomainCount(ctx sdk.Context, domain types.Domain) {
	// Check if domain is not TLD
	level := domain.GetDomainLevel()
	if level == 1 {
		panic("domain is TLD")
	}

	// Increment parent's subdomain count
	parentsName, parentsParent := domain.ParseParent()
	parentDomain, found := k.GetDomain(ctx, parentsName, parentsParent)
	if !found {
		panic("parent not found")
	}
	parentDomain.SubdomainCount++
	k.SetDomain(ctx, parentDomain)
}

func (k Keeper) RegisterDomain(ctx sdk.Context, domain types.Domain, owner sdk.AccAddress, registrationPeriodIYear uint64) (err error) {
	// Validate domain
	err = k.ValidateDomain(ctx, domain)
	if err != nil {
		return err
	}

	// Pay registration fee
	domainLevel := domain.GetDomainLevel()
	switch domainLevel {
	case 1: // TLD
		// TODO: Register TLD
		return
	case 2: // SLD
		parentDomain, found := k.GetParentDomain(ctx, domain)

		if !found {
			panic("parent not found")
		}

		// Check if parent domain has subdomain registration config
		if parentDomain.SubdomainRegistrationConfig.MaxSubdomainRegistrations <= parentDomain.SubdomainCount {
			err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%d", parentDomain.SubdomainCount)), types.ErrMaxSubdomainCountReached.Error())
			return err
		}

		// Set subdomain registration config
		domain.SubdomainRegistrationConfig = &types.SubdomainRegistrationConfig{
			MaxSubdomainRegistrations: 100,
		}

		// Increment parents subdomain SubdomainCount
		k.IncrementParentsSubdomainCount(ctx, domain)
		// Pay SLD registration fee
		err = k.PaySLDRegstrationFee(ctx, owner, domain, registrationPeriodIYear)
		if err != nil {
			return err
		}
	default: // Subdomain
		// Increment parents subdomain SubdomainCount
		k.IncrementParentsSubdomainCount(ctx, domain)
	}

	// Append to owned domain
	k.AppendToOwnedDomain(ctx, owner.String(), domain.Name, domain.Parent)

	// Set domain
	k.SetDomain(ctx, domain)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeRegsterDomain,
			sdk.NewAttribute(types.AttributeRegisterDomainEventName, domain.Name),
			sdk.NewAttribute(types.AttributeRegisterDomainEventParent, domain.Parent),
			sdk.NewAttribute(types.AttributeRegisterDomainEventExpirationDate, strconv.FormatInt(domain.ExpirationDate, 10)),
			sdk.NewAttribute(types.AttributeRegisterDomainLevel, strconv.Itoa(domainLevel)),
		),
	)

	return err
}
