package keeper

import (
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/furnace/types"
)

// SetEpochBurnConfig set epochBurnConfig in the store
func (k Keeper) SetEpochBurnConfig(ctx sdk.Context, epochBurnConfig types.EpochBurnConfig) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))

	b := k.cdc.MustMarshal(&epochBurnConfig)
	store.Set([]byte{0}, b)
}

// GetEpochBurnConfig returns epochBurnConfig
func (k Keeper) GetEpochBurnConfig(ctx sdk.Context) (val types.EpochBurnConfig, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEpochBurnConfig removes epochBurnConfig from the store
func (k Keeper) RemoveEpochBurnConfig(ctx sdk.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))

	store.Delete([]byte{0})
}
