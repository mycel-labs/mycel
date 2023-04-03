package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mycel/x/incentives/types"
)

// SetValidatorIncentive set a specific validatorIncentive in the store from its index
func (k Keeper) SetValidatorIncentive(ctx sdk.Context, validatorIncentive types.ValidatorIncentive) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorIncentiveKeyPrefix))
	b := k.cdc.MustMarshal(&validatorIncentive)
	store.Set(types.ValidatorIncentiveKey(
		validatorIncentive.Address,
	), b)
}

// GetValidatorIncentive returns a validatorIncentive from its index
func (k Keeper) GetValidatorIncentive(
	ctx sdk.Context,
	address string,

) (val types.ValidatorIncentive, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorIncentiveKeyPrefix))

	b := store.Get(types.ValidatorIncentiveKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidatorIncentive removes a validatorIncentive from the store
func (k Keeper) RemoveValidatorIncentive(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorIncentiveKeyPrefix))
	store.Delete(types.ValidatorIncentiveKey(
		address,
	))
}

// GetAllValidatorIncentive returns all validatorIncentive
func (k Keeper) GetAllValidatorIncentive(ctx sdk.Context) (list []types.ValidatorIncentive) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorIncentiveKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValidatorIncentive
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
