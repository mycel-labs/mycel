package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

// SetEpochBurnConfig set epochBurnConfig in the store
func (k Keeper) SetEpochBurnConfig(ctx context.Context, epochBurnConfig types.EpochBurnConfig) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EpochBurnConfigKey))
	b := k.cdc.MustMarshal(&epochBurnConfig)
	store.Set([]byte{0}, b)
}

// GetEpochBurnConfig returns epochBurnConfig
func (k Keeper) GetEpochBurnConfig(ctx context.Context) (val types.EpochBurnConfig, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EpochBurnConfigKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEpochBurnConfig removes epochBurnConfig from the store
func (k Keeper) RemoveEpochBurnConfig(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EpochBurnConfigKey))
	store.Delete([]byte{0})
}
