package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/mycel-domain/mycel/x/epochs/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EpochInfoAll(ctx context.Context, req *types.QueryAllEpochInfoRequest) (*types.QueryAllEpochInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var epochInfos []types.EpochInfo

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
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

func (k Keeper) EpochInfo(ctx context.Context, req *types.QueryGetEpochInfoRequest) (*types.QueryGetEpochInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetEpochInfo(
		ctx,
		req.Identifier,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEpochInfoResponse{EpochInfo: val}, nil
}
