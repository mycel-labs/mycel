package furnace_test

import (
	"testing"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/furnace"
	"github.com/mycel-domain/mycel/x/furnace/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EpochBurnConfig: &types.EpochBurnConfig{
			EpochIdentifier: "11",
		},
		BurnAmounts: []types.BurnAmount{
			{
				Identifier: 0,
			},
			{
				Identifier: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FurnaceKeeper(t)
	furnace.InitGenesis(ctx, *k, genesisState)
	got := furnace.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.EpochBurnConfig, got.EpochBurnConfig)
	require.ElementsMatch(t, genesisState.BurnAmounts, got.BurnAmountList)
	// this line is used by starport scaffolding # genesis/test/assert
}
