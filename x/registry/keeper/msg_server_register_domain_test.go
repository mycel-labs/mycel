package keeper_test

import (
	"errors"
	"fmt"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (suite *KeeperTestSuite) TestRegisterSubdomain() {
	testCases := []struct {
		creator                  string
		name                     string
		parent                   string
		registrationPeriodInYear uint64
		domainLevel              string
		domainOwnership          types.DomainOwnership
		expErr                   error
		fn                       func()
	}{
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "cel",
			registrationPeriodInYear: 1,
			domainLevel:              "2",
			domainOwnership: types.DomainOwnership{
				Owner:   testutil.Alice,
				Domains: []*types.OwnedDomain{{Name: "foo", Parent: "cel"}},
			},
			expErr: nil,
			fn:     func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "cel",
			registrationPeriodInYear: 4,
			domainLevel:              "2",
			domainOwnership: types.DomainOwnership{
				Owner:   testutil.Alice,
				Domains: []*types.OwnedDomain{{Name: "foo", Parent: "cel"}},
			},
			expErr: nil,
			fn:     func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "cel",
			registrationPeriodInYear: 1,
			domainLevel:              "2",
			expErr:                   sdkerrors.Wrapf(errors.New(fmt.Sprintf("foo.cel")), types.ErrDomainIsAlreadyTaken.Error()),
			fn: func() {
				// Register domain once
				domain := &types.MsgRegisterDomain{
					Creator:                  testutil.Alice,
					Name:                     "foo",
					Parent:                   "cel",
					RegistrationPeriodInYear: 1,
				}
				_, err := suite.msgServer.RegisterDomain(suite.ctx, domain)
				suite.Require().Nil(err)
			},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			registerMsg := &types.MsgRegisterDomain{
				Creator:                  tc.creator,
				Name:                     tc.name,
				Parent:                   tc.parent,
				RegistrationPeriodInYear: tc.registrationPeriodInYear,
			}

			domain := &types.Domain{
				Name:   tc.name,
				Parent: tc.parent,
			}
			parentsName, parentsParent := domain.ParseParent()
			parent, found := suite.app.RegistryKeeper.GetDomain(suite.ctx, parentsName, parentsParent)
			suite.Require().True(found)
			beforeSubdomainCount := parent.SubdomainCount

			// Run test case function
			tc.fn()

			// Before incentives
			beforeIncentives := suite.app.IncentivesKeeper.GetAllEpochIncentive(suite.ctx)
			beforeTotalAmount := sdk.NewInt(0)
			for _, incentive := range beforeIncentives {
				beforeTotalAmount = beforeTotalAmount.Add(incentive.Amount.AmountOf(types.MycelDenom))
			}

			// Register domain
			_, err := suite.msgServer.RegisterDomain(suite.ctx, registerMsg)

			if tc.expErr == nil {
				// Evalute domain ownership
				domainOwnership, found := suite.app.RegistryKeeper.GetDomainOwnership(suite.ctx, tc.creator)
				suite.Require().True(found)
				suite.Require().Equal(tc.domainOwnership, domainOwnership)

				// Evalute if domain is registered
				_, found = suite.app.RegistryKeeper.GetDomain(suite.ctx, tc.name, tc.parent)
				suite.Require().True(found)

				// Evalute if parent's subdomainCount is increased
				parent, found = suite.app.RegistryKeeper.GetDomain(suite.ctx, parentsName, parentsParent)
				suite.Require().True(found)
				afterSubdomainCount := parent.SubdomainCount
				suite.Require().Equal(beforeSubdomainCount+1, afterSubdomainCount)

				// Check if the total amount is increased by the fee
				incentives := suite.app.IncentivesKeeper.GetAllEpochIncentive(suite.ctx)
				afterTotalAmount := sdk.NewInt(0)
				for _, incentive := range incentives {
					afterTotalAmount = incentive.Amount.AmountOf(types.MycelDenom).Add(afterTotalAmount)
				}
				expFee, err := parent.SubdomainRegistrationConfig.GetRegistrationFee(tc.name, tc.registrationPeriodInYear)
				suite.Require().Nil(err)
				// Compare the total amount before and after
				suite.Require().Equal(beforeTotalAmount.Add(expFee.Amount), afterTotalAmount)

				// Evalute events
				suite.Require().Nil(err)
				events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
				eventIndex := len(events) - 1
				suite.Require().EqualValues(sdk.StringEvent{
					Type: types.EventTypeRegsterDomain,
					Attributes: []sdk.Attribute{
						{Key: types.AttributeRegisterDomainEventName, Value: tc.name},
						{Key: types.AttributeRegisterDomainEventParent, Value: tc.parent},
						{Key: types.AttributeRegisterDomainEventExpirationDate, Value: events[eventIndex].Attributes[2].Value},
						{Key: types.AttributeRegisterDomainLevel, Value: tc.domainLevel},
					},
				}, events[eventIndex])
			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}

		})
	}

}
