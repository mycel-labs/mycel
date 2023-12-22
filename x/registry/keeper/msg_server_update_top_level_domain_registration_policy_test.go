package keeper_test

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (suite *KeeperTestSuite) TestUpdateTopLevelDomainRegistrationPolicy() {
	testCases := []struct {
		creator            string
		name               string
		registrationPolicy string
		expErr             error
		fn                 func()
	}{
		{
			creator:            testutil.Alice,
			name:               "tcel",
			registrationPolicy: "PUBLIC",
			expErr:             nil,
			fn:                 func() {},
		},
		{
			creator:            testutil.Alice,
			name:               "tcel",
			registrationPolicy: "PRIVATE",
			expErr:             nil,
			fn:                 func() {},
		},
		{
			creator:            testutil.Alice,
			name:               "tcel",
			registrationPolicy: "INVALID",
			expErr:             errorsmod.Wrapf(types.ErrInvalidRegistrationPolicy, "INVALID"),
			fn:                 func() {},
		},
		{
			creator:            testutil.Alice,
			name:               "invalid",
			registrationPolicy: "PUBLIC",
			expErr:             errorsmod.Wrapf(types.ErrTopLevelDomainNotFound, "invalid"),
			fn:                 func() {},
		},
		{
			creator:            testutil.Bob,
			name:               "tcel",
			registrationPolicy: "PUBLIC",
			expErr:             errorsmod.Wrapf(types.ErrTopLevelDomainNotEditable, "%s", testutil.Bob),
			fn:                 func() {},
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Register top level domain
			registerTopLevelDomainMsg := &types.MsgRegisterTopLevelDomain{
				Creator:                  testutil.Alice,
				Name:                     tc.name,
				RegistrationPeriodInYear: 1,
			}
			if tc.name != "invalid" {
				_, err := suite.msgServer.RegisterTopLevelDomain(suite.ctx, registerTopLevelDomainMsg)
				suite.Require().Nil(err)
			}

			// Run test case function
			tc.fn()

			// Update dns record
			msgUpdateTopLevelDomainRegistrationPolicy := &types.MsgUpdateTopLevelDomainRegistrationPolicy{
				Creator:            tc.creator,
				Name:               tc.name,
				RegistrationPolicy: tc.registrationPolicy,
			}

			_, err := suite.msgServer.UpdateTopLevelDomainRegistrationPolicy(suite.ctx, msgUpdateTopLevelDomainRegistrationPolicy)

			if tc.expErr == nil {
				// Check if the record is updated
				suite.Require().Nil(err)
				res, _ := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, tc.name)
				suite.Require().Equal(tc.registrationPolicy, res.GetRegistrationPolicy().String())
				// Evalute events
				events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeUpdateTopLevelDomainRegistrationPolicy)
				suite.Require().True(found)

				for _, event := range events {
					suite.Require().Equal(tc.name, event.Attributes[0].Value)
					suite.Require().Equal(tc.registrationPolicy, event.Attributes[1].Value)
				}
			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}
		})
	}
}
