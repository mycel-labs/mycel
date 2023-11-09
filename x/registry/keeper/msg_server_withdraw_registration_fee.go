package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) WithdrawRegistrationFee(goCtx context.Context, msg *types.MsgWithdrawRegistrationFee) (*types.MsgWithdrawRegistrationFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Get top level domain
	topLevelDomain, found := k.Keeper.GetTopLevelDomain(ctx, msg.Name)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrSecondLevelDomainNotFound, "%s", msg.Name)
	}

	if topLevelDomain.TotalWithdrawalAmount.IsZero() {
		return nil, errorsmod.Wrapf(types.ErrNoWithdrawalAmountToWithdraw, "%s", msg.Name)
	}

	// Check if the creator is the owner of the domain
	role, ok := topLevelDomain.AccessControl[msg.Creator]
	if !ok || role != types.DomainRole_OWNER {
		return nil, errorsmod.Wrapf(types.ErrNoPermissionToWithdraw, "%s", msg.Creator)
	}

	// Send coins from module account to Creator
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	registrationFee := topLevelDomain.TotalWithdrawalAmount
	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddress, registrationFee)
	if err != nil {
		return nil, err
	}
	topLevelDomain.TotalWithdrawalAmount = sdk.NewCoins()
	k.Keeper.SetTopLevelDomain(ctx, topLevelDomain)

	// Emit event
	EmitWithdrawRegistrationFeeEvent(ctx, *msg, registrationFee)

	return &types.MsgWithdrawRegistrationFeeResponse{RegistrationFee: registrationFee}, nil
}
