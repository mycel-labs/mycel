package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) WithdrawRegistrationFee(goCtx context.Context, msg *types.MsgWithdrawRegistrationFee) (*types.MsgWithdrawRegistrationFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// Get top level domain
	topLevelDomain, found := k.Keeper.GetTopLevelDomain(ctx, msg.Name)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrDomainNotFound, "%s", msg.Name)
	}

	if topLevelDomain.RegistrationFees.IsZero() {
		// return nil, errorsmod.Wrapf(types.ErrNoRegistrationFeesToWithdraw, "%s", msg.Name)
	}

	// Check if the creator is the owner of the domain
	role, ok := topLevelDomain.AccessControl[msg.Creator]
	if !ok || role != types.DomainRole_OWNER {
		return nil, errorsmod.Wrapf(types.ErrNoPermissionToWithdrawFees, "%s", msg.Creator)
	}

	// Send coins from module account to Creator
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	registrationFees := topLevelDomain.RegistrationFees
	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddress, registrationFees)
	if err != nil {
		return nil, err
	}
	topLevelDomain.RegistrationFees = sdk.NewCoins()
	k.Keeper.SetTopLevelDomain(ctx, topLevelDomain)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeWithdrawRegistrationFees,
			sdk.NewAttribute(types.AttributeWithdrawRegistrationFeesEventDomainName, msg.Name),
			sdk.NewAttribute(types.AttributeWithdrawRegistrationFeesEventDomainFees, registrationFees.String()),
		),
	)

	return &types.MsgWithdrawRegistrationFeeResponse{RegistrationFees: registrationFees}, nil
}
