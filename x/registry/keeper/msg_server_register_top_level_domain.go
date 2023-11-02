package keeper

import (
	"context"

	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
)

func (k msgServer) RegisterTopLevelDomain(goCtx context.Context, msg *types.MsgRegisterTopLevelDomain) (*types.MsgRegisterTopLevelDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.RegistrationPeriodInYear < 1 || msg.RegistrationPeriodInYear > 4 {
		return nil, errorsmod.Wrapf(types.ErrInvalidRegistrationPeriod, "%d year(s)", msg.RegistrationPeriodInYear)
	}

	currentTime := ctx.BlockTime()
	expirationDate := currentTime.AddDate(0, 0, params.OneYearInDays*int(msg.RegistrationPeriodInYear))
	accessControl := map[string]types.DomainRole{
		msg.Creator: types.DomainRole_OWNER,
	}

	defaultRegistrationConfig := types.GetDefaultSubdomainConfig(3030)
	domain := types.TopLevelDomain{
		Name:                  msg.Name,
		ExpirationDate:        expirationDate,
		Metadata:              nil,
		SubdomainConfig:       &defaultRegistrationConfig,
		AccessControl:         accessControl,
		TotalWithdrawalAmount: sdk.NewCoins(),
	}

	fee, err := k.Keeper.RegisterTopLevelDomain(ctx, domain, msg.Creator, msg.RegistrationPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterTopLevelDomainResponse{
		TopLevelDomain: &domain,
		Fee:            &fee,
	}, nil
}
