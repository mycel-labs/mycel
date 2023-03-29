package keeper_test

import (
	"strconv"
	"testing"

	keepertest "mycel/testutil/keeper"
	"mycel/testutil/nullify"
	"mycel/x/incentives/keeper"
	"mycel/x/incentives/types"

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
