package keeper

import (
	"context"
	"time"

	"mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterDomain(goCtx context.Context, msg *types.MsgRegisterDomain) (*types.MsgRegisterDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now().In(time.UTC)
	expirationDate := currentTime.AddDate(int(msg.RegistrationPeriodInYear), 0, 0)

	domain := types.Domain{
		Name:           msg.Name,
		Owner:          msg.Creator,
		ExpirationDate: expirationDate.UnixNano(),
		Parent:         msg.Parent,
		DNSRecords:     nil,
		WalletRecords:  nil,
		Metadata:       nil,
	}

	registrationPeriodInWeek := uint(msg.RegistrationPeriodInYear * 12)

	err = k.Keeper.RegisterDomain(ctx, domain, creatorAddress, registrationPeriodInWeek)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterDomainResponse{}, nil
}
