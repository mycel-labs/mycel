package keeper

import (
	"context"
	"errors"
	"fmt"

	"mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateDnsRecord(goCtx context.Context, msg *types.MsgUpdateDnsRecord) (*types.MsgUpdateDnsRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	domain, isFound := k.Keeper.GetDomain(ctx, msg.Name, msg.Parent)
	if !isFound {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s.%s", msg.Name, msg.Parent)), types.ErrDomainNotFound.Error())
	}

	// Check if the domain is owned by the creator
	if domain.Owner != msg.Creator {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s.%s", msg.Name, msg.Parent)), types.ErrDomainNotOwned.Error())
	}

	err := domain.UpdateDNSRecord(msg.DnsRecordType, msg.Value)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetDomain(ctx, domain)
	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeUpdateDNSRecord,
			sdk.NewAttribute(types.AttributeUpdateDNSRecordEventDomainName, msg.Name),
			sdk.NewAttribute(types.AttributeUpdateDNSRecordEventDomainParent, msg.Parent),
			sdk.NewAttribute(types.AttributeUpdateDNSRecordEventDNSRecordType, msg.DnsRecordType),
			sdk.NewAttribute(types.AttributeUpdateDNSRecordEventValue, msg.Value),
		),
	)

	return &types.MsgUpdateDnsRecordResponse{}, nil
}
