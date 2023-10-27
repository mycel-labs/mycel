package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

func calculateBurntAmount(burnAmount *types.BurnAmount) sdk.Coin {
	if burnAmount.TotalBurnAmount.Amount.GTE(sdk.NewInt(int64(burnAmount.TotalEpochs))) {
		quotient := burnAmount.TotalBurnAmount.Amount.QuoRaw(int64(burnAmount.TotalEpochs))
		remainder := burnAmount.TotalBurnAmount.Amount.ModRaw(int64(burnAmount.TotalEpochs))
		if remainder.IsZero() || burnAmount.CurrentEpoch+1 != burnAmount.TotalEpochs {
			return sdk.NewCoin(params.DefaultBondDenom, quotient)
		}
		return sdk.NewCoin(params.BaseCoinUnit, quotient.Add(remainder))
	} else if burnAmount.CurrentEpoch == 0 {
		return sdk.NewCoin(params.DefaultBondDenom, burnAmount.TotalBurnAmount.Amount)
	}
	return sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0))
}

func createNextBurnAmount(ctx sdk.Context, k Keeper, config types.EpochBurnConfig) (burnAmount types.BurnAmount) {
	// Create burn amount
	burnAmount = types.BurnAmount{
		Index:                 uint64(config.CurrentBurnAmountIndex),
		TotalEpochs:           config.DefaultTotalEpochs,
		CurrentEpoch:          0,
		TotalBurnAmount:       sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0)),
		CumulativeBurntAmount: sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0)),
	}
	k.SetBurnAmount(ctx, burnAmount)

	// Emit event
	EmitBurnAmountCreatedEvent(ctx, &burnAmount)

	return burnAmount
}

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
	if !found {
		burnAmount = createNextBurnAmount(ctx, k, config)
	}

	// Check if CurrentEpoch is smaller than TotalEpochs
	if burnAmount.CurrentEpoch > burnAmount.TotalEpochs {
		panic("current epoch is greater than total epochs")
	}

	// Check if CurrentEpoch is final epoch
	if burnAmount.CurrentEpoch == burnAmount.TotalEpochs {
		config.CurrentBurnAmountIndex++
		k.SetEpochBurnConfig(ctx, config)

		burnAmount, found = k.GetBurnAmount(ctx, uint64(config.CurrentBurnAmountIndex))
		if !found {
			burnAmount = createNextBurnAmount(ctx, k, config)
		}
	}

	// Calculate burnt amount
	if burnAmount.CumulativeBurntAmount.IsLT(burnAmount.TotalBurnAmount) {
		burnt = calculateBurntAmount(&burnAmount)
	}

	// TODO: Burn coins

	// Update burn amount
	burnAmount.CumulativeBurntAmount = burnAmount.CumulativeBurntAmount.Add(burnt)
	burnAmount.CurrentEpoch++
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
