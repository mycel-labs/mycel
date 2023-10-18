package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterTopLevelDomain(goCtx context.Context, msg *types.MsgRegisterTopLevelDomain) (*types.MsgRegisterTopLevelDomainResponse, error) {
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

	defaultRegistrationConfig := types.GetDefaultSubdomainConfig(3030)
	domain := types.TopLevelDomain{
		Name:             msg.Name,
		ExpirationDate:   expirationDate.UnixNano(),
		Metadata:         nil,
		SubdomainConfig:  &defaultRegistrationConfig,
		AccessControl:    accessControl,
		RegistrationFee: sdk.NewCoins(),
	}

	err = k.Keeper.RegisterTopLevelDomain(ctx, domain, creatorAddress, msg.RegistrationPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterTopLevelDomainResponse{}, nil
}
