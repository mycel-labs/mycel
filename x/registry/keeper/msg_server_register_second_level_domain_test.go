package keeper_test

import (
	"errors"
	"fmt"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (suite *KeeperTestSuite) TestRegisterSecondLevelDomain() {
	testCases := []struct {
		creator                  string
		name                     string
		parent                   string
		registrationPeriodInYear uint64
		domainOwnership          types.DomainOwnership
		expErr                   error
		fn                       func()
	}{
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "cel",
			registrationPeriodInYear: 1,
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
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "xxx",
			registrationPeriodInYear: 1,
			expErr:                   sdkerrors.Wrapf(errors.New(fmt.Sprintf("xxx")), types.ErrParentDomainDoesNotExist.Error()),
			fn: func() {
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

			// domain := &types.SecondLevelDomain{
			// 	Name:   tc.name,
			// 	Parent: tc.parent,
			// }
			// parentsName := domain.ParseParent()
			// parent, found := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, parentsName)
			// suite.Require().True(found)
			// beforeSubdomainCount := parent.SubdomainCount

			// Run test case function
			tc.fn()

			// Register domain
			_, err := suite.msgServer.RegisterDomain(suite.ctx, registerMsg)
			fmt.Println("----Case_", i , "---01", err)

			if err == nil {
				// Evalute domain ownership
				domainOwnership, found := suite.app.RegistryKeeper.GetDomainOwnership(suite.ctx, tc.creator)
				suite.Require().True(found)
				suite.Require().Equal(tc.domainOwnership, domainOwnership)

				// Evalute if domain is registered
				_, found = suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx,  tc.parent)
				suite.Require().True(found)

				// // Evalute if parent's subdomainCount is increased
				// parent, found = suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, parentsName)
				// suite.Require().True(found)
				// afterSubdomainCount := parent.SubdomainCount
				// suite.Require().Equal(beforeSubdomainCount+1, afterSubdomainCount)

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
					},
				}, events[eventIndex])
			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}

		})
	}

}
