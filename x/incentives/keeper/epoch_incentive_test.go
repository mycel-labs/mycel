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

func createNEpochIncentive(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EpochIncentive {
	items := make([]types.EpochIncentive, n)
	for i := range items {
		items[i].Epoch = int64(i)

		keeper.SetEpochIncentive(ctx, items[i])
	}
	return items
}

func TestEpochIncentiveGet(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNEpochIncentive(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEpochIncentive(ctx,
			item.Epoch,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEpochIncentiveRemove(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNEpochIncentive(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEpochIncentive(ctx,
			item.Epoch,
		)
		_, found := keeper.GetEpochIncentive(ctx,
			item.Epoch,
		)
		require.False(t, found)
	}
}

func TestEpochIncentiveGetAll(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNEpochIncentive(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEpochIncentive(ctx)),
	)
}
func (suite *KeeperTestSuite) TestSetEpochIncentivesOnRegistration() {
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
			amount:                   sdk.NewInt(10000000),
			regestrationPeriodInWeek: 12 * 4 * 100,
			expCurrentEpoch:          0,
			expAmount:                sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, sdk.NewInt(10000000))),
			fn: func() {
			},
		},
		{
			amount:                   sdk.NewInt(100),
			regestrationPeriodInWeek: 12,
			expCurrentEpoch:          1,
			expAmount:                sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, sdk.NewInt(100))),
			fn: func() {
			},
		},
		{
			amount:                   sdk.NewInt(100),
			regestrationPeriodInWeek: 12 * 4,
			expCurrentEpoch:          2,
			expAmount:                sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, sdk.NewInt(100))),
			fn: func() {
				// Set epoch incentives at epoch 0
				suite.app.IncentivesKeeper.SetEpochIncentivesOnRegistration(suite.ctx, 12, sdk.NewCoin(registrytypes.MycelDenom, sdk.NewInt(100)))
				suite.ctx = suite.ctx.WithBlockHeight(2).WithBlockTime(now.Add(time.Hour * 24 * 90))
				suite.app.EpochsKeeper.BeginBlocker(suite.ctx)
			},
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			tc.fn()

			// Before incentives
			beforeIncentives := suite.app.IncentivesKeeper.GetAllEpochIncentive(suite.ctx)
			// Sum up all the incentives
			beforeTotalAmount := sdk.NewInt(0)
			for _, incentive := range beforeIncentives {
				beforeTotalAmount = beforeTotalAmount.Add(incentive.Amount.AmountOf(registrytypes.MycelDenom))
			}

			// Set incentives
			suite.app.IncentivesKeeper.SetEpochIncentivesOnRegistration(suite.ctx, tc.regestrationPeriodInWeek, sdk.NewCoin(registrytypes.MycelDenom, tc.amount))

			// Check incentive start epoch
			incentives := suite.app.IncentivesKeeper.GetAllEpochIncentive(suite.ctx)
			afterTotalAmount := sdk.NewInt(0)
			// Sum up all the incentives
			for _, incentive := range incentives {
				afterTotalAmount = incentive.Amount.AmountOf(registrytypes.MycelDenom).Add(afterTotalAmount)
			}

			// Check total incentive amount
			suite.Require().Equal(tc.expAmount, sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, afterTotalAmount.Sub(beforeTotalAmount))))

		})
	}
}
