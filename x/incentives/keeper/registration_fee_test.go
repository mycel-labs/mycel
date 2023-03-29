package keeper_test

import (
	"fmt"
	registrytypes "mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestSetIncentivesOnRegistration() {
	testCases := []struct {
		expCurrentEpoch        int64
		expIncentiveStartEpoch int64
		expIncentiveEndEpoch   int64
		expTotalAmount         sdk.Coins
		fn                     func()
	}{
		{
			expCurrentEpoch:        1,
			expIncentiveStartEpoch: 2,
			expIncentiveEndEpoch:   3,
			expTotalAmount:         sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, sdk.NewInt(100))),
			fn: func() {
				suite.app.IncentivesKeeper.SetIncentivesOnRegistration(suite.ctx, 1, sdk.NewInt(100))
			},
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			tc.fn()

			incentives := suite.app.IncentivesKeeper.GetAllIncentive(suite.ctx)

			var totalAmount sdk.Coins
			for i, incentive := range incentives {
				totalAmount = totalAmount.Add(incentive.Amount...)

				// Check incentive start epoch
				suite.Require().Equal(tc.expCurrentEpoch+int64(i)+1, incentive.Epoch)
			}

			// Check total incentive amount
			suite.Require().Equal(tc.expTotalAmount, totalAmount)

		})
	}
}
