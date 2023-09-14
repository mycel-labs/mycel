package keeper

import (
	// "errors"
	// "fmt"
	"github.com/mycel-domain/mycel/x/registry/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Pay TLD registration fee
func (k Keeper) PayTLDRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.TopLevelDomain, registrationPeriodInYear uint64) (err error) {
	// TODO: Pay fee
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
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeRegsterTopLevelDomain,
			sdk.NewAttribute(types.AttributeRegisterTopLevelDomainEventName, domain.Name),
			sdk.NewAttribute(types.AttributeRegisterTopLevelDomainEventExpirationDate, strconv.FormatInt(domain.ExpirationDate, 10)),
		),
	)

	return err
}
