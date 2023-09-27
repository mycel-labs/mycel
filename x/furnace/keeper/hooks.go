package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

// BeforeEpochStart is the epoch start hook.
func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
}

// AfterEpochEnd is the epoch end hook.
func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	var burnt = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0))

	// Get the epoch burn config.
	config, found := k.GetEpochBurnConfig(ctx)
	if !found {
		panic("epoch burn config not found")
	}

	if config.EpochIdentifier == epochIdentifier {
		burnAmount, found := k.GetBurnAmount(ctx, uint64(config.CurrentBurnAmountIndex))
		if !found {
			panic("burn amount not found")
		}

		if burnAmount.CurrentEpoch <= burnAmount.TotalEpochs {
			// Check if the current epoch is the last epoch.
			if burnAmount.CurrentEpoch == burnAmount.TotalEpochs {
				config.CurrentBurnAmountIndex++
				k.SetEpochBurnConfig(ctx, config)
				burnAmount, found = k.GetBurnAmount(ctx, uint64(config.CurrentBurnAmountIndex))
				if !found {
					panic("burn amount not found")
				}
			}

			// Check if the current epoch is less than the total epochs.
			if burnAmount.CumulativeBurntAmount.IsLT(burnAmount.TotalBurnAmount) {
				// Check if the total burn amount is greater than the total epochs.
				if burnAmount.TotalBurnAmount.Amount.GTE(sdk.NewInt(int64(burnAmount.TotalEpochs))) {
					quotient := burnAmount.TotalBurnAmount.Amount.QuoRaw(int64(burnAmount.TotalEpochs))
					remander := burnAmount.TotalBurnAmount.Amount.ModRaw(int64(burnAmount.TotalEpochs))
					// Check if the remander is zero.
					if remander.IsZero() {
						// Set the burnt amount to the quotient.
						burnt = sdk.NewCoin(sdk.DefaultBondDenom, quotient)
					} else {
						// Check if the current epoch is the last epoch.
						if burnAmount.CurrentEpoch+1 == burnAmount.TotalEpochs {
							// Set the burnt amount to the quotient plus the remander.
							burnt = sdk.NewCoin(sdk.DefaultBondDenom, quotient.Add(remander))
						} else {
							// Set the burnt amount to the quotient.
							burnt = sdk.NewCoin(sdk.DefaultBondDenom, quotient)
						}
					}
				} else {
					if burnAmount.CurrentEpoch == 0 {
						burnt = sdk.NewCoin(sdk.DefaultBondDenom, burnAmount.TotalBurnAmount.Amount)
					}
				}
			}
		} else {
			panic("current epoch is greater than total epochs")
		}

		// TODO: Burn coins

		// Add the burn amount to burntAmount
		burnAmount.CumulativeBurntAmount = burnAmount.CumulativeBurntAmount.Add(burnt)
		burnAmount.CurrentEpoch++
		k.SetBurnAmount(ctx, burnAmount)

		// Emit Events
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeEpochBurn,
				sdk.NewAttribute(types.AttributeKeyEpochIdentifier, epochIdentifier),
				sdk.NewAttribute(types.AttributeKeyEpochNumber, sdk.NewInt(epochNumber).String()),
				sdk.NewAttribute(types.AttributeKeyBurnIndex, sdk.NewInt(int64(burnAmount.Index)).String()),
				sdk.NewAttribute(types.AttributeKeyBurnTotalEpochs, sdk.NewInt(int64(burnAmount.TotalEpochs)).String()),
				sdk.NewAttribute(types.AttributeKeyBurnCurrentEpoch, sdk.NewInt(int64(burnAmount.CurrentEpoch)).String()),
				sdk.NewAttribute(types.AttributeKeybBurnAmount, burnt.String()),
				sdk.NewAttribute(types.AttributeKeyBurnCumulativeAmount, burnAmount.CumulativeBurntAmount.String()),
				sdk.NewAttribute(types.AttributeKeyBurnTimestamp, ctx.BlockTime().String()),
			),
		)
	}
}

// ___________________________________________________________________________________________________

// Hooks is the wrapper struct for the incentives keeper.
type Hooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

// Hooks returns the hook wrapper struct.
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// BeforeEpochStart is the epoch start hook.
func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	h.k.BeforeEpochStart(ctx, epochIdentifier, epochNumber)
}

// AfterEpochEnd is the epoch end hook.
func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	h.k.AfterEpochEnd(ctx, epochIdentifier, epochNumber)
}
