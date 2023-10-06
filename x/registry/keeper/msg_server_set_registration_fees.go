package keeper

import (
	"context"
	"errors"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) SetRegistrationFees(goCtx context.Context, msg *types.MsgSetRegistrationFees) (*types.MsgSetRegistrationFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get TLD
	domain, isFound := k.Keeper.GetTopLevelDomain(ctx, msg.Domain)
	if !isFound {
		return nil, errorsmod.Wrapf(errors.New(fmt.Sprintf("%s", msg.Domain)), types.ErrDomainNotFound.Error())
	}

	// Check if domain is editable
	if domain.AccessControl[msg.Creator] != types.DomainRole_OWNER {
		return nil, errorsmod.Wrapf(errors.New(fmt.Sprintf("%s", msg.Creator)), types.ErrDomainNotEditable.Error())
	}

	// Check if request is empty

	// Set fees
	domain.SetRegistrationFees(msg.FeesByName, msg.FeesByLength, msg.DefaultFee)

	// Validate fees
	err := domain.ValidateRegistrationFees()
	if err != nil {
		return nil, err
	}

	// Store fees
	k.Keeper.SetTopLevelDomain(ctx, domain)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeSetRegistrationFees,
			sdk.NewAttribute(types.AttributeSetRegistrationFeesDomain, msg.Domain),
		),
	)

	return &types.MsgSetRegistrationFeesResponse{}, nil
}
