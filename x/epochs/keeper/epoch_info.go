package keeper

import (
	"mycel/x/epochs/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEpochInfo set a specific epochInfo in the store from its index
func (k Keeper) SetEpochInfo(ctx sdk.Context, epochInfo types.EpochInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochInfoKeyPrefix))
	b := k.cdc.MustMarshal(&epochInfo)
	store.Set(types.EpochInfoKey(
		epochInfo.Identifier,
	), b)
}

// GetEpochInfo returns a epochInfo from its index
func (k Keeper) GetEpochInfo(
	ctx sdk.Context,
	identifier string,

) (val types.EpochInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochInfoKeyPrefix))

	b := store.Get(types.EpochInfoKey(
		identifier,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEpochInfo removes a epochInfo from the store
func (k Keeper) RemoveEpochInfo(
	ctx sdk.Context,
	identifier string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochInfoKeyPrefix))
	store.Delete(types.EpochInfoKey(
		identifier,
	))
}

// GetAllEpochInfo returns all epochInfo
func (k Keeper) GetAllEpochInfo(ctx sdk.Context) (list []types.EpochInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EpochInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Iterate though epochs
func (k Keeper) IterateEpochInfo(ctx sdk.Context, fn func(index int64, epochInfo types.EpochInfo) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochInfoKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		epoch := types.EpochInfo{}
		k.cdc.MustUnmarshal(iterator.Value(), &epoch)

		stop := fn(i, epoch)

		if stop {
			break
		}
		i++
	}
}
