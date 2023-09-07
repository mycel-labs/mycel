package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) RegisterTopLevelDomain(goCtx context.Context, msg *types.MsgRegisterTopLevelDomain) (*types.MsgRegisterTopLevelDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterTopLevelDomainResponse{}, nil
}
