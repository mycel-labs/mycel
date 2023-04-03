package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mycel/x/incentives/types"
)

func (k Keeper) DelegetorIncentiveAll(goCtx context.Context, req *types.QueryAllDelegetorIncentiveRequest) (*types.QueryAllDelegetorIncentiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var delegetorIncentives []types.DelegetorIncentive
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	delegetorIncentiveStore := prefix.NewStore(store, types.KeyPrefix(types.DelegetorIncentiveKeyPrefix))

	pageRes, err := query.Paginate(delegetorIncentiveStore, req.Pagination, func(key []byte, value []byte) error {
		var delegetorIncentive types.DelegetorIncentive
		if err := k.cdc.Unmarshal(value, &delegetorIncentive); err != nil {
			return err
		}

		delegetorIncentives = append(delegetorIncentives, delegetorIncentive)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDelegetorIncentiveResponse{DelegetorIncentive: delegetorIncentives, Pagination: pageRes}, nil
}

func (k Keeper) DelegetorIncentive(goCtx context.Context, req *types.QueryGetDelegetorIncentiveRequest) (*types.QueryGetDelegetorIncentiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDelegetorIncentive(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDelegetorIncentiveResponse{DelegetorIncentive: val}, nil
}
