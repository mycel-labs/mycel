package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mycel-domain/mycel/x/resolver/types"
)

func (k Keeper) AllRecords(goCtx context.Context, req *types.QueryAllRecordsRequest) (*types.QueryAllRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Query domain record
	_, err := k.registryKeeper.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.registryKeeper.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}

	values := secondLevelDomain.Records

	return &types.QueryAllRecordsResponse{Values: values}, nil
}
