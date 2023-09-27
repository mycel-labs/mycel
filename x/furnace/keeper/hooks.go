package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

func emitEpochBurnEvent(ctx sdk.Context, epochIdentifier string, epochNumber int64, burnAmount *types.BurnAmount, burnt sdk.Coin) {
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

func calculateBurntAmount(burnAmount *types.BurnAmount) sdk.Coin {
	if burnAmount.TotalBurnAmount.Amount.GTE(sdk.NewInt(int64(burnAmount.TotalEpochs))) {
		quotient := burnAmount.TotalBurnAmount.Amount.QuoRaw(int64(burnAmount.TotalEpochs))
		remander := burnAmount.TotalBurnAmount.Amount.ModRaw(int64(burnAmount.TotalEpochs))
		if remander.IsZero() || burnAmount.CurrentEpoch+1 != burnAmount.TotalEpochs {
			return sdk.NewCoin(sdk.DefaultBondDenom, quotient)
		}
		return sdk.NewCoin(sdk.DefaultBondDenom, quotient.Add(remander))
	} else if burnAmount.CurrentEpoch == 0 {
		return sdk.NewCoin(sdk.DefaultBondDenom, burnAmount.TotalBurnAmount.Amount)
	}
	return sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0))
}

// BeforeEpochStart is the epoch start hook.
func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
}

// AfterEpochEnd is the epoch end hook.
func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	var burnt = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0))

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
		panic("burn amount not found")
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
			panic("burn amount not found")
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
	emitEpochBurnEvent(ctx, epochIdentifier, epochNumber, &burnAmount, burnt)

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
