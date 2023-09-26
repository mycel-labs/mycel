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
	// Get the epoch burn config.
	config, found := k.GetEpochBurnConfig(ctx)
	if !found {
		panic("epoch burn config not found")
	}

	if config.EpochIdentifier == epochIdentifier {
		burnAmount, found := k.GetBurnAmount(ctx, config.CurrentBurnAmountIndex)
		if !found {
			panic("burn amount not found")
		}

		// Calc burn amount for this epoch.
		epochInfo, found := k.epochsKeeper.GetEpochInfo(ctx, epochIdentifier)
		if !found {
			panic("epoch info not found")
		}

		// TODO: Burn coins
		_ = epochInfo
		_ = burnAmount

		// Add the burn amount to burntAmount

		// Emit Events
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeEpochBurn,
				sdk.NewAttribute(types.AttributeKeyEpochIdentifier, epochIdentifier),
				sdk.NewAttribute(types.AttributeKeyEpochNumber, sdk.NewInt(epochNumber).String()),
				sdk.NewAttribute(types.AtributeKeyEpochBurnAmount, burnAmount.String()),
				sdk.NewAttribute(types.AtributeKeyEpochBurnTimestamp, ctx.BlockTime().String()),
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
