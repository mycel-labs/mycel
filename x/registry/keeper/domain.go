package keeper

import (
	"errors"
	"fmt"
	"mycel/x/registry/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) GetIsDomainAlreadyTaken(ctx sdk.Context, domain types.Domain) (isDomainAlreadyTaken bool) {
	_, isDomainAlreadyTaken = k.GetDomain(ctx, domain.Name, domain.Parent)
	return isDomainAlreadyTaken
}

func (k Keeper) GetIsParentDomainExist(ctx sdk.Context, domain types.Domain) (isParentDomainExist bool) {
	name, parent := domain.ParseParent()
	_, isParentDomainExist = k.GetDomain(ctx, name, parent)
	return isParentDomainExist
}

func (k Keeper) ValidateIsDomainAlreadyTaken(ctx sdk.Context, domain types.Domain) (err error) {
	isDomainAlreadyTaken := k.GetIsDomainAlreadyTaken(ctx, domain)
	if isDomainAlreadyTaken {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s.%s", domain.Name, domain.Parent)),
			types.ErrDomainIsAlreadyTaken.Error())
	}
	return err
}

func (k Keeper) ValidateRegisterTLD(ctx sdk.Context, domain types.Domain) (err error) {
	if domain.Parent != "" {
		err = sdkerrors.Wrapf(errors.New(domain.Parent),
			types.ErrParentDomainMustBeEmpty.Error())
	}
	// TODO: Is Staked enough to register TLD
	return err
}

func (k Keeper) ValidateRegsiterSLD(ctx sdk.Context, domain types.Domain) (err error) {
	isParentDomainExist := k.GetIsParentDomainExist(ctx, domain)
	if !isParentDomainExist {
		err = sdkerrors.Wrapf(errors.New(domain.Parent),
			types.ErrParentDomainDoesNotExist.Error())
	}
	return err
}
