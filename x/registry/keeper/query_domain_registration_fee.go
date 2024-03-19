package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func createErrorResponse(err error) *types.QueryDomainRegistrationFeeResponse {
	return &types.QueryDomainRegistrationFeeResponse{
		IsRegistrable:             false,
		Fee:                       sdk.NewCoins(),
		RegistrationPeriodInYear:  0,
		MaxSubDomainRegistrations: 0,
		ErrorMessage:              err.Error(),
	}
}

func (k Keeper) DomainRegistrationFee(ctx context.Context, req *types.QueryDomainRegistrationFeeRequest) (*types.QueryDomainRegistrationFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.Parent == "" {
		// Top level domain
		config := types.GetDefaultSubdomainConfig(1)
		domain := types.TopLevelDomain{
			Name:            req.Name,
			SubdomainConfig: &config,
		}
		err := k.ValidateTopLevelDomainIsRegistrable(ctx, domain)
		if err != nil {
			return createErrorResponse(err), nil
		}
		fee, err := k.GetTopLevelDomainFee(ctx, domain, req.RegistrationPeriodInYear)
		if err != nil {
			return createErrorResponse(err), nil
		}
		return &types.QueryDomainRegistrationFeeResponse{
			IsRegistrable:             true,
			Fee:                       fee.TotalFee,
			RegistrationPeriodInYear:  req.RegistrationPeriodInYear,
			MaxSubDomainRegistrations: config.MaxSubdomainRegistrations,
			ErrorMessage:              "",
		}, nil
	}

	// Second level domain
	domain := types.SecondLevelDomain{Name: req.Name, Parent: req.Parent}
	registerer, err := sdk.AccAddressFromBech32(req.Registerer)
	if err != nil {
		return nil, err
	}
	err = k.ValidateSecondLevelDomainIsRegistrable(ctx, domain, registerer)
	if err != nil {
		return createErrorResponse(err), nil
	}
	config := k.GetSecondLevelDomainParentsSubdomainConfig(ctx, domain)
	fee, err := config.GetRegistrationFee(domain.Name, req.RegistrationPeriodInYear)
	if err != nil {
		return createErrorResponse(err), nil
	}
	return &types.QueryDomainRegistrationFeeResponse{
		IsRegistrable:             true,
		Fee:                       sdk.NewCoins(fee),
		RegistrationPeriodInYear:  req.RegistrationPeriodInYear,
		MaxSubDomainRegistrations: 0,
		ErrorMessage:              "",
	}, nil
}
