package epochs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/epochs/keeper"
	"github.com/mycel-domain/mycel/x/epochs/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the epochInfo
	for _, elem := range genState.Epochs {
		k.SetEpochInfo(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.Epochs = k.GetAllEpochInfo(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
