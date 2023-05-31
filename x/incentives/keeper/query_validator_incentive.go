package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/mycel-domain/mycel/x/incentives/types"
)

func (k Keeper) ValidatorIncentiveAll(goCtx context.Context, req *types.QueryAllValidatorIncentiveRequest) (*types.QueryAllValidatorIncentiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validatorIncentives []types.ValidatorIncentive
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	validatorIncentiveStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorIncentiveKeyPrefix))

	pageRes, err := query.Paginate(validatorIncentiveStore, req.Pagination, func(key []byte, value []byte) error {
		var validatorIncentive types.ValidatorIncentive
		if err := k.cdc.Unmarshal(value, &validatorIncentive); err != nil {
			return err
		}

		validatorIncentives = append(validatorIncentives, validatorIncentive)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorIncentiveResponse{ValidatorIncentive: validatorIncentives, Pagination: pageRes}, nil
}

func (k Keeper) ValidatorIncentive(goCtx context.Context, req *types.QueryGetValidatorIncentiveRequest) (*types.QueryGetValidatorIncentiveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetValidatorIncentive(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetValidatorIncentiveResponse{ValidatorIncentive: val}, nil
}
