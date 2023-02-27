package keeper

import (
	"context"
	"time"

	"mycel/x/mycel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDomain(goCtx context.Context, msg *types.MsgCreateDomain) (*types.MsgCreateDomainResponse, error) {
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

	k.Keeper.SetDomain(ctx, newDomain)

	return &types.MsgCreateDomainResponse{}, nil
}
