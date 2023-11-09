package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/furnace/types"
)

func EmitEpochBurnEvent(ctx sdk.Context, epochIdentifier string, epochNumber int64, burnAmount *types.BurnAmount, burnt sdk.Coin) {
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

func EmitBurnAmountCreatedEvent(ctx sdk.Context, burnAmount *types.BurnAmount) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBurnAmountCreated,
			sdk.NewAttribute(types.AttributeKeyBurnAmountIndex, sdk.NewInt(int64(burnAmount.Index)).String()),
		),
	)
}
