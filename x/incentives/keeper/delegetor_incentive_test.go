package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "mycel/testutil/keeper"
	"mycel/testutil/nullify"
	"mycel/x/incentives/keeper"
	"mycel/x/incentives/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDelegetorIncentive(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DelegetorIncentive {
	items := make([]types.DelegetorIncentive, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetDelegetorIncentive(ctx, items[i])
	}
	return items
}

func TestDelegetorIncentiveGet(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNDelegetorIncentive(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDelegetorIncentive(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDelegetorIncentiveRemove(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNDelegetorIncentive(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDelegetorIncentive(ctx,
			item.Address,
		)
		_, found := keeper.GetDelegetorIncentive(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestDelegetorIncentiveGetAll(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNDelegetorIncentive(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDelegetorIncentive(ctx)),
	)
}
