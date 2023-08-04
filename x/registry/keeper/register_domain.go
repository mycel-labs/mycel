package keeper

import (
	incentivestypes "github.com/mycel-domain/mycel/x/incentives/types"
	"github.com/mycel-domain/mycel/x/registry/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Pay SLD registration fee
func (k Keeper) PaySLDRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.Domain, registrationPeriodInWeek uint) (err error) {
	fee := domain.GetRegistrationFee()
	k.incentivesKeeper.SetEpochIncentivesOnRegistration(ctx, registrationPeriodInWeek, fee)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, payer, incentivestypes.ModuleName, sdk.NewCoins(fee))
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

func (k Keeper) RegisterDomain(ctx sdk.Context, domain types.Domain, owner sdk.AccAddress, registrationPeriodInWeek uint) (err error) {
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
		// Increment parents subdomain SubdomainCount
		k.IncrementParentsSubdomainCount(ctx, domain)
		// Pay SLD registration fee
		err = k.PaySLDRegstrationFee(ctx, owner, domain, registrationPeriodInWeek)
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
