package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) ExtendTopLevelDomainExpirationDate(goCtx context.Context, msg *types.MsgExtendTopLevelDomainExpirationDate) (*types.MsgExtendTopLevelDomainExpirationDateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	topLevelDomain, fee, err := k.Keeper.ExtendTopLevelDomainExpirationDate(ctx, msg.Creator, msg.Name, msg.ExtensionPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgExtendTopLevelDomainExpirationDateResponse{
		TopLevelDomain: &topLevelDomain,
		Fee:            &fee,
	}, nil
}
