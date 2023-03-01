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

	newDomain := types.Domain{
		Name:           msg.Name,
		Owner:          msg.Creator,
		ExpirationDate: expirationDate.UnixNano(),
		Parent:         msg.Parent,
		DNSRecords:     nil,
		WalletRecords:  nil,
		Metadata:       nil,
	}

	err := newDomain.ValidateDomain()
	if err != nil {
		return nil, err
	}
	_, err = k.Keeper.GetIsDomainAlreadyTaken(ctx, newDomain.Name, newDomain.Parent)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetDomain(ctx, newDomain)

	return &types.MsgRegisterDomainResponse{}, nil
}

func (k msgServer) RegisterTLD(goCtx context.Context, domain types.Domain) (err error) {
	return err
}

func (k msgServer) RegisterRootDomain(goCtx context.Context, domain types.Domain) (err error) {
	return err
}

func (k msgServer) RegisterSubDomain(goCtx context.Context, domain types.Domain) (err error) {
	return err
}
