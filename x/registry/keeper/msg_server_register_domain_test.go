package keeper_test

import (
	"errors"
	"fmt"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (suite *KeeperTestSuite) TestRegisterDomain() {
	testCases := []struct {
		creator                  string
		name                     string
		parent                   string
		registrationPeriodInYear uint64
		domainLevel              string
		expErr                   error
		fn                       func()
	}{
		{
			creator:                  testutil.Alice,
			name:                     "foo",
			parent:                   "cel",
			registrationPeriodInYear: 1,
			domainLevel:              "2",
			expErr:                   nil,
			fn:                       func() {},
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

			domain := &types.MsgRegisterDomain{
				Creator:                  tc.creator,
				Name:                     tc.name,
				Parent:                   tc.parent,
				RegistrationPeriodInYear: tc.registrationPeriodInYear,
			}

			// Run test case function
			tc.fn()

			// Register domain
			_, err := suite.msgServer.RegisterDomain(suite.ctx, domain)

			if tc.expErr == nil {
				// Evalute events
				suite.Require().Nil(err)
				events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
				eventIndex := len(events) - 2
				suite.Require().EqualValues(sdk.StringEvent{
					Type: types.EventTypeRegsterDomain,
					Attributes: []sdk.Attribute{
						{Key: types.AttributeRegisterDomainEventName, Value: tc.name},
						{Key: types.AttributeRegisterDomainEventParent, Value: tc.parent},
						{Key: types.AttributeRegisterDomainEventExpirationDate, Value: events[3].Attributes[2].Value},
						{Key: types.AttributeRegisterDomainLevel, Value: tc.domainLevel},
					},
				}, events[eventIndex])
			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}

		})
	}

}
