package keeper_test

import (
	"fmt"

	"github.com/mycel-domain/mycel/x/registry/types"

	"github.com/mycel-domain/mycel/testutil"

	errorsmod "cosmossdk.io/errors"
)

func (suite *KeeperTestSuite) TestUpdateWalletRecord() {
	testCases := []struct {
		creator          string
		name             string
		parent           string
		walletRecordType string
		value            string
		expErr           error
		fn               func()
	}{
		{
			creator:          testutil.Alice,
			name:             "foo",
			parent:           "cel",
			walletRecordType: "ETHEREUM_MAINNET_MAINNET",
			value:            "0x1234567890123456789012345678901234567890",
			expErr:           nil,
			fn:               func() {},
		},
		{
			creator:          testutil.Alice,
			name:             "foo",
			parent:           "cel",
			walletRecordType: "ETHEREUM_MAINNET_MAINNET",
			value:            "0x1234567890123456789012345678901234567891",
			expErr:           nil,
			fn:               func() {},
		},
		{
			creator:          testutil.Alice,
			name:             "foo",
			parent:           "cel",
			walletRecordType: "ETHEREUM_TESTNET_GOERLI",
			value:            "0x1234567890123456789012345678901234567890",
			expErr:           nil,
			fn:               func() {},
		},
		{
			creator:          testutil.Alice,
			name:             "foo",
			parent:           "cel",
			walletRecordType: "POLYGON_MAINNET_MAINNET",
			value:            "0x1234567890123456789012345678901234567890",
			expErr:           nil,
			fn:               func() {},
		},
		{
			creator:          testutil.Alice,
			name:             "foo",
			parent:           "cel",
			walletRecordType: "POLYGON_TESTNET_MUMBAI",
			value:            "0x1234567890123456789012345678901234567890",
			expErr:           nil,
			fn:               func() {},
		},
		{
			creator:          testutil.Alice,
			name:             "hoge",
			parent:           "fuga",
			walletRecordType: "ETHEREUM_MAINNET_MAINNET",
			value:            "0x1234567890123456789012345678901234567890",
			expErr:           errorsmod.Wrapf(types.ErrSecondLevelDomainNotFound, "hoge.fuga"),
			fn:               func() {},
		},
		{
			creator:          testutil.Bob,
			name:             "foo",
			parent:           "cel",
			walletRecordType: "ETHEREUM_MAINNET_MAINNET",
			value:            "0x1234567890123456789012345678901234567890",
			expErr:           errorsmod.Wrapf(types.ErrSecondLevelDomainNotEditable, "%s", testutil.Bob),
			fn:               func() {},
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

			// Update wallet record
			msgUpdateRecord := &types.MsgUpdateWalletRecord{
				Creator:          tc.creator,
				Name:             tc.name,
				Parent:           tc.parent,
				WalletRecordType: tc.walletRecordType,
				Value:            tc.value,
			}
			_, err = suite.msgServer.UpdateWalletRecord(suite.ctx, msgUpdateRecord)

			if tc.expErr == nil {
				// Check if the record is updated
				suite.Require().Nil(err)
				res, _ := suite.app.RegistryKeeper.GetSecondLevelDomain(suite.ctx, domain.Name, domain.Parent)
				suite.Require().Equal(tc.value, res.Records[tc.walletRecordType].GetWalletRecord().GetValue())

				// Evalute events
				events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeUpdateWalletRecord)
				suite.Require().True(found)

				for _, event := range events {
					suite.Require().Equal(tc.name, event.Attributes[0].Value)
					suite.Require().Equal(tc.parent, event.Attributes[1].Value)
					suite.Require().Equal(tc.walletRecordType, event.Attributes[2].Value)
					suite.Require().Equal(tc.value, event.Attributes[3].Value)
				}
			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}

		})
	}

}
