package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/mycel-domain/mycel/x/incentives/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IncentiveAll(goCtx context.Context, req *types.QueryAllIncentiveRequest) (*types.QueryAllIncentiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var incentives []types.Incentive
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	incentiveStore := prefix.NewStore(store, types.KeyPrefix(types.IncentiveKeyPrefix))

	pageRes, err := query.Paginate(incentiveStore, req.Pagination, func(key []byte, value []byte) error {
		var incentive types.Incentive
		if err := k.cdc.Unmarshal(value, &incentive); err != nil {
			return err
		}

		incentives = append(incentives, incentive)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIncentiveResponse{Incentive: incentives, Pagination: pageRes}, nil
}

func (k Keeper) Incentive(goCtx context.Context, req *types.QueryGetIncentiveRequest) (*types.QueryGetIncentiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetIncentive(
		ctx,
		req.Epoch,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetIncentiveResponse{Incentive: val}, nil
}
