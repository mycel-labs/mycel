package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
)

// BeforeEpochStart is the epoch start hook.
func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
}

// AfterEpochEnd is the epoch end hook.
func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	var burnt = sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0))

	config, found := k.GetEpochBurnConfig(ctx)
	if !found {
		panic("epoch burn config not found")
	}

	// Check epoch identifier
	if config.EpochIdentifier != epochIdentifier {
		return
	}

	// Get burn amount
	burnAmount, found := k.GetBurnAmount(ctx, uint64(config.CurrentBurnAmountIndex))
	// If not found, set default burn amount
	if !found {
		burnAmount = k.NewBurnAmount(ctx, config, uint64(config.CurrentBurnAmountIndex))
	}

	// Check if CurrentEpoch is smaller than TotalEpochs
	if burnAmount.CurrentEpoch > burnAmount.TotalEpochs {
		panic("current epoch is greater than total epochs")
	}

	// If CurrentEpoch is final epoch, update index and set next burnAmount
	if burnAmount.CurrentEpoch == burnAmount.TotalEpochs {
		config.CurrentBurnAmountIndex++
		k.SetEpochBurnConfig(ctx, config)

		burnAmount, found = k.GetBurnAmount(ctx, uint64(config.CurrentBurnAmountIndex))
		if !found {
			burnAmount = k.NewBurnAmount(ctx, config, uint64(config.CurrentBurnAmountIndex))
		}
	}

	// Calculate burnt amount
	if burnAmount.CumulativeBurntAmount.IsLT(burnAmount.TotalBurnAmount) {
		burnt = burnAmount.CalculateBurntAmount()
	}

	// TODO: Burn coins

	// Update burn amount
	burnAmount.CumulateBurntAmount(burnt)

	// Set burn amount
	k.SetBurnAmount(ctx, burnAmount)

	// Emit event
	EmitEpochBurnEvent(ctx, epochIdentifier, epochNumber, &burnAmount, burnt)

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
