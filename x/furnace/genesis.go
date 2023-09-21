package furnace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/x/furnace/keeper"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set genesis time
	genState.EpochBurnConfig.StartTime = ctx.BlockTime()
	k.SetEpochBurnConfig(ctx, genState.EpochBurnConfig)
	// Set all the burnAmount
	for _, elem := range genState.BurnAmountList {
		k.SetBurnAmount(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all epochBurnConfig
	epochBurnConfig, found := k.GetEpochBurnConfig(ctx)
	if found {
		genesis.EpochBurnConfig = epochBurnConfig
	}
	genesis.BurnAmountList = k.GetAllBurnAmount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
