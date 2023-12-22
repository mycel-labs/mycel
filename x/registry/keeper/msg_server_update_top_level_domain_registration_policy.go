package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) UpdateTopLevelDomainRegistrationPolicy(goCtx context.Context, msg *types.MsgUpdateTopLevelDomainRegistrationPolicy) (*types.MsgUpdateTopLevelDomainRegistrationPolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.UpdateTopLevelDomainRegistrationPolicy(ctx, msg.Creator, msg.Name, msg.RegistrationPolicy)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateTopLevelDomainRegistrationPolicyResponse{}, nil
}
