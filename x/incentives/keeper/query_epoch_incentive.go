package keeper

import (
	"context"

	"mycel/x/incentives/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EpochIncentiveAll(goCtx context.Context, req *types.QueryAllEpochIncentiveRequest) (*types.QueryAllEpochIncentiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var epochIncentives []types.EpochIncentive
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	epochIncentiveStore := prefix.NewStore(store, types.KeyPrefix(types.EpochIncentiveKeyPrefix))

	pageRes, err := query.Paginate(epochIncentiveStore, req.Pagination, func(key []byte, value []byte) error {
		var epochIncentive types.EpochIncentive
		if err := k.cdc.Unmarshal(value, &epochIncentive); err != nil {
			return err
		}

		epochIncentives = append(epochIncentives, epochIncentive)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEpochIncentiveResponse{EpochIncentive: epochIncentives, Pagination: pageRes}, nil
}

func (k Keeper) EpochIncentive(goCtx context.Context, req *types.QueryGetEpochIncentiveRequest) (*types.QueryGetEpochIncentiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEpochIncentive(
		ctx,
		req.Epoch,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEpochIncentiveResponse{EpochIncentive: val}, nil
}
