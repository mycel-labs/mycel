package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/mycel-domain/mycel/x/furnace/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BurnAmountAll(goCtx context.Context, req *types.QueryAllBurnAmountRequest) (*types.QueryAllBurnAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var burnAmounts []types.BurnAmount
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
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

func (k Keeper) BurnAmount(goCtx context.Context, req *types.QueryGetBurnAmountRequest) (*types.QueryGetBurnAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetBurnAmount(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBurnAmountResponse{BurnAmount: val}, nil
}
