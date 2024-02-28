package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k Keeper) DomainOwnershipAll(ctx context.Context, req *types.QueryAllDomainOwnershipRequest) (*types.QueryAllDomainOwnershipResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var domainOwnerships []types.DomainOwnership

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

	domainOwnershipStore := prefix.NewStore(store, types.KeyPrefix(types.DomainOwnershipKeyPrefix))

	pageRes, err := query.Paginate(domainOwnershipStore, req.Pagination, func(key []byte, value []byte) error {
		var domainOwnership types.DomainOwnership
		if err := k.cdc.Unmarshal(value, &domainOwnership); err != nil {
			return err
		}

		domainOwnerships = append(domainOwnerships, domainOwnership)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDomainOwnershipResponse{DomainOwnership: domainOwnerships, Pagination: pageRes}, nil
}

func (k Keeper) DomainOwnership(ctx context.Context, req *types.QueryGetDomainOwnershipRequest) (*types.QueryGetDomainOwnershipResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetDomainOwnership(
		ctx,
		req.Owner,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDomainOwnershipResponse{DomainOwnership: val}, nil
}
