package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DomainRegistrationFee(goCtx context.Context, req *types.QueryDomainRegistrationFeeRequest) (*types.QueryDomainRegistrationFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	domain := types.SecondLevelDomain{Name: req.Name, Parent: req.Parent}
	config := k.GetSecondLevelDomainParentsSubdomainConfig(ctx, domain)
	fee, err := config.GetRegistrationFee(domain.Name, 1)
	if err != nil {
		return nil, err
	}

	return &types.QueryDomainRegistrationFeeResponse{Fee: fee}, nil
}
