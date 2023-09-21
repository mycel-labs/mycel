package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/furnace/keeper"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

func createTestEpochBurnConfig(keeper *keeper.Keeper, ctx sdk.Context) types.EpochBurnConfig {
	item := types.EpochBurnConfig{}
	keeper.SetEpochBurnConfig(ctx, item)
	return item
}

func TestEpochBurnConfigGet(t *testing.T) {
	keeper, ctx := keepertest.FurnaceKeeper(t)
	item := createTestEpochBurnConfig(keeper, ctx)
	rst, found := keeper.GetEpochBurnConfig(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestEpochBurnConfigRemove(t *testing.T) {
	keeper, ctx := keepertest.FurnaceKeeper(t)
	createTestEpochBurnConfig(keeper, ctx)
	keeper.RemoveEpochBurnConfig(ctx)
	_, found := keeper.GetEpochBurnConfig(ctx)
	require.False(t, found)
}
