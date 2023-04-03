package keeper

import (
	"errors"
	"fmt"
	incentivestypes "mycel/x/incentives/types"
	"mycel/x/registry/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Get is domain already taken
func (k Keeper) GetIsDomainAlreadyTaken(ctx sdk.Context, domain types.Domain) (isDomainAlreadyTaken bool) {
	_, isDomainAlreadyTaken = k.GetDomain(ctx, domain.Name, domain.Parent)
	return isDomainAlreadyTaken
}

// Get is parent domain exist
func (k Keeper) GetIsParentDomainExist(ctx sdk.Context, domain types.Domain) (isParentDomainExist bool) {
	name, parent := domain.ParseParent()
	_, isParentDomainExist = k.GetDomain(ctx, name, parent)
	return isParentDomainExist
}

// Validate TLD registration
func (k Keeper) ValidateRegisterTLD(ctx sdk.Context, domain types.Domain) (err error) {
	if domain.Parent != "" {
		err = sdkerrors.Wrapf(errors.New(domain.Parent),
			types.ErrParentDomainMustBeEmpty.Error())
	}
	// TODO: Is Staked enough to register TLD
	return err
}

// Validate SLD registration
func (k Keeper) ValidateRegsiterSLD(ctx sdk.Context, domain types.Domain) (err error) {
	isParentDomainExist := k.GetIsParentDomainExist(ctx, domain)
	if !isParentDomainExist {
		err = sdkerrors.Wrapf(errors.New(domain.Parent),
			types.ErrParentDomainDoesNotExist.Error())
	}
	return err
}

// Validate subdomain GetRegistrationFee
func (k Keeper) ValidateRegsiterSubdomain(ctx sdk.Context, domain types.Domain) (err error) {
	isParentDomainExist := k.GetIsParentDomainExist(ctx, domain)
	if !isParentDomainExist {
		err = sdkerrors.Wrapf(errors.New(domain.Parent),
			types.ErrParentDomainDoesNotExist.Error())
	}
	return err
}

// Validate domain
func (k Keeper) ValidateDomain(ctx sdk.Context, domain types.Domain) (err error) {
	// Type check
	err = domain.Validate()
	if err != nil {
		return err
	}
	// Check if domain is already taken
	isDomainAlreadyTaken := k.GetIsDomainAlreadyTaken(ctx, domain)
	if isDomainAlreadyTaken {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s.%s", domain.Name, domain.Parent)),
			types.ErrDomainIsAlreadyTaken.Error())
		return err
	}

	domainLevel := domain.GetDomainLevel()
	switch domainLevel {
	case 1: // TLD
		// Validate TLD
		err = k.ValidateRegisterTLD(ctx, domain)
		if err != nil {
			return err
		}
	case 2: // TLD
		// Validate SLD
		err = k.ValidateRegsiterSLD(ctx, domain)
		if err != nil {
			return err
		}
	default: // Subdomain
		// Validate Subdomain
		err = k.ValidateRegsiterSubdomain(ctx, domain)
		if err != nil {
			return err
		}
	}

	return err
}

// Pay SLD registration fee
func (k Keeper) PayTLDRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.Domain, registrationPeriodInWeek uint) (err error) {
	// TODO: Pay TLD registration fee
	return err
}

// Pay SLD registration fee
func (k Keeper) PaySLDRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.Domain, registrationPeriodInWeek uint) (err error) {
	fee := domain.GetRegistrationFee()
	k.incentivesKeeper.SetEpochIncentivesOnRegistration(ctx, registrationPeriodInWeek, fee)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, payer, incentivestypes.ModuleName, sdk.NewCoins(fee))
	return err
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
		// Pay TLD registration fee
		err = k.PayTLDRegstrationFee(ctx, owner, domain, registrationPeriodInWeek)
		if err != nil {
			return err
		}
	case 2: // TLD
		// Pay SLD registration fee
		err = k.PaySLDRegstrationFee(ctx, owner, domain, registrationPeriodInWeek)
		if err != nil {
			return err
		}
	default: // Subdomain
	}

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
