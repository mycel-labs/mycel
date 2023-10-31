package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterSecondLevelDomain(goCtx context.Context, msg *types.MsgRegisterSecondLevelDomain) (*types.MsgRegisterSecondLevelDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.RegistrationPeriodInYear < 1 || msg.RegistrationPeriodInYear > 4 {
		return nil, errorsmod.Wrapf(errors.New(fmt.Sprintf("%d year(s)", msg.RegistrationPeriodInYear)), types.ErrInvalidRegistrationPeriod.Error())
	}

	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	currentTime := ctx.BlockTime()
	expirationDate := currentTime.AddDate(int(msg.RegistrationPeriodInYear), 0, 0)
	accessControl := map[string]types.DomainRole{
		msg.Creator: types.DomainRole_OWNER,
	}

	domain := types.SecondLevelDomain{
		Name:           msg.Name,
		Owner:          msg.Creator,
		ExpirationDate: expirationDate.UnixNano(),
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
