package epochs_test

import (
	"testing"

	keepertest "mycel/testutil/keeper"
	"mycel/testutil/nullify"
	"mycel/x/epochs"
	"mycel/x/epochs/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Epochs: []types.EpochInfo{
			{
				Identifier: "0",
			},
			{
				Identifier: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EpochsKeeper(t)
	epochs.InitGenesis(ctx, *k, genesisState)
	got := epochs.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.Epochs, got.Epochs)
	// this line is used by starport scaffolding # genesis/test/assert
}
