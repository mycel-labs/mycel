package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/epochs/types"
)

func (k Keeper) EpochInfos(goCtx context.Context, req *types.QueryEpochsInfoRequest) (*types.QueryEpochsInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryEpochsInfoResponse{Epochs: k.GetAllEpochInfo(ctx)}, nil
}

func (k Keeper) CurrentEpoch(goCtx context.Context, req *types.QueryCurrentEpochRequest) (*types.QueryCurrentEpochResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEpochInfo(
		ctx,
		req.Identifier,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryCurrentEpochResponse{EpochInfo: val}, nil
}
