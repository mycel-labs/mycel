package incentives_test

import (
	"testing"

	keepertest "mycel/testutil/keeper"
	"mycel/testutil/nullify"
	"mycel/x/incentives"
	"mycel/x/incentives/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EpochIncentiveList: []types.EpochIncentive{
			{
				Epoch: 0,
			},
			{
				Epoch: 1,
			},
		},
		ValidatorIncentiveList: []types.ValidatorIncentive{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		DelegetorIncentiveList: []types.DelegetorIncentive{
			{
				Address: "0",
			},
			{
				Address: "1",
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

	require.ElementsMatch(t, genesisState.EpochIncentiveList, got.EpochIncentiveList)
	require.ElementsMatch(t, genesisState.ValidatorIncentiveList, got.ValidatorIncentiveList)
	require.ElementsMatch(t, genesisState.DelegetorIncentiveList, got.DelegetorIncentiveList)
	// this line is used by starport scaffolding # genesis/test/assert
}
