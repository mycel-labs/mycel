package registry

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the domain
	for _, elem := range genState.Domains {
		k.SetDomain(ctx, elem)
	}
	// Set all the domainOwnership
	for _, elem := range genState.DomainOwnerships {
		k.SetDomainOwnership(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.Domains = k.GetAllDomain(ctx)
	genesis.DomainOwnerships = k.GetAllDomainOwnership(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
