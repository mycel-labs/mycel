package keeper_test

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (suite *KeeperTestSuite) TestQueryDomainRegistrationFee() {
	config := types.GetDefaultSubdomainConfig(1)
	tcs := []struct {
		desc   string
		req    *types.QueryDomainRegistrationFeeRequest
		resp   *types.QueryDomainRegistrationFeeResponse
		expErr error
		fn     func()
	}{
		{
			desc: "Should register top-level-domain",
			req: &types.QueryDomainRegistrationFeeRequest{
				Name:                     "first",
				Parent:                   "",
				RegistrationPeriodInYear: 1,
				Registerer:               testutil.Alice,
			},
			resp: &types.QueryDomainRegistrationFeeResponse{
				IsRegistrable:             true,
				RegistrationPeriodInYear:  1,
				MaxSubDomainRegistrations: config.MaxSubdomainRegistrations,
				ErrorMessage:              "",
			},
			expErr: nil,
			fn:     func() {},
		},
		{
			desc: "Should register second-level-domain",
			req: &types.QueryDomainRegistrationFeeRequest{
				Name:                     "first",
				Parent:                   "cel",
				RegistrationPeriodInYear: 1,
				Registerer:               testutil.Alice,
			},
			resp: &types.QueryDomainRegistrationFeeResponse{
				IsRegistrable:             true,
				RegistrationPeriodInYear:  1,
				MaxSubDomainRegistrations: 0,
				ErrorMessage:              "",
			},
			expErr: nil,
			fn:     func() {},
		},
		{
			desc: "Should not register second-level-domain because of registration policy",
			req: &types.QueryDomainRegistrationFeeRequest{
				Name:                     "first",
				Parent:                   "private",
				RegistrationPeriodInYear: 1,
				Registerer:               testutil.Bob,
			},
			resp: &types.QueryDomainRegistrationFeeResponse{
				IsRegistrable:             false,
				RegistrationPeriodInYear:  0,
				MaxSubDomainRegistrations: 0,
				ErrorMessage:              errorsmod.Wrapf(types.ErrNotAllowedRegisterDomain, "private").Error(),
			},
			expErr: nil,
			fn: func() {
				_, _, err := suite.app.RegistryKeeper.RegisterTopLevelDomain(suite.ctx, testutil.Alice, "private", 1)
				suite.Require().Nil(err)
			},
		},
	}
	for _, tc := range tcs {
		suite.Run(tc.desc, func() {
			// Setup
			suite.SetupTest()
			tc.fn()

			// Calculate fee
			k := suite.app.RegistryKeeper

			// Get DomainRegistrationFee
			resp, err := k.DomainRegistrationFee(suite.ctx, tc.req)
			suite.Require().Nil(err)

			if tc.expErr == nil {
				suite.Require().Nil(err)
				suite.Require().Equal(tc.resp.IsRegistrable, resp.IsRegistrable)
				suite.Require().Equal(tc.resp.RegistrationPeriodInYear, resp.RegistrationPeriodInYear)
				suite.Require().Equal(tc.resp.MaxSubDomainRegistrations, resp.MaxSubDomainRegistrations)
				// suite.Require().Equal(tc.resp.Fee, resp.Fee)
				// suite.Require().Equal(tc.resp, resp)
			} else {
				suite.Require().NotNil(err)
				suite.Require().Equal(tc.expErr.Error(), err.Error())
			}
		})
	}
}
