package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/resolver/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	registrytypes "github.com/mycel-domain/mycel/x/registry/types"
)

func (k Keeper) WalletRecord(goCtx context.Context, req *types.QueryWalletRecordRequest) (*types.QueryWalletRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate request parameters
	_, err := registrytypes.GetWalletAddressFormat(req.WalletRecordType)
	if err != nil {
		return nil, err
	}

	// Query domain record
	_, err = k.registryKeeper.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.registryKeeper.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}

	return &types.QueryWalletRecordResponse{Value: secondLevelDomain.WalletRecords[req.WalletRecordType]}, nil
}
