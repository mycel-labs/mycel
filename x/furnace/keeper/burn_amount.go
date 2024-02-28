package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/app/params"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

// SetBurnAmount set a specific burnAmount in the store from its index
func (k Keeper) SetBurnAmount(goCtx context.Context, burnAmount types.BurnAmount) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))
	b := k.cdc.MustMarshal(&burnAmount)
	store.Set(types.BurnAmountKey(
		burnAmount.Index,
	), b)
}

// GetBurnAmount returns a burnAmount from its index
func (k Keeper) GetBurnAmount(
	goCtx context.Context,
	index uint64,
) (val types.BurnAmount, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))

	b := store.Get(types.BurnAmountKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBurnAmount removes a burnAmount from the store
func (k Keeper) RemoveBurnAmount(
	goCtx context.Context,
	index uint64,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))

	store.Delete(types.BurnAmountKey(
		index,
	))
}

// GetAllBurnAmount returns all burnAmount
func (k Keeper) GetAllBurnAmount(goCtx context.Context) (list []types.BurnAmount) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BurnAmount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

// Create a next burnAmount
func (k Keeper) NewBurnAmount(goCtx context.Context, config types.EpochBurnConfig, index uint64) (burnAmount types.BurnAmount) {
	// Create burn amount
	burnAmount = types.NewBurnAmount(config, index)
	k.SetBurnAmount(goCtx, burnAmount)

	// Emit event
	EmitBurnAmountCreatedEvent(goCtx, &burnAmount)

	return burnAmount
}

// Add to total burn BurnAmount
func (k Keeper) AddToTotalBurnAmount(goCtx context.Context, index uint64, amount sdk.Coin) (newBurnAmount types.BurnAmount) {
	// Get burn amount
	burnAmount, found := k.GetBurnAmount(goCtx, index)
	if !found {
		panic("burn amount not found")
	}
	// Update burn amount
	burnAmount.TotalBurnAmount = burnAmount.TotalBurnAmount.Add(amount)
	k.SetBurnAmount(goCtx, burnAmount)
	return burnAmount
}

// Add registration fee to burnAmounts
func (k Keeper) AddRegistrationFeeToBurnAmounts(goCtx context.Context, registrationPeriodInYear uint64, amount sdk.Coin) (burnAmounts []types.BurnAmount, err error) {
	// Check registrationPeriodInYear
	if registrationPeriodInYear == 0 {
		return nil, errorsmod.Wrapf(types.ErrInvalidRegistrationPeriod, "%d", registrationPeriodInYear)
	}
	epochBurnConfig, found := k.GetEpochBurnConfig(goCtx)
	if !found {
		panic("epoch burn config not found")
	}

	remainDays := registrationPeriodInYear * params.OneYearInDays
	for i := epochBurnConfig.CurrentBurnAmountIndex + 1; remainDays > 0; i++ {
		burnAmount, found := k.GetBurnAmount(goCtx, i)
		// Create new burn amount if not found
		if !found {
			burnAmount = k.NewBurnAmount(goCtx, epochBurnConfig, i)
		}

		burnAmounts = append(burnAmounts, burnAmount)

		if remainDays >= burnAmount.TotalEpochs {
			remainDays -= burnAmount.TotalEpochs
		} else {
			remainDays = 0
		}
	}

	quotient := amount.Amount.QuoRaw(int64(len(burnAmounts)))
	reminder := amount.Amount.ModRaw(int64(len(burnAmounts)))

	// Set burnAmount
	for i, burnAmount := range burnAmounts {
		if !reminder.IsZero() && i == 0 {
			amount = sdk.NewCoin(amount.Denom, quotient.Add(reminder))
		} else {
			amount = sdk.NewCoin(amount.Denom, quotient)
		}
		burnAmounts[i].TotalBurnAmount = amount
		k.AddToTotalBurnAmount(goCtx, burnAmount.Index, amount)
	}
	return burnAmounts, err
}
