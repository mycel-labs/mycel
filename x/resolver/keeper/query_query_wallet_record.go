package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/resolver/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	registrytypes "github.com/mycel-domain/mycel/x/registry/types"
)

func (k Keeper) QueryWalletRecord(goCtx context.Context, req *types.QueryQueryWalletRecordRequest) (*types.QueryQueryWalletRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate request parameters
	walletAddressFormat, err := registrytypes.GetWalletAddressFormat(req.NetworkName)
	if err != nil {
		return nil, err
	}
	_ = walletAddressFormat

	// TODO: Query domain QueryWalletRecord
	_, err = k.registryKeeper.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.registryKeeper.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}

	return &types.QueryQueryWalletRecordResponse{Value: secondLevelDomain.WalletRecords[req.NetworkName]}, nil
}
