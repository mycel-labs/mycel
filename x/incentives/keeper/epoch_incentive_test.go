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
