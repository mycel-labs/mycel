package keeper

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	registrytypes "github.com/mycel-domain/mycel/x/registry/types"
)

func (k Keeper) GetTopLevelDomain(ctx sdk.Context, name string) (topLevelDomain registrytypes.TopLevelDomain, err error) {
	// Regex validation
	err = registrytypes.ValidateTopLevelDomainName(name)
	if err != nil {
		return topLevelDomain, err
	}

	// Get top level domain
	topLevelDomain, isFound := k.registryKeeper.GetTopLevelDomain(ctx, name)
	if !isFound {
		return topLevelDomain, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", name)), registrytypes.ErrDomainNotFound.Error())
	}

	// Check if domain is valid


	return topLevelDomain, nil
}

func (k Keeper) GetSecondLevelDomain(ctx sdk.Context, name string, parent string) (secondLevelDomain registrytypes.SecondLevelDomain, err error) {
	// Regex validation
	err = registrytypes.ValidateSecondLevelDomainName(name)
	if err != nil {
		return secondLevelDomain, err
	}
	err = registrytypes.ValidateSecondLevelDomainParent(parent)
	if err != nil {
		return secondLevelDomain, err
	}
	// Get second level domain
	secondLevelDomain, isFound := k.registryKeeper.GetSecondLevelDomain(ctx, name, parent)
	if !isFound {
		return secondLevelDomain, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s.%s", name, parent)), registrytypes.ErrDomainNotFound.Error())
	}
	// Check if domain is valid

	return secondLevelDomain, nil
}
