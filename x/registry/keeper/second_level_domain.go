package keeper

import (
	"github.com/mycel-domain/mycel/x/registry/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSecondLevelDomain set a specific second-level-domain in the store from its index
func (k Keeper) SetSecondLevelDomain(ctx sdk.Context, domain types.SecondLevelDomain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecondLevelDomainKeyPrefix))
	b := k.cdc.MustMarshal(&domain)
	store.Set(types.SecondLevelDomainKey(
		domain.Name,
		domain.Parent,
	), b)
}

// GetSecondLevelDomain returns a second-level-domain from its index
func (k Keeper) GetSecondLevelDomain(
	ctx sdk.Context,
	name string,
	parent string,

) (val types.SecondLevelDomain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecondLevelDomainKeyPrefix))

	b := store.Get(types.SecondLevelDomainKey(
		name,
		parent,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecondLevelDomain removes a second-level-domain from the store
func (k Keeper) RemoveSecondLevelDomain(
	ctx sdk.Context,
	name string,
	parent string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecondLevelDomainKeyPrefix))
	store.Delete(types.SecondLevelDomainKey(
		name,
		parent,
	))
}

// GetAllSecondLevelDomain returns all second-level-domain
func (k Keeper) GetAllSecondLevelDomain(ctx sdk.Context) (list []types.SecondLevelDomain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecondLevelDomainKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecondLevelDomain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
