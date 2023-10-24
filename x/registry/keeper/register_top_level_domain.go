package keeper

import (
	// "errors"
	// "fmt"
	"github.com/mycel-domain/mycel/x/registry/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
	furnacetypes "github.com/mycel-domain/mycel/x/furnace/types"
)

func (k Keeper) GetStatkingRatio(ctx sdk.Context) (ratio sdk.Int) {
	denom := params.DefaultBondDenom
	// Calc staking ratio
	totalSupply := k.bankKeeper.GetSupply(ctx, denom)
	moduleAccount := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	stakedAmount := k.bankKeeper.GetBalance(ctx, moduleAccount.GetAddress(), denom)
	stakingRatio := stakedAmount.Amount.Quo(totalSupply.Amount)
	return stakingRatio
}

// Pay TLD registration fee
func (k Keeper) PayTLDRegstrationFee(ctx sdk.Context, payer sdk.AccAddress, domain types.TopLevelDomain, registrationPeriodInYear uint64) (err error) {
	// TODO: Support other denoms
	denom := params.DefaultBondDenom

	// Calc fee
	fee, err := domain.GetRegistrationFeeByDenom(denom, registrationPeriodInYear)
	if err != nil {
		return err
	}

	// Send coins to furnace module
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, payer, furnacetypes.ModuleName, sdk.NewCoins(fee))
	if err != nil {
		return err
	}

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
