package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/furnace/keeper"
	"github.com/mycel-domain/mycel/x/furnace/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBurnAmount(keeper keeper.Keeper, ctx context.Context, n int) []types.BurnAmount {
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
