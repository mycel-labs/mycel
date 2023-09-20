package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/testutil"
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

type ExpEvent struct {
	EpochNumber string
}

func (suite *KeeperTestSuite) TestAfterEpochHooks() {
	var (
		now            = time.Now()
		oneDayDuration = time.Hour*24 + time.Second
	)
	testCases := []struct {
		expEpochNumber            string
		expBeforeEpochStartEvents []ExpEvent
		expAfterEpochEndEvents    []ExpEvent
		fn                        func()
	}{
		{
			expBeforeEpochStartEvents: []ExpEvent{
				{
					EpochNumber: "2",
				},
			},
			expAfterEpochEndEvents: []ExpEvent{
				{
					EpochNumber: "1",
				},
			},
			fn: func() {
				// Begin first block
				suite.ctx = suite.ctx.WithBlockHeight(2).WithBlockTime(now.Add(time.Second))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)

				// Check if curent epoch is expected
				epochInfo, found := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, types.DayEpochId)
				suite.Require().True(found)
				suite.Require().Equal(int64(2), epochInfo.CurrentEpoch)
			},
		},
		{
			expBeforeEpochStartEvents: []ExpEvent{
				{
					EpochNumber: "2",
				},
				{
					EpochNumber: "3",
				},
			},
			expAfterEpochEndEvents: []ExpEvent{
				{
					EpochNumber: "1",
				},
{
					EpochNumber: "2",
				},
			},
			fn: func() {
				// Begin first block
				suite.ctx = suite.ctx.WithBlockHeight(2).WithBlockTime(now.Add(time.Second))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)

				// Check if curent epoch is expected
				epochInfo, found := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, types.DayEpochId)
				suite.Require().True(found)
				suite.Require().Equal(int64(2), epochInfo.CurrentEpoch)

				// Begin second block
				suite.ctx = suite.ctx.WithBlockHeight(3).WithBlockTime(now.Add(oneDayDuration))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)

				// Check if curent epoch is expected
				epochInfo, found = suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, types.DayEpochId)
				suite.Require().True(found)
				suite.Require().Equal(int64(3), epochInfo.CurrentEpoch)
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

			// Check before epoch start events
			if len(tc.expBeforeEpochStartEvents) != 0 {
				beforeEpochStartEvents, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), BeforeEpochStartEventType)
				suite.Require().True(found)
				for i, expEvent := range tc.expBeforeEpochStartEvents {
					event := beforeEpochStartEvents[i]
					suite.Require().Equal(BeforeEpochStartEventType, event.Type)
					suite.Require().Equal(EpochIdentifier, event.Attributes[0].Value)
					suite.Require().Equal(expEvent.EpochNumber, event.Attributes[1].Value)
					suite.Require().True(found)
				}
			}

			if len(tc.expAfterEpochEndEvents) != 0 {
				afterEpochEndEvents, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), AfterEpochEndEventType)
				suite.Require().True(found)
				for i, expEvent := range tc.expAfterEpochEndEvents {
					event := afterEpochEndEvents[i]
					suite.Require().Equal(AfterEpochEndEventType, event.Type)
					suite.Require().Equal(EpochIdentifier, event.Attributes[0].Value)
					suite.Require().Equal(expEvent.EpochNumber, event.Attributes[1].Value)
					suite.Require().True(found)
				}

			}

		})
	}

}
