package keeper_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/incentives/keeper"
	"github.com/mycel-domain/mycel/x/incentives/types"
	registrytypes "github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNIncentive(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Incentive {
	items := make([]types.Incentive, n)
	for i := range items {
		items[i].Epoch = int64(i)

		keeper.SetIncentive(ctx, items[i])
	}
	return items
}

func TestIncentiveGet(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNIncentive(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetIncentive(ctx,
			item.Epoch,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestIncentiveRemove(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNIncentive(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIncentive(ctx,
			item.Epoch,
		)
		_, found := keeper.GetIncentive(ctx,
			item.Epoch,
		)
		require.False(t, found)
	}
}

func TestIncentiveGetAll(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNIncentive(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllIncentive(ctx)),
	)
}

func (suite *KeeperTestSuite) TestSetIncentivesOnRegistration() {
	now := time.Now()
	testCases := []struct {
		amount                   sdk.Int
		regestrationPeriodInWeek uint
		expCurrentEpoch          int64
		expAmount                sdk.Coins
		fn                       func()
	}{
		{
			amount:                   sdk.NewInt(100),
			regestrationPeriodInWeek: 12,
			expCurrentEpoch:          0,
			expAmount:                sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, sdk.NewInt(100))),
			fn: func() {
			},
		},
		{
			amount:                   sdk.NewInt(100),
			regestrationPeriodInWeek: 12,
			expCurrentEpoch:          1,
			expAmount:                sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, sdk.NewInt(100))),
			fn: func() {
				suite.ctx = suite.ctx.WithBlockHeight(2).WithBlockTime(now.Add(time.Hour * 24 * 7))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)
			},
		},
		{
			amount:                   sdk.NewInt(100),
			regestrationPeriodInWeek: 12,
			expCurrentEpoch:          2,
			expAmount:                sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, sdk.NewInt(100))),
			fn: func() {
				suite.ctx = suite.ctx.WithBlockHeight(2).WithBlockTime(now.Add(time.Hour * 24 * 7))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)
				suite.ctx = suite.ctx.WithBlockHeight(2).WithBlockTime(now.Add(time.Hour * 24 * (7 + 1)))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)
			},
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			tc.fn()

			// Before incentives
			beforeIncentives := suite.app.IncentivesKeeper.GetAllIncentive(suite.ctx)
			beforeTotalAmount := sdk.NewInt(0)
			for _, incentive := range beforeIncentives {
				beforeTotalAmount = beforeTotalAmount.Add(incentive.Amount.AmountOf(registrytypes.MycelDenom))
			}

			// Set incentives
			suite.app.IncentivesKeeper.SetIncentivesOnRegistration(suite.ctx, tc.regestrationPeriodInWeek, sdk.NewCoin(registrytypes.MycelDenom, tc.amount))

			// Check incentive start epoch
			incentives := suite.app.IncentivesKeeper.GetAllIncentive(suite.ctx)
			afterTotalAmount := sdk.NewInt(0)
			for i, incentive := range incentives {
				afterTotalAmount = incentive.Amount.AmountOf(registrytypes.MycelDenom).Add(afterTotalAmount)
				suite.Require().Equal(tc.expCurrentEpoch+int64(i)+1, incentive.Epoch)
			}

			// Check total incentive amount
			suite.Require().Equal(tc.expAmount, sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, afterTotalAmount.Sub(beforeTotalAmount))))

		})
	}
}
