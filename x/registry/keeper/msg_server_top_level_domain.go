package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

// RegisterTopLevelDomain registers a new top-level domain
func (k msgServer) RegisterTopLevelDomain(goCtx context.Context, msg *types.MsgRegisterTopLevelDomain) (*types.MsgRegisterTopLevelDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.RegistrationPeriodInYear < 1 || msg.RegistrationPeriodInYear > 4 {
		return nil, errorsmod.Wrapf(types.ErrTopLevelDomainInvalidRegistrationPeriod, "%d year(s)", msg.RegistrationPeriodInYear)
	}

	topLevelDomain, fee, err := k.Keeper.RegisterTopLevelDomain(ctx, msg.Creator, msg.Name, msg.RegistrationPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterTopLevelDomainResponse{
		TopLevelDomain: &topLevelDomain,
		Fee:            &fee,
	}, nil
}

// ExtendTopLevelDomainExpirationDate extends the expiration date of a top-level domain
func (k msgServer) ExtendTopLevelDomainExpirationDate(goCtx context.Context, msg *types.MsgExtendTopLevelDomainExpirationDate) (*types.MsgExtendTopLevelDomainExpirationDateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	topLevelDomain, fee, err := k.Keeper.ExtendTopLevelDomainExpirationDate(ctx, msg.Creator, msg.Name, msg.ExtensionPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgExtendTopLevelDomainExpirationDateResponse{
		TopLevelDomain: &topLevelDomain,
		Fee:            &fee,
	}, nil
}

// WithdrawRegistrationFee withdraws the registration fee of a sub-domain
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
	role := topLevelDomain.GetRole(msg.Creator)
	if role != types.DomainRole_OWNER {
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
