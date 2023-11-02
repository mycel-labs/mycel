package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) ExtendTopLevelDomainExpiration(goCtx context.Context, msg *types.MsgExtendTopLevelDomainExpiration) (*types.MsgExtendTopLevelDomainExpirationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgExtendTopLevelDomainExpirationResponse{}, nil
}
