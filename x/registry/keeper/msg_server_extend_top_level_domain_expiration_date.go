package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) ExtendTopLevelDomainExpirationDate(goCtx context.Context, msg *types.MsgExtendTopLevelDomainExpirationDate) (*types.MsgExtendTopLevelDomainExpirationDateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	domain, found := k.GetTopLevelDomain(ctx, msg.Name)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrDomainNotFound, "%s", msg.Name)
	}

	// Check if the domain is editable
	_, err := domain.IsEditable(msg.Creator)
	if err != nil {
		return nil, err
	}

	fee, err := k.Keeper.ExtendTopLevelDomainExpirationDate(ctx, msg.Creator, &domain, msg.ExtensionPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgExtendTopLevelDomainExpirationDateResponse{
		TopLevelDomain: &domain,
		Fee:            &fee,
	}, nil
}
