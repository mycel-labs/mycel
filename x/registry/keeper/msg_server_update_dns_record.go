package keeper

import (
	"context"

	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateDnsRecord(goCtx context.Context, msg *types.MsgUpdateDnsRecord) (*types.MsgUpdateDnsRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	domain, isFound := k.Keeper.GetSecondLevelDomain(ctx, msg.Name, msg.Parent)
	if !isFound {
		return nil, errorsmod.Wrapf(types.ErrSecondLevelDomainNotFound, "%s.%s", msg.Name, msg.Parent)
	}

	// Check if the domain is owned by the creator
	isEditable, err := domain.IsRecordEditable(msg.Creator)
	if !isEditable {
		return nil, err
	}

	err = domain.UpdateDnsRecord(msg.DnsRecordType, msg.Value)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetSecondLevelDomain(ctx, domain)

	// Emit event
	EmitUpdateDnsRecordEvent(ctx, *msg)

	return &types.MsgUpdateDnsRecordResponse{}, nil
}
