package registry_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RegistryKeeper(t)
	registry.InitGenesis(ctx, *k, genesisState)
	got := registry.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TopLevelDomains, got.TopLevelDomains)
	require.ElementsMatch(t, genesisState.SecondLevelDomains, got.SecondLevelDomains)
	require.ElementsMatch(t, genesisState.DomainOwnerships, got.DomainOwnerships)
	// this line is used by starport scaffolding # genesis/test/assert
}
