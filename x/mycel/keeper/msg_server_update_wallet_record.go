package keeper

import (
	"context"

	"mycel/x/mycel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateWalletRecord(goCtx context.Context, msg *types.MsgUpdateWalletRecord) (*types.MsgUpdateWalletRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	domain, isFound := k.GetDomain(ctx, msg.Name, msg.Parent)
	if !isFound {
		return nil, types.ErrDomainNotFound
	}

	// Check if the domain is owned by the creator
	if domain.Owner != msg.Creator {
		return nil, types.ErrDomainNotOwned
	}

	domain.UpdateWalletRecord(msg.WalletRecordType, msg.Value)

	return &types.MsgUpdateWalletRecordResponse{}, nil
}
