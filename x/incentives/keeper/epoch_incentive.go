package keeper

import (
	epochstypes "mycel/x/epochs/types"
	"mycel/x/incentives/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEpochIncentive set a specific epochIncentive in the store from its index
func (k Keeper) SetEpochIncentive(ctx sdk.Context, epochIncentive types.EpochIncentive) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochIncentiveKeyPrefix))
	b := k.cdc.MustMarshal(&epochIncentive)
	store.Set(types.EpochIncentiveKey(
		epochIncentive.Epoch,
	), b)
}

// GetEpochIncentive returns a epochIncentive from its index
func (k Keeper) GetEpochIncentive(
	ctx sdk.Context,
	epoch int64,

) (val types.EpochIncentive, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochIncentiveKeyPrefix))

	b := store.Get(types.EpochIncentiveKey(
		epoch,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEpochIncentive removes a epochIncentive from the store
func (k Keeper) RemoveEpochIncentive(
	ctx sdk.Context,
	epoch int64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochIncentiveKeyPrefix))
	store.Delete(types.EpochIncentiveKey(
		epoch,
	))
}

// GetAllEpochIncentive returns all epochIncentive
func (k Keeper) GetAllEpochIncentive(ctx sdk.Context) (list []types.EpochIncentive) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochIncentiveKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EpochIncentive
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
func (k Keeper) SetEpochIncentivesOnRegistration(ctx sdk.Context, registrationPeriodInWeek uint, fee sdk.Coin) {
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
		incentive, found := k.GetEpochIncentive(ctx, nextEpoch+int64(i))
		amount := sdk.NewCoin(fee.Denom, amountPerEpoch)

		if !found {
			incentive = types.EpochIncentive{
				Epoch:  nextEpoch + int64(i),
				Amount: sdk.NewCoins(amount),
			}
		} else {
			incentive.Amount = incentive.Amount.Add(amount)
		}
		k.SetEpochIncentive(ctx, incentive)
	}
}

func (k Keeper) SendRegistrationFeeToIncentiveModule(ctx sdk.Context) {}
