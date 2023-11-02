package keeper

import (
	"context"

	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterTopLevelDomain(goCtx context.Context, msg *types.MsgRegisterTopLevelDomain) (*types.MsgRegisterTopLevelDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.RegistrationPeriodInYear < 1 || msg.RegistrationPeriodInYear > 4 {
		return nil, errorsmod.Wrapf(types.ErrTopLevelDomainInvalidRegistrationPeriod, "%d year(s)", msg.RegistrationPeriodInYear)
	}

	topLevelDomain, fee, err := k.Keeper.RegisterTopLevelDomain(ctx, msg.Creator, msg.Name, msg.RegistrationPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterTopLevelDomainResponse{
		TopLevelDomain: &topLevelDomain,
		Fee:            &fee,
	}, nil
}
