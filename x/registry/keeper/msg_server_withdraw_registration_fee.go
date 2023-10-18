package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) WithdrawRegistrationFee(goCtx context.Context, msg *types.MsgWithdrawRegistrationFee) (*types.MsgWithdrawRegistrationFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWithdrawRegistrationFeeResponse{}, nil
}
