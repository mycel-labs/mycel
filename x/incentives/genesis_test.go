package incentives_test

import (
	"testing"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/incentives"
	"github.com/mycel-domain/mycel/x/incentives/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		IncentiveList: []types.Incentive{
			{
				Epoch: 0,
			},
			{
				Epoch: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IncentivesKeeper(t)
	incentives.InitGenesis(ctx, *k, genesisState)
	got := incentives.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.IncentiveList, got.IncentiveList)
	// this line is used by starport scaffolding # genesis/test/assert
}
