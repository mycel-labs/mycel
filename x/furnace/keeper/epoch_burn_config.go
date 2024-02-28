package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/mycel-domain/mycel/x/furnace/types"
)

// SetEpochBurnConfig set epochBurnConfig in the store
func (k Keeper) SetEpochBurnConfig(goCtx context.Context, epochBurnConfig types.EpochBurnConfig) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))

	b := k.cdc.MustMarshal(&epochBurnConfig)
	store.Set([]byte{0}, b)
}

// GetEpochBurnConfig returns epochBurnConfig
func (k Keeper) GetEpochBurnConfig(goCtx context.Context) (val types.EpochBurnConfig, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEpochBurnConfig removes epochBurnConfig from the store
func (k Keeper) RemoveEpochBurnConfig(goCtx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))

	store.Delete([]byte{0})
}
