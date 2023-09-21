package types_test

import (
	"testing"

	"github.com/mycel-domain/mycel/x/furnace/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
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

				EpochBurnConfig: types.EpochBurnConfig{
					EpochIdentifier: "35",
				},
				BurnAmountList: []types.BurnAmount{
					{
						Identifier: 0,
					},
					{
						Identifier: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated burnAmount",
			genState: &types.GenesisState{
				BurnAmountList: []types.BurnAmount{
					{
						Identifier: 0,
					},
					{
						Identifier: 0,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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
