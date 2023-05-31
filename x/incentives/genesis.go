package incentives

import (
	"github.com/mycel-domain/mycel/x/incentives/keeper"
	"github.com/mycel-domain/mycel/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the epochIncentive
	for _, elem := range genState.EpochIncentiveList {
		k.SetEpochIncentive(ctx, elem)
	}
	// Set all the validatorIncentive
	for _, elem := range genState.ValidatorIncentiveList {
		k.SetValidatorIncentive(ctx, elem)
	}
	// Set all the delegetorIncentive
	for _, elem := range genState.DelegetorIncentiveList {
		k.SetDelegetorIncentive(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.EpochIncentiveList = k.GetAllEpochIncentive(ctx)
	genesis.ValidatorIncentiveList = k.GetAllValidatorIncentive(ctx)
	genesis.DelegetorIncentiveList = k.GetAllDelegetorIncentive(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
