package keeper_test

import (
	"strconv"
	"testing"

	keepertest "mycel/testutil/keeper"
	"mycel/testutil/nullify"
	"mycel/x/registry/keeper"
	"mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDomain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Domain {
	items := make([]types.Domain, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)
		items[i].Parent = strconv.Itoa(i)

		keeper.SetDomain(ctx, items[i])
	}
	return items
}

func TestDomainGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNDomain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDomain(ctx,
			item.Name,
			item.Parent,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDomainRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNDomain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDomain(ctx,
			item.Name,
			item.Parent,
		)
		_, found := keeper.GetDomain(ctx,
			item.Name,
			item.Parent,
		)
		require.False(t, found)
	}
}

func TestDomainGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNDomain(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDomain(ctx)),
	)
}
