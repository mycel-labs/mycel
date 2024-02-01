package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k msgServer) SubmitTopLevelDomainProposal(goCtx context.Context, msg *types.MsgSubmitTopLevelDomainProposal) (*types.MsgSubmitTopLevelDomainProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitTopLevelDomainProposalResponse{}, nil
}
