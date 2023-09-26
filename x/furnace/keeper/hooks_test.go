package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/furnace/types"
	"time"
)

type ExpEvent struct {
	EpochIndex string
	EpochNumber     string
	CumulativeBurntAmount     string
}

func (suite *KeeperTestSuite) TestAfterEpochEnd() {
	var (
		now = time.Now()
	)
	testCases := []struct {
		totalBurnAmount int64
		index      uint64
		expectedEvents  []ExpEvent
		fn              func()
	}{
		{
			totalBurnAmount: 100,
			index:      1,
			expectedEvents: []ExpEvent{
				{
					EpochIndex: "daily",
					EpochNumber:     "1",
					CumulativeBurntAmount:     "0stake",
				},
				{
					EpochIndex: "daily",
					EpochNumber:     "2",
					CumulativeBurntAmount:     "30stake",
				},
			},
			fn: func() {
				// Begin first block
				suite.ctx = suite.ctx.WithBlockHeight(2).WithBlockTime(now.Add(time.Second))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)
				// Begin first block
				suite.ctx = suite.ctx.WithBlockHeight(3).WithBlockTime(now.Add(time.Hour * 24))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)

			},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Set burn totalBurnAmount
			suite.app.FurnaceKeeper.SetBurnAmount(suite.ctx, types.BurnAmount{
				Index:      tc.index,
				TotalBurnAmount: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(tc.totalBurnAmount)),
				CumulativeBurntAmount:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0)),
			})

			tc.fn()

			// TODO: check if token is burnt

			// Check if event is emitted
			events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeEpochBurn)
			suite.Require().True(found)
			for i, event := range events {
				suite.Require().Equal(tc.expectedEvents[i].EpochIndex, event.Attributes[0].Value)
				suite.Require().Equal(tc.expectedEvents[i].EpochNumber, event.Attributes[1].Value)
				suite.Require().Equal(tc.expectedEvents[i].CumulativeBurntAmount, event.Attributes[2].Value)
			}
		})
	}

}
