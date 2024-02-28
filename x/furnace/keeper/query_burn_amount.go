package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/mycel-domain/mycel/x/furnace/types"
)

func (k Keeper) BurnAmountAll(ctx context.Context, req *types.QueryAllBurnAmountRequest) (*types.QueryAllBurnAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var burnAmounts []types.BurnAmount

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnAmountKeyPrefix))

	burnAmountStore := prefix.NewStore(store, types.KeyPrefix(types.BurnAmountKeyPrefix))

	pageRes, err := query.Paginate(burnAmountStore, req.Pagination, func(key []byte, value []byte) error {
		var burnAmount types.BurnAmount
		if err := k.cdc.Unmarshal(value, &burnAmount); err != nil {
			return err
		}

		burnAmounts = append(burnAmounts, burnAmount)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBurnAmountResponse{BurnAmount: burnAmounts, Pagination: pageRes}, nil
}

func (k Keeper) BurnAmount(ctx context.Context, req *types.QueryGetBurnAmountRequest) (*types.QueryGetBurnAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetBurnAmount(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBurnAmountResponse{BurnAmount: val}, nil
}
