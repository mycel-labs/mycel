package keeper

import (
	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
	"github.com/mycel-domain/mycel/x/incentives/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetIncentive set a specific incentive in the store from its index
func (k Keeper) SetIncentive(ctx sdk.Context, incentive types.Incentive) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IncentiveKeyPrefix))
	b := k.cdc.MustMarshal(&incentive)
	store.Set(types.IncentiveKey(
		incentive.Epoch,
	), b)
}

// GetIncentive returns a incentive from its index
func (k Keeper) GetIncentive(
	ctx sdk.Context,
	epoch int64,

) (val types.Incentive, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IncentiveKeyPrefix))

	b := store.Get(types.IncentiveKey(
		epoch,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveIncentive removes a incentive from the store
func (k Keeper) RemoveIncentive(
	ctx sdk.Context,
	epoch int64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IncentiveKeyPrefix))
	store.Delete(types.IncentiveKey(
		epoch,
	))
}

// GetAllIncentive returns all incentive
func (k Keeper) GetAllIncentive(ctx sdk.Context) (list []types.Incentive) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IncentiveKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Incentive
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetIncentivesOnRegistration(ctx sdk.Context, registrationPeriodInWeek uint, fee sdk.Coin) {
	//Get current epoch
	epoch, found := k.epochsKeeper.GetEpochInfo(ctx, epochstypes.WeeklyEpochId)
	if !found {
		panic("current epoch not found")
	}
	nextEpoch := epoch.CurrentEpoch + 1

	amount := fee.Amount

	// Calculate amount to be distributed per epoch
	quotient := amount.QuoRaw(int64(registrationPeriodInWeek))
	remainder := amount.ModRaw(int64(registrationPeriodInWeek))

	amounts := make([]sdk.Int, registrationPeriodInWeek)

	for i := 0; i < int(registrationPeriodInWeek); i++ {
		amounts[i] = quotient
	}
	for i := 0; i < int(remainder.Int64()); i++ {
		amounts[i] = amounts[i].AddRaw(1)
	}

	// Set incentives
	for i, amountPerEpoch := range amounts {
		incentive, found := k.GetIncentive(ctx, nextEpoch+int64(i))
		amount := sdk.NewCoin(fee.Denom, amountPerEpoch)

		if !found {
			incentive = types.Incentive{
				Epoch:  nextEpoch + int64(i),
				Amount: sdk.NewCoins(amount),
			}
		} else {
			incentive.Amount = incentive.Amount.Add(amount)
		}
		k.SetIncentive(ctx, incentive)
	}
}

func (k Keeper) SendRegistrationFeeToIncentiveModule(ctx sdk.Context) {}
