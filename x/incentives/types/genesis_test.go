package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"mycel/x/incentives/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				IncentiveList: []types.Incentive{
					{
						Epoch: 0,
					},
					{
						Epoch: 1,
					},
				},
				EpochIncentiveList: []types.EpochIncentive{
					{
						Epoch: 0,
					},
					{
						Epoch: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated incentive",
			genState: &types.GenesisState{
				IncentiveList: []types.Incentive{
					{
						Epoch: 0,
					},
					{
						Epoch: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated epochIncentive",
			genState: &types.GenesisState{
				EpochIncentiveList: []types.EpochIncentive{
					{
						Epoch: 0,
					},
					{
						Epoch: 0,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
