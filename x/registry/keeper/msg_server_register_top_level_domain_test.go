package keeper_test

import (
	"fmt"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestRegisterTopLevelDomain() {
	testCases := []struct {
		creator                  string
		name                     string
		registrationPeriodInYear uint64
		expErr                   error
		fn                       func()
	}{
		{
			creator:                  testutil.Alice,
			name:                     "cel0",
			registrationPeriodInYear: 1,
			expErr:                   nil,
			fn:                       func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "cel1",
			registrationPeriodInYear: 4,
			expErr:                   nil,
			fn:                       func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "cel2",
			registrationPeriodInYear: 1,
			expErr:                   errorsmod.Wrapf(types.ErrDomainIsAlreadyTaken, "cel2"),
			fn: func() {
				// Register domain once
				domain := &types.MsgRegisterTopLevelDomain{
					Creator:                  testutil.Alice,
					Name:                     "cel2",
					RegistrationPeriodInYear: 1,
				}
				_, err := suite.msgServer.RegisterTopLevelDomain(suite.ctx, domain)
				suite.Require().Nil(err)
			},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			registerMsg := &types.MsgRegisterTopLevelDomain{
				Creator:                  tc.creator,
				Name:                     tc.name,
				RegistrationPeriodInYear: tc.registrationPeriodInYear,
			}

			// Run test case function
			tc.fn()

			// Register domain
			_, err := suite.msgServer.RegisterTopLevelDomain(suite.ctx, registerMsg)
			fmt.Println("----Case_", i, "---01", err)

			if tc.expErr == nil {
				// Evalute if domain is registered
				domain, found := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, tc.name)
				suite.Require().True(found)
				suite.Require().Equal(domain.AccessControl[tc.creator], types.DomainRole_OWNER)

				// Evalute events
				suite.Require().Nil(err)
				events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
				eventIndex := len(events) - 1
				suite.Require().EqualValues(sdk.StringEvent{
					Type: types.EventTypeRegsterTopLevelDomain,
					Attributes: []sdk.Attribute{
						{Key: types.AttributeRegisterTopLevelDomainEventName, Value: tc.name},
						{Key: types.AttributeRegisterTopLevelDomainEventExpirationDate, Value: events[eventIndex].Attributes[1].Value},
					},
				}, events[eventIndex])
			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}
		})
	}
}
