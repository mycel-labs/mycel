package keeper

import (
	"mycel/x/mycel/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDomain set a specific domain in the store from its index
func (k Keeper) SetDomain(ctx sdk.Context, domain types.Domain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DomainKeyPrefix))
	b := k.cdc.MustMarshal(&domain)
	store.Set(types.DomainKey(
		domain.Name,
		domain.Parent,
	), b)
}

// GetDomain returns a domain from its index
func (k Keeper) GetDomain(
	ctx sdk.Context,
	name string,
	parent string,

) (val types.Domain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DomainKeyPrefix))

	b := store.Get(types.DomainKey(
		name,
		parent,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDomain removes a domain from the store
func (k Keeper) RemoveDomain(
	ctx sdk.Context,
	name string,
	parent string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DomainKeyPrefix))
	store.Delete(types.DomainKey(
		name,
		parent,
	))
}

// GetAllDomain returns all domain
func (k Keeper) GetAllDomain(ctx sdk.Context) (list []types.Domain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DomainKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Domain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
