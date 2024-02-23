package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k Keeper) AllRecords(goCtx context.Context, req *types.QueryAllRecordsRequest) (*types.QueryAllRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Query domain record
	_, err := k.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}

	// Convert repeated Record to map
	values := make(map[string]*types.Record)
	for _, record := range secondLevelDomain.Records {
		key := generateRecordKey(record)
		if key != "" {
			values[key] = record
		}
	}
	return &types.QueryAllRecordsResponse{Values: values}, nil
}

func generateRecordKey(record *types.Record) string {
	switch {
	case record.GetDnsRecord() != nil:
		return record.GetDnsRecord().DnsRecordType.String()
	case record.GetWalletRecord() != nil:
		return record.GetWalletRecord().WalletRecordType.String()
	case record.GetTextRecord() != nil:
		return record.GetTextRecord().Key
	default:
		return ""
	}
}

func (k Keeper) DnsRecord(goCtx context.Context, req *types.QueryDnsRecordRequest) (*types.QueryDnsRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate request parameters
	err := types.ValidateDnsRecordType(req.DnsRecordType)
	if err != nil {
		return nil, err
	}

	// Query domain record
	_, err = k.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}

	value := secondLevelDomain.GetDnsRecord(req.DnsRecordType)
	recordType := types.DnsRecordType(types.DnsRecordType_value[req.DnsRecordType])

	return &types.QueryDnsRecordResponse{
		Value: &types.DnsRecord{DnsRecordType: recordType, Value: value},
	}, nil
}

func (k Keeper) TextRecord(goCtx context.Context, req *types.QueryTextRecordRequest) (*types.QueryTextRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate request parameters
	err := types.ValidateTextRecordKey(req.Key)
	if err != nil {
		return nil, err
	}

	// Query domain record
	_, err = k.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}

	value := secondLevelDomain.GetTextRecord(req.Key)

	return &types.QueryTextRecordResponse{
		Value: &types.TextRecord{Key: req.Key, Value: value},
	}, nil
}

func (k Keeper) WalletRecord(goCtx context.Context, req *types.QueryWalletRecordRequest) (*types.QueryWalletRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate request parameters
	_, err := types.GetWalletAddressFormat(req.WalletRecordType)
	if err != nil {
		return nil, err
	}

	// Query domain record
	_, err = k.GetValidTopLevelDomain(ctx, req.DomainParent)
	if err != nil {
		return nil, err
	}
	secondLevelDomain, err := k.GetValidSecondLevelDomain(ctx, req.DomainName, req.DomainParent)
	if err != nil {
		return nil, err
	}
	value := secondLevelDomain.GetWalletRecord(req.WalletRecordType)
	recordType := types.NetworkName(types.NetworkName_value[req.WalletRecordType])

	return &types.QueryWalletRecordResponse{
		Value: &types.WalletRecord{WalletRecordType: recordType, Value: value},
	}, nil
}
