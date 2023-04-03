package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mycel/x/incentives/types"
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
	epoch int32,

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
	epoch int32,

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
