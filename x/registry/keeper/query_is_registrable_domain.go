package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IsRegistrableDomain(goCtx context.Context, req *types.QueryIsRegistrableDomainRequest) (*types.QueryIsRegistrableDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	domain := types.SecondLevelDomain{Name: req.Name, Parent: req.Parent}
	err := k.ValidateSecondLevelDomainIsRegistrable(ctx, domain)
	if err != nil {
		return &types.QueryIsRegistrableDomainResponse{IsRegstrable: false, ErrorMessage: err.Error()}, nil
	}

	return &types.QueryIsRegistrableDomainResponse{IsRegstrable: true, ErrorMessage: ""}, nil
}
