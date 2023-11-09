package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (k Keeper) DomainRegistrationFee(goCtx context.Context, req *types.QueryDomainRegistrationFeeRequest) (*types.QueryDomainRegistrationFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
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
		} else {
			return &types.QueryDomainRegistrationFeeResponse{
				IsRegistrable:             true,
				Fee:                       fee.TotalFee,
				RegistrationPeriodInYear:  1,
				MaxSubDomainRegistrations: config.MaxSubdomainRegistrations,
				ErrorMessage:              "",
			}, nil
		}
	} else {
		// Second level domain
		domain := types.SecondLevelDomain{Name: req.Name, Parent: req.Parent}
		err := k.ValidateSecondLevelDomainIsRegistrable(ctx, domain)
		if err != nil {
			return createErrorResponse(err), nil
		}
		config := k.GetSecondLevelDomainParentsSubdomainConfig(ctx, domain)
		fee, err := config.GetRegistrationFee(domain.Name, req.RegistrationPeriodInYear)
		if err != nil {
			return createErrorResponse(err), nil
		} else {
			return &types.QueryDomainRegistrationFeeResponse{
				IsRegistrable:             true,
				Fee:                       sdk.NewCoins(fee),
				RegistrationPeriodInYear:  1,
				MaxSubDomainRegistrations: 0,
				ErrorMessage:              "",
			}, nil
		}
	}
}
