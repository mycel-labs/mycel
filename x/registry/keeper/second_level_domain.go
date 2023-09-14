package keeper

import (
	"errors"
	"fmt"
	"github.com/mycel-domain/mycel/x/registry/types"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// Get valid second level domain
func (k Keeper) GetValidSecondLevelDomain(ctx sdk.Context, name string, parent string) (secondLevelDomain types.SecondLevelDomain, err error) {
	// Regex validation
	err = types.ValidateSecondLevelDomainName(name)
	if err != nil {
		return secondLevelDomain, err
	}
	err = types.ValidateSecondLevelDomainParent(parent)
	if err != nil {
		return secondLevelDomain, err
	}
	// Get second level domain
	secondLevelDomain, isFound := k.GetSecondLevelDomain(ctx, name, parent)
	if !isFound {
		return secondLevelDomain, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s.%s", name, parent)), types.ErrDomainNotFound.Error())
	}

	// Check if domain is not expired
	if time.Unix(secondLevelDomain.ExpirationDate, 0).Before(ctx.BlockTime()) && secondLevelDomain.ExpirationDate != 0 {
		return secondLevelDomain, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", name)), types.ErrDomainExpired.Error())
	}

	return secondLevelDomain, nil
}
