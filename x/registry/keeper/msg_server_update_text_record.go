package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) UpdateTextRecord(goCtx context.Context, msg *types.MsgUpdateTextRecord) (*types.MsgUpdateTextRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx
	domain, isFound := k.Keeper.GetSecondLevelDomain(ctx, msg.Name, msg.Parent)
	if !isFound {
		return nil, errorsmod.Wrapf(types.ErrSecondLevelDomainNotFound, "%s.%s", msg.Name, msg.Parent)
	}

	// Check if the domain is owned by the creator
	isEditable, err := domain.IsRecordEditable(msg.Creator)
	if !isEditable {
		return nil, err
	}

	err = domain.UpdateTextRecord(msg.Key, msg.Value)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetSecondLevelDomain(ctx, domain)

	// Emit event
	EmitUpdateTextRecordEvent(ctx, *msg)

	return &types.MsgUpdateTextRecordResponse{}, nil
}
