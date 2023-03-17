package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"mycel/x/mycel/types"
)

func (k msgServer) UpdateWalletRecord(goCtx context.Context, msg *types.MsgUpdateWalletRecord) (*types.MsgUpdateWalletRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateWalletRecordResponse{}, nil
}
