package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/furnace/types"
)

func (k Keeper) EpochBurnConfig(goCtx context.Context, req *types.QueryGetEpochBurnConfigRequest) (*types.QueryGetEpochBurnConfigResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEpochBurnConfig(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEpochBurnConfigResponse{EpochBurnConfig: val}, nil
}
