package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) SubmitTopLevelDomainProposal(goCtx context.Context, msg *types.MsgSubmitTopLevelDomainProposal) (*types.MsgSubmitTopLevelDomainProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// NOTE: Same as x/registry/keeper/msg_server_register_top_level_domain.go
	if msg.RegistrationPeriodInYear < 1 || msg.RegistrationPeriodInYear > 4 {
		return nil, errorsmod.Wrapf(types.ErrTopLevelDomainInvalidRegistrationPeriod, "%d year(s)", msg.RegistrationPeriodInYear)
	}

	topLevelDomain, fee, err := k.Keeper.RegisterTopLevelDomain(ctx, msg.Creator, msg.Name, msg.RegistrationPeriodInYear)
	if err != nil {
		return nil, err
	}

	return &types.MsgSubmitTopLevelDomainProposalResponse{
		TopLevelDomain: &topLevelDomain,
		Fee:            &fee,
	}, nil
}
