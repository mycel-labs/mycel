package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "mycel/testutil/keeper"
	"mycel/testutil/nullify"
	"mycel/x/epochs/keeper"
	"mycel/x/epochs/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNEpochInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EpochInfo {
	items := make([]types.EpochInfo, n)
	for i := range items {
		items[i].Identifier = strconv.Itoa(i)

		keeper.SetEpochInfo(ctx, items[i])
	}
	return items
}

func TestEpochInfoGet(t *testing.T) {
	keeper, ctx := keepertest.EpochsKeeper(t)
	items := createNEpochInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEpochInfo(ctx,
			item.Identifier,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEpochInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.EpochsKeeper(t)
	items := createNEpochInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEpochInfo(ctx,
			item.Identifier,
		)
		_, found := keeper.GetEpochInfo(ctx,
			item.Identifier,
		)
		require.False(t, found)
	}
}

func TestEpochInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.EpochsKeeper(t)
	items := createNEpochInfo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEpochInfo(ctx)),
	)
}
