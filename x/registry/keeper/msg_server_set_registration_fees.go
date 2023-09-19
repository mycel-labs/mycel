package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) SetRegistrationFees(goCtx context.Context, msg *types.MsgSetRegistrationFees) (*types.MsgSetRegistrationFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetRegistrationFeesResponse{}, nil
}
