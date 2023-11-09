package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mycel-domain/mycel/x/epochs/types"
)

func (k Keeper) EpochInfoAll(goCtx context.Context, req *types.QueryAllEpochInfoRequest) (*types.QueryAllEpochInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var epochInfos []types.EpochInfo
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	epochInfoStore := prefix.NewStore(store, types.KeyPrefix(types.EpochInfoKeyPrefix))

	pageRes, err := query.Paginate(epochInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var epochInfo types.EpochInfo
		if err := k.cdc.Unmarshal(value, &epochInfo); err != nil {
			return err
		}

		epochInfos = append(epochInfos, epochInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEpochInfoResponse{EpochInfo: epochInfos, Pagination: pageRes}, nil
}

func (k Keeper) EpochInfo(goCtx context.Context, req *types.QueryGetEpochInfoRequest) (*types.QueryGetEpochInfoResponse, error) {
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

	return &types.QueryGetEpochInfoResponse{EpochInfo: val}, nil
}
