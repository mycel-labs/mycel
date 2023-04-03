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

func createNValidatorIncentive(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ValidatorIncentive {
	items := make([]types.ValidatorIncentive, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetValidatorIncentive(ctx, items[i])
	}
	return items
}

func TestValidatorIncentiveGet(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNValidatorIncentive(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidatorIncentive(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestValidatorIncentiveRemove(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNValidatorIncentive(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidatorIncentive(ctx,
			item.Address,
		)
		_, found := keeper.GetValidatorIncentive(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestValidatorIncentiveGetAll(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNValidatorIncentive(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllValidatorIncentive(ctx)),
	)
}
