package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mycel-domain/mycel/x/furnace/types"
)

func (k Keeper) EpochBurnConfig(ctx context.Context, req *types.QueryGetEpochBurnConfigRequest) (*types.QueryGetEpochBurnConfigResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetEpochBurnConfig(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEpochBurnConfigResponse{EpochBurnConfig: val}, nil
}
