package keeper_test

import (
	"fmt"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestSetRegistrationFees() {
	fee1, err := sdk.ParseCoinNormalized("100mycel")
	fee2, err := sdk.ParseCoinNormalized("200mycel")
	fee3, err := sdk.ParseCoinNormalized("300mycel")
	suite.Require().Nil(err)

	testCases := []struct {
		creator      string
		domain       string
		feesByName   []types.ReqRegistrationFeeByName
		feesByLength []types.ReqRegistrationFeeByLength
		defaultFee   sdk.Coin
		expErr       error
		fn           func()
	}{
		{
			creator: testutil.Alice,
			domain:  "foo",
			feesByName: []types.ReqRegistrationFeeByName{
				{Name: "aaa", IsRegistrable: true, Fee: &fee1},
				{Name: "bbb", IsRegistrable: true, Fee: &fee2},
				{Name: "ccc", IsRegistrable: true, Fee: &fee3},
			},
			feesByLength: []types.ReqRegistrationFeeByLength{
				{Length: 1, IsRegistrable: true, Fee: &fee1},
				{Length: 2, IsRegistrable: true, Fee: &fee2},
				{Length: 3, IsRegistrable: true, Fee: &fee3},
			},
			defaultFee: fee1,
			expErr:     nil,
			fn:         func() {},
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Register domain
			domain := &types.MsgRegisterTopLevelDomain{
				Creator:                  testutil.Alice,
				Name:                     "foo",
				RegistrationPeriodInYear: 1,
			}
			_, err := suite.msgServer.RegisterTopLevelDomain(suite.ctx, domain)
			suite.Require().Nil(err)
			// Run test case function
			tc.fn()

			// set registration fees
			msgSetRegistrationFees := &types.MsgSetRegistrationFees{
				Creator:      tc.creator,
				Domain:       tc.domain,
				FeesByName:   tc.feesByName,
				FeesByLength: tc.feesByLength,
				DefaultFee:   tc.defaultFee,
			}
			_, err = suite.msgServer.SetRegistrationFees(suite.ctx, msgSetRegistrationFees)

			if tc.expErr == nil {
				// Evalute events
				suite.Require().Nil(err)
				res, _ := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, domain.Name)
				resFees := res.SubdomainConfig.SubdomainRegistrationFees

				// Evaluate stored values
				for _, fee := range tc.feesByName {
					suite.Require().Equal(resFees.FeeByName[fee.Name].Fee, fee.Fee)
				}
				for _, fee := range tc.feesByLength {
					suite.Require().Equal(resFees.FeeByLength[fee.Length].Fee, fee.Fee)
				}
				suite.Require().Equal(resFees.DefaultFee, tc.defaultFee)

				// Event check
				events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
				eventIndex := len(events) - 1
				suite.Require().EqualValues(sdk.StringEvent{
					Type: types.EventTypeSetRegistrationFees,
					Attributes: []sdk.Attribute{
						{Key: types.AttributeSetRegistrationFeesDomain, Value: tc.domain},
					},
				},
					events[eventIndex])

			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}

		})
	}

}
