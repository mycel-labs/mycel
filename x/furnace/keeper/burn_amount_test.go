package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/app/params"
	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/furnace/keeper"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBurnAmount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.BurnAmount {
	items := make([]types.BurnAmount, n)
	for i := range items {
		items[i].Index = uint64(i)

		keeper.SetBurnAmount(ctx, items[i])
	}
	return items
}

func TestBurnAmountGet(t *testing.T) {
	keeper, ctx := keepertest.FurnaceKeeper(t)
	items := createNBurnAmount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBurnAmount(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestBurnAmountRemove(t *testing.T) {
	keeper, ctx := keepertest.FurnaceKeeper(t)
	items := createNBurnAmount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBurnAmount(ctx,
			item.Index,
		)
		_, found := keeper.GetBurnAmount(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestBurnAmountGetAll(t *testing.T) {
	keeper, ctx := keepertest.FurnaceKeeper(t)
	items := createNBurnAmount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBurnAmount(ctx)),
	)
}

func (suite *KeeperTestSuite) TestAddRegistrationFeeToBurnAmounts() {
	testCases := []struct {
		registrationFee         sdk.Coin
		regitrationPeriodInYear uint64
		expStartBurnAmountIndex uint64
		fn                      func()
	}{
		{
			registrationFee:         sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(365)),
			regitrationPeriodInYear: 1,
			expStartBurnAmountIndex: 2,
		},
		{
			registrationFee:         sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(36000000)),
			regitrationPeriodInYear: 2,
			expStartBurnAmountIndex: 3,
			fn: func() {
				for i := int64(1); i <= int64(defaultConfig.DefaultTotalEpochs+1); i++ {
					suite.ctx = suite.ctx.WithBlockHeight(i).WithBlockTime(now.Add(oneDayDuration))
					suite.app.EpochsKeeper.BeginBlocker(suite.ctx)
				}
			},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			if tc.fn != nil {
				tc.fn()
			}

			// Set burn totalBurnAmount
			burnAmounts, err := suite.app.FurnaceKeeper.AddRegistrationFeeToBurnAmounts(suite.ctx, tc.regitrationPeriodInYear, tc.registrationFee)
			suite.Require().NoError(err)

			totalBurnAmount := sdk.NewInt(0)

			for i, burnAmount := range burnAmounts {
				if i == 0 {
					suite.Require().Equal(tc.expStartBurnAmountIndex, burnAmount.Index)
				}

				storedBurnAmount, found := suite.app.FurnaceKeeper.GetBurnAmount(suite.ctx, burnAmount.Index)
				suite.Require().True(found)
				suite.Require().Equal(burnAmount, storedBurnAmount)

				totalBurnAmount = totalBurnAmount.Add(burnAmount.TotalBurnAmount.Amount)
			}
			suite.Require().Equal(tc.registrationFee.Amount, totalBurnAmount)
		})
	}
}
