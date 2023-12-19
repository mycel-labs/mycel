package keeper

import (
	"context"
	"strings"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k Keeper) Role(goCtx context.Context, req *types.QueryRoleRequest) (*types.QueryRoleResponse, error) {
	// Return error when the request is empty
	if req == nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request: empty request")
	}

	// Setup Variables
	var (
		role  types.DomainRole
		found bool
	)
	ctx := sdk.UnwrapSDKContext(goCtx)
	dms := strings.Split(req.DomainName, ".") // "foo.cel" will be [ "foo", "cel" ]

	// Get Role
	switch len(dms) {
	case 1: // TLD
		role, found = k.GetTopLevelDomainRole(ctx, dms[0], req.Address)
	case 2: // SLD
		role, found = k.GetSecondLevelDomainRole(ctx, dms[0], dms[1], req.Address)
	default:
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request: domain name")
	}

	// Return results
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "domain not found")
	}
	return &types.QueryRoleResponse{Role: role.String()}, nil
}
