package keeper

import (
	"context"

	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateWalletRecord(goCtx context.Context, msg *types.MsgUpdateWalletRecord) (*types.MsgUpdateWalletRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	domain, isFound := k.Keeper.GetSecondLevelDomain(ctx, msg.Name, msg.Parent)
	if !isFound {
		return nil, errorsmod.Wrapf(types.ErrDomainNotFound, "%s.%s", msg.Name, msg.Parent)
	}

	// Check if the domain is owned by the creator
	isEditable, err := domain.IsRecordEditable(msg.Creator)
	if !isEditable {
		return nil, err
	}

	err = domain.UpdateWalletRecord(msg.WalletRecordType, msg.Value)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetSecondLevelDomain(ctx, domain)

	// Emit event
	EmitUpdateWalletRecordEvent(ctx, *msg)


	return &types.MsgUpdateWalletRecordResponse{}, nil
}
