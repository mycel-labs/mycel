package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	registrytypes "github.com/mycel-domain/mycel/x/registry/types"
	"github.com/mycel-domain/mycel/x/resolver/types"
)

func (k Keeper) TextRecord(goCtx context.Context, req *types.QueryTextRecordRequest) (*types.QueryTextRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate request parameters
	err := registrytypes.ValidateTextRecordKey(req.Key)
	if err != nil {
		return nil, err
	}

	// Query domain record
	_, err = k.registryKeeper.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.registryKeeper.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}

	value := secondLevelDomain.GetTextRecord(req.Key)

	return &types.QueryTextRecordResponse{
		Value: &registrytypes.TextRecord{Key: req.Key, Value: value},
	}, nil
}
