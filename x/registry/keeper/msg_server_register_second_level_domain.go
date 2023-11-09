package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) RegisterSecondLevelDomain(goCtx context.Context, msg *types.MsgRegisterSecondLevelDomain) (*types.MsgRegisterSecondLevelDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	accessControl := map[string]types.DomainRole{
		msg.Creator: types.DomainRole_OWNER,
	}

	domain := types.SecondLevelDomain{
		Name:           msg.Name,
		Owner:          msg.Creator,
		ExpirationDate: time.Time{},
		Parent:         msg.Parent,
		Records:        nil,
		AccessControl:  accessControl,
	}

	err = k.Keeper.RegisterSecondLevelDomain(ctx, domain, creatorAddress, msg.RegistrationPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterSecondLevelDomainResponse{}, nil
}
