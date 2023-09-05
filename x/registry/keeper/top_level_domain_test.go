package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTopLevelDomain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TopLevelDomain {
	items := make([]types.TopLevelDomain, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetTopLevelDomain(ctx, items[i])
	}
	return items
}

func TestTopLevelDomainGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNTopLevelDomain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTopLevelDomain(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTopLevelDomainRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNTopLevelDomain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTopLevelDomain(ctx,
			item.Name,
		)
		_, found := keeper.GetTopLevelDomain(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestTopLevelDomainGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNTopLevelDomain(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTopLevelDomain(ctx)),
	)
}
