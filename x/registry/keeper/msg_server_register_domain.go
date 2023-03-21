package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"mycel/x/registry/types"
)

func (k msgServer) RegisterDomain(goCtx context.Context, msg *types.MsgRegisterDomain) (*types.MsgRegisterDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterDomainResponse{}, nil
}
