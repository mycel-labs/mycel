package registry

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/mycel-domain/mycel/testutil/sample"
	registrysimulation "github.com/mycel-domain/mycel/x/registry/simulation"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = registrysimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgUpdateWalletRecord = "op_weight_msg_update_wallet_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWalletRecord int = 100

	opWeightMsgRegisterDomain = "op_weight_msg_register_domain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterDomain int = 100

	opWeightMsgRegisterTopLevelDomain = "op_weight_msg_register_top_level_domain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterTopLevelDomain int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	registryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&registryGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgUpdateWalletRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWalletRecord, &weightMsgUpdateWalletRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWalletRecord = defaultWeightMsgUpdateWalletRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWalletRecord,
		registrysimulation.SimulateMsgUpdateWalletRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterDomain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterDomain, &weightMsgRegisterDomain, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterDomain = defaultWeightMsgRegisterDomain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterDomain,
		registrysimulation.SimulateMsgRegisterDomain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterTopLevelDomain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterTopLevelDomain, &weightMsgRegisterTopLevelDomain, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterTopLevelDomain = defaultWeightMsgRegisterTopLevelDomain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterTopLevelDomain,
		registrysimulation.SimulateMsgRegisterTopLevelDomain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
