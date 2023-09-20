package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/epochs/types"
	"strconv"
	"time"
)

type MockHooks struct{}

const (
	EpochIdentifier           = types.DayEpochId
	BeforeEpochStartEventType = "BeforeEpochStart"
	AfterEpochEndEventType    = "AfterEpochEnd"
)

func (h *MockHooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	if epochIdentifier == EpochIdentifier {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				AfterEpochEndEventType,
				sdk.NewAttribute("epochIdentifier", epochIdentifier),
				sdk.NewAttribute("epochNumber", strconv.FormatInt(epochNumber, 10)),
			),
		)
	}
}

func (h *MockHooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	if epochIdentifier == EpochIdentifier {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				BeforeEpochStartEventType,
				sdk.NewAttribute("epochIdentifier", epochIdentifier),
				sdk.NewAttribute("epochNumber", strconv.FormatInt(epochNumber, 10)),
			),
		)
	}
}

func (suite *KeeperTestSuite) TestAfterEpochHooks() {
	var (
		now = time.Now()
		// oneDayDuration = time.Hour * 24
	)
	testCases := []struct {
		expEpochNumber string
		expEventIndex  int
		expEventType   string

		fn func()
	}{
		{
			expEpochNumber: "1",
			expEventIndex:  1,
			expEventType:   BeforeEpochStartEventType,
			fn: func() {
				// Begin first block
				suite.ctx = suite.ctx.WithBlockHeight(2).WithBlockTime(now.Add(time.Second))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)

				// Check if curent epoch is expected
				epochInfo, found := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, types.DayEpochId)
				suite.Require().True(found)
				suite.Require().Equal(int64(1), epochInfo.CurrentEpoch)
			},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Remove hooks
			suite.app.EpochsKeeper.RemoveHooks()

			_ = tc

			// Register hooks
			hook := new(MockHooks)
			suite.app.EpochsKeeper.SetHooks(hook)

			// Run test Case
			tc.fn()

			// Check events
			event := suite.ctx.EventManager().Events()[tc.expEventIndex]
			suite.Require().Equal(tc.expEventType, event.Type)
			suite.Require().Equal(EpochIdentifier, event.Attributes[0].Value)
			suite.Require().Equal(tc.expEpochNumber, event.Attributes[1].Value)
		})
	}

}
