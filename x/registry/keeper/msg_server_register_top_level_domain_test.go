package keeper_test

import (
	"errors"
	"fmt"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (suite *KeeperTestSuite) TestRegisterTopLevelDomain() {
	testCases := []struct {
		creator                  string
		name                     string
		registrationPeriodInYear uint64
		domainOwnership          types.DomainOwnership
		expErr                   error
		fn                       func()
	}{
		{
			creator:                  testutil.Alice,
			name:                     "cel",
			registrationPeriodInYear: 1,
			expErr:                   nil,
			fn:                       func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "cel",
			registrationPeriodInYear: 4,
			expErr:                   nil,
			fn:                       func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "cel",
			registrationPeriodInYear: 1,
			expErr:                   sdkerrors.Wrapf(errors.New(fmt.Sprintf("foo.cel")), types.ErrDomainIsAlreadyTaken.Error()),
			fn: func() {
				// Register domain once
				domain := &types.MsgRegisterTopLevelDomain{
					Creator:                  testutil.Alice,
					Name:                     "cel",
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

			if err == nil {
				// Evalute if domain is registered
				_, found := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, tc.name)
				suite.Require().True(found)

				// Evalute events
				suite.Require().Nil(err)
				events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
				eventIndex := len(events) - 1
				suite.Require().EqualValues(sdk.StringEvent{
					Type: types.EventTypeRegsterDomain,
					Attributes: []sdk.Attribute{
						{Key: types.AttributeRegisterTopLevelDomainEventName, Value: tc.name},
						{Key: types.AttributeRegisterTopLevelDomainEventExpirationDate, Value: events[eventIndex].Attributes[2].Value},
					},
				}, events[eventIndex])
			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}
		})
	}
}
