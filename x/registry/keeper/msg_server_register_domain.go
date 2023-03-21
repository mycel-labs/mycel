package keeper

import (
	"context"
	"time"

	"mycel/x/registry/types"

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
	err = k.Keeper.ValidateIsDomainAlreadyTaken(ctx, domain)
	if err != nil {
		return nil, err
	}

	// Register domain
	domainLevel := domain.GetDomainLevel()

	switch domainLevel {
	case 1:
		err = k.Keeper.ValidateRegisterTLD(ctx, domain)
	default:
		err = k.Keeper.ValidateRegsiterSLD(ctx, domain)
	}
	if err != nil {
		return nil, err
	}

	// Store domain
	k.Keeper.SetDomain(ctx, domain)

	return &types.MsgRegisterDomainResponse{}, nil
}
