package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mycel-domain/mycel/x/registry/types"
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
				TopLevelDomains: []types.TopLevelDomain{
					{
						Name: "0",
					},
					{
						Name: "1",
					},
				},
				SecondLevelDomains: []types.SecondLevelDomain{
					{
						Name:   "0",
						Parent: "0",
					},
					{
						Name:   "1",
						Parent: "1",
					},
				},
				DomainOwnerships: []types.DomainOwnership{
					{
						Owner: "0",
					},
					{
						Owner: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated topLevelDomain",
			genState: &types.GenesisState{
				TopLevelDomains: []types.TopLevelDomain{
					{
						Name: "0",
					},
					{
						Name: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated secondLevelDomain",
			genState: &types.GenesisState{
				SecondLevelDomains: []types.SecondLevelDomain{
					{
						Name:   "0",
						Parent: "0",
					},
					{
						Name:   "0",
						Parent: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated domainOwnership",
			genState: &types.GenesisState{
				DomainOwnerships: []types.DomainOwnership{
					{
						Owner: "0",
					},
					{
						Owner: "0",
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
