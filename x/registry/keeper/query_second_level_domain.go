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

func (k Keeper) SecondLevelDomainAll(ctx context.Context, req *types.QueryAllSecondLevelDomainRequest) (*types.QueryAllSecondLevelDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var secondLevelDomains []types.SecondLevelDomainResponse

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SecondLevelDomainKeyPrefix))

	secondLevelDomainStore := prefix.NewStore(store, types.KeyPrefix(types.SecondLevelDomainKeyPrefix))

	pageRes, err := query.Paginate(secondLevelDomainStore, req.Pagination, func(key []byte, value []byte) error {
		var secondLevelDomain types.SecondLevelDomain
		if err := k.cdc.Unmarshal(value, &secondLevelDomain); err != nil {
			return err
		}
		secondLevelDomainResponse := types.SecondLevelDomainResponse{
			Name:           secondLevelDomain.Name,
			Parent:         secondLevelDomain.Parent,
			ExpirationDate: secondLevelDomain.ExpirationDate,
		}

		secondLevelDomains = append(secondLevelDomains, secondLevelDomainResponse)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSecondLevelDomainResponse{SecondLevelDomain: secondLevelDomains, Pagination: pageRes}, nil
}

func (k Keeper) SecondLevelDomain(ctx context.Context, req *types.QueryGetSecondLevelDomainRequest) (*types.QueryGetSecondLevelDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetSecondLevelDomain(
		ctx,
		req.Name,
		req.Parent,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	secondLevelDomainResponse := types.SecondLevelDomainResponse{
		Name:           val.Name,
		Parent:         val.Parent,
		ExpirationDate: val.ExpirationDate,
	}

	return &types.QueryGetSecondLevelDomainResponse{SecondLevelDomain: secondLevelDomainResponse}, nil
}
