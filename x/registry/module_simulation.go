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
	opWeightMsgUpdateWalletRecord = "op_weight_msg_update_wallet_record" // #nosec G101
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWalletRecord int = 100

	opWeightMsgRegisterSecondLevelDomain = "op_weight_msg_register_domain" // #nosec G101
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterSecondLevelDomain int = 100

	opWeightMsgRegisterTopLevelDomain = "op_weight_msg_register_top_level_domain" // #nosec G101
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterTopLevelDomain int = 100

	opWeightMsgWithdrawRegistrationFee = "op_weight_msg_withdraw_registration_fee" // #nosec G101
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawRegistrationFee int = 100

	opWeightMsgExtendTopLevelDomainExpirationDate = "op_weight_msg_extend_top_level_domain_expiration" // #nosec G101
	// TODO: Determine the simulation weight value
	defaultWeightMsgExtendTopLevelDomainExpirationDate int = 100

	opWeightMsgUpdateTextRecord = "op_weight_msg_update_text_record" // #nosec G101
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTextRecord int = 100

	opWeightMsgUpdateTopLevelDomainRegistrationPolicy = "op_weight_msg_update_top_level_domain_registration_policy" //#nosec G101
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTopLevelDomainRegistrationPolicy int = 100

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

// ProposalMsgs doesn't return any content functions for governance proposals
func (AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
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

	var weightMsgRegisterSecondLevelDomain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterSecondLevelDomain, &weightMsgRegisterSecondLevelDomain, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterSecondLevelDomain = defaultWeightMsgRegisterSecondLevelDomain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterSecondLevelDomain,
		registrysimulation.SimulateMsgRegisterSecondLevelDomain(am.accountKeeper, am.bankKeeper, am.keeper),
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

	var weightMsgWithdrawRegistrationFee int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawRegistrationFee, &weightMsgWithdrawRegistrationFee, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawRegistrationFee = defaultWeightMsgWithdrawRegistrationFee
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawRegistrationFee,
		registrysimulation.SimulateMsgWithdrawRegistrationFee(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgExtendTopLevelDomainExpirationDate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgExtendTopLevelDomainExpirationDate, &weightMsgExtendTopLevelDomainExpirationDate, nil,
		func(_ *rand.Rand) {
			weightMsgExtendTopLevelDomainExpirationDate = defaultWeightMsgExtendTopLevelDomainExpirationDate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgExtendTopLevelDomainExpirationDate,
		registrysimulation.SimulateMsgExtendTopLevelDomainExpirationDate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTextRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTextRecord, &weightMsgUpdateTextRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTextRecord = defaultWeightMsgUpdateTextRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTextRecord,
		registrysimulation.SimulateMsgUpdateTextRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTopLevelDomainRegistrationPolicy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTopLevelDomainRegistrationPolicy, &weightMsgUpdateTopLevelDomainRegistrationPolicy, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTopLevelDomainRegistrationPolicy = defaultWeightMsgUpdateTopLevelDomainRegistrationPolicy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTopLevelDomainRegistrationPolicy,
		registrysimulation.SimulateMsgUpdateTopLevelDomainRegistrationPolicy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
