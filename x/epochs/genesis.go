package epochs

import (
	"mycel/x/epochs/keeper"
	"mycel/x/epochs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the epochInfo
	for _, elem := range genState.Epochs {
		k.SetEpochInfo(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Epochs = k.GetAllEpochInfo(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
