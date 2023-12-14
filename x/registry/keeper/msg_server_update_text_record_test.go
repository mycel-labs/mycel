package keeper_test

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (suite *KeeperTestSuite) TestUpdateTextRecord() {
	testCases := []struct {
		creator string
		name    string
		parent  string
		key     string
		value   string
		expErr  error
		fn      func()
	}{
		{
			creator: testutil.Alice,
			name:    "foo",
			parent:  "cel",
			key:     "com.github",
			value:   "mycel-domain",
			expErr:  nil,
			fn:      func() {},
		},
		{
			creator: testutil.Alice,
			name:    "foo",
			parent:  "cel",
			key:     "ETHEREUM_MAINNET_MAINNET",
			value:   "mycel-domain",
			expErr:  errorsmod.Wrapf(types.ErrInvalidTextRecordKey, "%s", "ETHEREUM_MAINNET_MAINNET"),
			fn:      func() {},
		},
		{
			creator: testutil.Alice,
			name:    "hoge",
			parent:  "fuga",
			key:     "com.github",
			value:   "mycel-domain",
			expErr:  errorsmod.Wrapf(types.ErrSecondLevelDomainNotFound, "hoge.fuga"),
			fn:      func() {},
		},
		{
			creator: testutil.Bob,
			name:    "foo",
			parent:  "cel",
			key:     "com.github",
			value:   "mycel-domain",
			expErr:  errorsmod.Wrapf(types.ErrSecondLevelDomainNotEditable, "%s", testutil.Bob),
			fn:      func() {},
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Register domain
			domain := &types.MsgRegisterSecondLevelDomain{
				Creator:                  testutil.Alice,
				Name:                     "foo",
				Parent:                   "cel",
				RegistrationPeriodInYear: 1,
			}
			_, err := suite.msgServer.RegisterSecondLevelDomain(suite.ctx, domain)
			suite.Require().Nil(err)
			// Run test case function
			tc.fn()

			// Update dns record
			msgUpdateRecord := &types.MsgUpdateTextRecord{
				Creator: tc.creator,
				Name:    tc.name,
				Parent:  tc.parent,
				Key:     tc.key,
				Value:   tc.value,
			}
			_, err = suite.msgServer.UpdateTextRecord(suite.ctx, msgUpdateRecord)

			if tc.expErr == nil {
				// Check if the record is updated
				suite.Require().Nil(err)
				res, _ := suite.app.RegistryKeeper.GetSecondLevelDomain(suite.ctx, domain.Name, domain.Parent)
				suite.Require().Equal(tc.value, res.GetTextRecord(tc.key))
				// Evalute events
				events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeUpdateTextRecord)
				suite.Require().True(found)

				for _, event := range events {
					suite.Require().Equal(tc.name, event.Attributes[0].Value)
					suite.Require().Equal(tc.parent, event.Attributes[1].Value)
					suite.Require().Equal(tc.key, event.Attributes[2].Value)
					suite.Require().Equal(tc.value, event.Attributes[3].Value)
				}
			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}

		})
	}

}
