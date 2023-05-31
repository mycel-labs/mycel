package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/incentives/types"
)

// SetDelegetorIncentive set a specific delegetorIncentive in the store from its index
func (k Keeper) SetDelegetorIncentive(ctx sdk.Context, delegetorIncentive types.DelegetorIncentive) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegetorIncentiveKeyPrefix))
	b := k.cdc.MustMarshal(&delegetorIncentive)
	store.Set(types.DelegetorIncentiveKey(
		delegetorIncentive.Address,
	), b)
}

// GetDelegetorIncentive returns a delegetorIncentive from its index
func (k Keeper) GetDelegetorIncentive(
	ctx sdk.Context,
	address string,

) (val types.DelegetorIncentive, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegetorIncentiveKeyPrefix))

	b := store.Get(types.DelegetorIncentiveKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegetorIncentive removes a delegetorIncentive from the store
func (k Keeper) RemoveDelegetorIncentive(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegetorIncentiveKeyPrefix))
	store.Delete(types.DelegetorIncentiveKey(
		address,
	))
}

// GetAllDelegetorIncentive returns all delegetorIncentive
func (k Keeper) GetAllDelegetorIncentive(ctx sdk.Context) (list []types.DelegetorIncentive) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegetorIncentiveKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegetorIncentive
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
