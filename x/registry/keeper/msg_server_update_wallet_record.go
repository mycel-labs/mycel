package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateWalletRecord(goCtx context.Context, msg *types.MsgUpdateWalletRecord) (*types.MsgUpdateWalletRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	domain, isFound := k.Keeper.GetDomain(ctx, msg.Name, msg.Parent)
	if !isFound {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s.%s", msg.Name, msg.Parent)), types.ErrDomainNotFound.Error())
	}

	// Check if the domain is owned by the creator
	if domain.Owner != msg.Creator {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s.%s", msg.Name, msg.Parent)), types.ErrDomainNotOwned.Error())
	}

	err := domain.UpdateWalletRecord(msg.WalletRecordType, msg.Value)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetDomain(ctx, domain)
	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeUpdateWalletRecord,
			sdk.NewAttribute(types.AttributeUpdateWalletRecordEventDomainName, msg.Name),
			sdk.NewAttribute(types.AttributeRegisterDomainEventParent, msg.Parent),
			sdk.NewAttribute(types.AttributeUpdateWalletRecordEventWalletRecordType, msg.WalletRecordType),
			sdk.NewAttribute(types.AttributeUpdateWalletRecordEventValue, msg.Value),
		),
	)

	return &types.MsgUpdateWalletRecordResponse{}, nil
}
