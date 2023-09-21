package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

// SetBurnAmount set a specific burnAmount in the store from its index
func (k Keeper) SetBurnAmount(ctx sdk.Context, burnAmount types.BurnAmount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurnAmountKeyPrefix))
	b := k.cdc.MustMarshal(&burnAmount)
	store.Set(types.BurnAmountKey(
		burnAmount.Identifier,
	), b)
}

// GetBurnAmount returns a burnAmount from its index
func (k Keeper) GetBurnAmount(
	ctx sdk.Context,
	identifier uint64,

) (val types.BurnAmount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurnAmountKeyPrefix))

	b := store.Get(types.BurnAmountKey(
		identifier,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBurnAmount removes a burnAmount from the store
func (k Keeper) RemoveBurnAmount(
	ctx sdk.Context,
	identifier uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurnAmountKeyPrefix))
	store.Delete(types.BurnAmountKey(
		identifier,
	))
}

// GetAllBurnAmount returns all burnAmount
func (k Keeper) GetAllBurnAmount(ctx sdk.Context) (list []types.BurnAmount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurnAmountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BurnAmount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
