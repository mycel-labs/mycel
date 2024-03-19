package keeper

import (
	"context"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) RegisterSecondLevelDomain(goCtx context.Context, msg *types.MsgRegisterSecondLevelDomain) (*types.MsgRegisterSecondLevelDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	accessControl := types.AccessControl{
		Address: msg.Creator,
		Role:    types.DomainRole_OWNER,
	}
	domain := types.SecondLevelDomain{
		Name:           msg.Name,
		Owner:          msg.Creator,
		ExpirationDate: time.Time{},
		Parent:         msg.Parent,
		Records:        nil,
		AccessControl:  []*types.AccessControl{&accessControl},
	}

	err = k.Keeper.RegisterSecondLevelDomain(ctx, domain, creatorAddress, msg.RegistrationPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterSecondLevelDomainResponse{}, nil
}

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

func (k msgServer) UpdateTextRecord(goCtx context.Context, msg *types.MsgUpdateTextRecord) (*types.MsgUpdateTextRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx
	domain, isFound := k.Keeper.GetSecondLevelDomain(ctx, msg.Name, msg.Parent)
	if !isFound {
		return nil, errorsmod.Wrapf(types.ErrSecondLevelDomainNotFound, "%s.%s", msg.Name, msg.Parent)
	}

	// Check if the domain is owned by the creator
	isEditable, err := domain.IsRecordEditable(msg.Creator)
	if !isEditable {
		return nil, err
	}

	err = domain.UpdateTextRecord(msg.Key, msg.Value)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetSecondLevelDomain(ctx, domain)

	// Emit event
	EmitUpdateTextRecordEvent(ctx, *msg)

	return &types.MsgUpdateTextRecordResponse{}, nil
}

func (k msgServer) UpdateWalletRecord(goCtx context.Context, msg *types.MsgUpdateWalletRecord) (*types.MsgUpdateWalletRecordResponse, error) {
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

	err = domain.UpdateWalletRecord(msg.WalletRecordType, msg.Value)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetSecondLevelDomain(ctx, domain)

	// Emit event
	EmitUpdateWalletRecordEvent(ctx, *msg)

	return &types.MsgUpdateWalletRecordResponse{}, nil
}

func (k msgServer) UpdateTopLevelDomainRegistrationPolicy(goCtx context.Context, msg *types.MsgUpdateTopLevelDomainRegistrationPolicy) (*types.MsgUpdateTopLevelDomainRegistrationPolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.UpdateTopLevelDomainRegistrationPolicy(ctx, msg.Creator, msg.Name, msg.RegistrationPolicy)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateTopLevelDomainRegistrationPolicyResponse{}, nil
}
