package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/mycel-domain/mycel/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DomainAll(goCtx context.Context, req *types.QueryAllDomainRequest) (*types.QueryAllDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var domains []types.Domain
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	domainStore := prefix.NewStore(store, types.KeyPrefix(types.DomainKeyPrefix))

	pageRes, err := query.Paginate(domainStore, req.Pagination, func(key []byte, value []byte) error {
		var domain types.Domain
		if err := k.cdc.Unmarshal(value, &domain); err != nil {
			return err
		}

		domains = append(domains, domain)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDomainResponse{Domain: domains, Pagination: pageRes}, nil
}

func (k Keeper) Domain(goCtx context.Context, req *types.QueryGetDomainRequest) (*types.QueryGetDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDomain(
		ctx,
		req.Name,
		req.Parent,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDomainResponse{Domain: val}, nil
}
