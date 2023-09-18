package keeper

import (
	"errors"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// SetTopLevelDomain set a specific topLevelDomain in the store from its index
func (k Keeper) SetTopLevelDomain(ctx sdk.Context, topLevelDomain types.TopLevelDomain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopLevelDomainKeyPrefix))
	b := k.cdc.MustMarshal(&topLevelDomain)
	store.Set(types.TopLevelDomainKey(
		topLevelDomain.Name,
	), b)
}

// GetTopLevelDomain returns a topLevelDomain from its index
func (k Keeper) GetTopLevelDomain(
	ctx sdk.Context,
	name string,

) (val types.TopLevelDomain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopLevelDomainKeyPrefix))

	b := store.Get(types.TopLevelDomainKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTopLevelDomain removes a topLevelDomain from the store
func (k Keeper) RemoveTopLevelDomain(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopLevelDomainKeyPrefix))
	store.Delete(types.TopLevelDomainKey(
		name,
	))
}

// GetAllTopLevelDomain returns all topLevelDomain
func (k Keeper) GetAllTopLevelDomain(ctx sdk.Context) (list []types.TopLevelDomain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopLevelDomainKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TopLevelDomain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Get valid top level domain

func (k Keeper) GetValidTopLevelDomain(ctx sdk.Context, name string) (topLevelDomain types.TopLevelDomain, err error) {
	// Regex validation
	err = types.ValidateTopLevelDomainName(name)
	if err != nil {
		return topLevelDomain, err
	}

	// Get top level domain
	topLevelDomain, isFound := k.GetTopLevelDomain(ctx, name)
	if !isFound {
		return topLevelDomain, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", name)), types.ErrDomainNotFound.Error())
	}

	// Check if domain is not expired
	expirationDate := time.Unix(0, topLevelDomain.ExpirationDate)
	if ctx.BlockTime().After(expirationDate) && topLevelDomain.ExpirationDate != 0 {
		return topLevelDomain, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", name)), types.ErrDomainExpired.Error())
	}

	return topLevelDomain, nil
}
