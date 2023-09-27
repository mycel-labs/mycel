package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/testutil"
	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
	"github.com/mycel-domain/mycel/x/furnace/types"
	"time"
)

type ExpEvent struct {
	EpochIndex                string
	EpochNumber               string
	BurnIndex                 string
	BurnTotalEpochs           string
	BurnCurrentEpoch          string
	BurnAmount                string
	BurnCumulativeBurntAmount string
}

func (suite *KeeperTestSuite) TestAfterEpochEnd() {
	var (
		now            = time.Now()
		oneDayDuration = time.Hour*24 + time.Second
	)
	testCases := []struct {
		totalBurnAmounts []int64
		expectedEvents   []ExpEvent
		epochsCount      int64
		fn               func()
	}{
		{
			totalBurnAmounts: []int64{30, 31},
			epochsCount:      4,
			expectedEvents: []ExpEvent{
				{
					EpochIndex:                "daily",
					EpochNumber:               "1",
					BurnIndex:                 "1",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "1",
					BurnAmount:                "10stake",
					BurnCumulativeBurntAmount: "10stake",
				},
				{
					EpochIndex:                "daily",
					EpochNumber:               "2",
					BurnIndex:                 "1",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "2",
					BurnAmount:                "10stake",
					BurnCumulativeBurntAmount: "20stake",
				},
				{
					EpochIndex:                "daily",
					EpochNumber:               "3",
					BurnIndex:                 "1",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "3",
					BurnAmount:                "10stake",
					BurnCumulativeBurntAmount: "30stake",
				},
				{
					EpochIndex:                "daily",
					EpochNumber:               "4",
					BurnIndex:                 "2",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "1",
					BurnAmount:                "10stake",
					BurnCumulativeBurntAmount: "10stake",
				},
			},
		},
		{
			totalBurnAmounts: []int64{31},
			epochsCount:      3,
			expectedEvents: []ExpEvent{
				{
					EpochIndex:                "daily",
					EpochNumber:               "1",
					BurnIndex:                 "1",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "1",
					BurnAmount:                "10stake",
					BurnCumulativeBurntAmount: "10stake",
				},
				{
					EpochIndex:                "daily",
					EpochNumber:               "2",
					BurnIndex:                 "1",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "2",
					BurnAmount:                "10stake",
					BurnCumulativeBurntAmount: "20stake",
				},
				{
					EpochIndex:                "daily",
					EpochNumber:               "3",
					BurnIndex:                 "1",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "3",
					BurnAmount:                "11stake",
					BurnCumulativeBurntAmount: "31stake",
				},
			},
		},

		{
			totalBurnAmounts: []int64{1},
			epochsCount:      3,
			expectedEvents: []ExpEvent{
				{
					EpochIndex:                "daily",
					EpochNumber:               "1",
					BurnIndex:                 "1",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "1",
					BurnAmount:                "1stake",
					BurnCumulativeBurntAmount: "1stake",
				},
				{
					EpochIndex:                "daily",
					EpochNumber:               "2",
					BurnIndex:                 "1",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "2",
					BurnAmount:                "0stake",
					BurnCumulativeBurntAmount: "1stake",
				},
				{
					EpochIndex:                "daily",
					EpochNumber:               "3",
					BurnIndex:                 "1",
					BurnTotalEpochs:           "3",
					BurnCurrentEpoch:          "3",
					BurnAmount:                "0stake",
					BurnCumulativeBurntAmount: "1stake",
				},
			},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Set burn totalBurnAmount
			for i, _ := range tc.totalBurnAmounts {
				suite.app.FurnaceKeeper.SetBurnAmount(suite.ctx, types.BurnAmount{
					Index:                 uint64(i + 1),
					TotalEpochs:           3,
					CurrentEpoch:          0,
					TotalBurnAmount:       sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(tc.totalBurnAmounts[i])),
					CumulativeBurntAmount: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0)),
				})
			}

			// Run fn
			if tc.fn != nil {
				tc.fn()
			}

			for i := int64(1); i <= tc.epochsCount; i++ {
				suite.ctx = suite.ctx.WithBlockHeight(i + 1).WithBlockTime(now.Add(oneDayDuration))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)
				// Check if curent epoch is expected
				epochInfo, found := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, epochstypes.DailyEpochId)
				suite.Require().True(found)
				suite.Require().Equal(i+1, epochInfo.CurrentEpoch)

				// Check if burn amount is expected
				config, found := suite.app.FurnaceKeeper.GetEpochBurnConfig(suite.ctx)
				suite.Require().True(found)
				_, found = suite.app.FurnaceKeeper.GetBurnAmount(suite.ctx, uint64(config.CurrentBurnAmountIndex))
				suite.Require().True(found)
			}

			// TODO: check if token is burnt

			// Check if event is emitted
			events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeEpochBurn)
			suite.Require().True(found)
			suite.Require().Equal(len(tc.expectedEvents), len(events))
			for i, event := range events {
				suite.Require().Equal(tc.expectedEvents[i].EpochIndex, event.Attributes[0].Value)
				suite.Require().Equal(tc.expectedEvents[i].EpochNumber, event.Attributes[1].Value)
				suite.Require().Equal(tc.expectedEvents[i].BurnIndex, event.Attributes[2].Value)
				suite.Require().Equal(tc.expectedEvents[i].BurnTotalEpochs, event.Attributes[3].Value)
				suite.Require().Equal(tc.expectedEvents[i].BurnCurrentEpoch, event.Attributes[4].Value)
				suite.Require().Equal(tc.expectedEvents[i].BurnAmount, event.Attributes[5].Value)
				suite.Require().Equal(tc.expectedEvents[i].BurnCumulativeBurntAmount, event.Attributes[6].Value)
			}
		})
	}

}
