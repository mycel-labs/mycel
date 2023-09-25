package keeper_test

import (
	"fmt"
	"github.com/mycel-domain/mycel/testutil"
	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
	"github.com/mycel-domain/mycel/x/furnace/types"
	"time"
)

type ExpEvent struct {
	EpochIdentifier string
	EpochNumber     string
	BurntAmount     string
}

func (suite *KeeperTestSuite) TestAfterEpochEnd() {
	var (
		now = time.Now()
	)
	testCases := []struct {
		totalBurnAmount int64
		identifier      uint64
		expectedEvent   ExpEvent
		fn              func()
	}{
		{
			totalBurnAmount: 100,
			identifier:      1,
			expectedEvent: ExpEvent{
				EpochIdentifier: "1",
				EpochNumber:     "1",
				BurntAmount:     "100",
			},
			fn: func() {
				// Begin first block
				suite.ctx = suite.ctx.WithBlockHeight(2).WithBlockTime(now.Add(time.Second))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)

				// Check if curent epoch is expected
				epochInfo, found := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, epochstypes.DailyEpochId)
				suite.Require().True(found)
				suite.Require().Equal(int64(2), epochInfo.CurrentEpoch)

			},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			tc.fn()

			// TODO: check if token is burnt

			// Check if event is emitted
			events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeEpochBurn)
			suite.Require().True(found)
			event := events[0]
			suite.Require().Equal(tc.expectedEvent.EpochIdentifier, event.Attributes[0].Value)
			suite.Require().Equal(tc.expectedEvent.EpochNumber, event.Attributes[1].Value)
			suite.Require().Equal(tc.expectedEvent.BurntAmount, event.Attributes[2].Value)
		})
	}

}
