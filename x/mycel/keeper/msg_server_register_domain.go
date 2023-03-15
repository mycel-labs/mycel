package keeper

import (
	"context"
	"time"

	"mycel/x/mycel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterDomain(goCtx context.Context, msg *types.MsgRegisterDomain) (*types.MsgRegisterDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentTime := time.Now()
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

	// Validate domain
	err := domain.ValidateDomain()
	if err != nil {
		return nil, err
	}
	_, err = k.Keeper.GetIsDomainAlreadyTaken(ctx, domain.Name, domain.Parent)
	if err != nil {
		return nil, err
	}

	// Register domain
	domainLevel := domain.GetDomainLevel()

	switch domainLevel {
	case 1:
		k.RegisterSubDomainValidate(ctx, domain)
	case 2:
		k.RegisterSecondLevelDomainValidate(ctx, domain)
	default:
		k.RegisterSubDomainValidate(ctx, domain)
	}

	// Store domain
	k.Keeper.SetDomain(ctx, domain)

	return &types.MsgRegisterDomainResponse{}, nil
}

func (k msgServer) RegisterTopLevelDomainValidate(ctx sdk.Context, domain types.Domain) (err error) {
	// TODO: check the validator is alive and send token as registration fee
	return err
}

func (k msgServer) RegisterSecondLevelDomainValidate(ctx sdk.Context, domain types.Domain) (err error) {
	// TODO: send token as registration fee
	return err
}

func (k msgServer) RegisterSubDomainValidate(ctx sdk.Context, domain types.Domain) (err error) {
	// TODO: no fee
	return err
}
