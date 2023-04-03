package types_test

import (
	"testing"

	"mycel/x/incentives/types"

	"github.com/stretchr/testify/require"
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc:     "duplicated incentive",
			genState: &types.GenesisState{},
			valid:    false,
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
		{
			desc: "duplicated validatorIncentive",
			genState: &types.GenesisState{
				ValidatorIncentiveList: []types.ValidatorIncentive{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated delegetorIncentive",
			genState: &types.GenesisState{
				DelegetorIncentiveList: []types.DelegetorIncentive{
					{
						Address: "0",
					},
					{
						Address: "0",
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
