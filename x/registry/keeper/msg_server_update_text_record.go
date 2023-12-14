package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) UpdateTextRecord(goCtx context.Context, msg *types.MsgUpdateTextRecord) (*types.MsgUpdateTextRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateTextRecordResponse{}, nil
}
