package keeper

import (
	"mycel/x/incentives/types"

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
	epoch uint64,

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
	epoch uint64,

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
