package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k Keeper) TopLevelDomainAll(goCtx context.Context, req *types.QueryAllTopLevelDomainRequest) (*types.QueryAllTopLevelDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var topLevelDomains []types.TopLevelDomain
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	topLevelDomainStore := prefix.NewStore(store, types.KeyPrefix(types.TopLevelDomainKeyPrefix))

	pageRes, err := query.Paginate(topLevelDomainStore, req.Pagination, func(key []byte, value []byte) error {
		var topLevelDomain types.TopLevelDomain
		if err := k.cdc.Unmarshal(value, &topLevelDomain); err != nil {
			return err
		}

		topLevelDomains = append(topLevelDomains, topLevelDomain)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTopLevelDomainResponse{TopLevelDomain: topLevelDomains, Pagination: pageRes}, nil
}

func (k Keeper) TopLevelDomain(goCtx context.Context, req *types.QueryGetTopLevelDomainRequest) (*types.QueryGetTopLevelDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTopLevelDomain(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTopLevelDomainResponse{TopLevelDomain: val}, nil
}
