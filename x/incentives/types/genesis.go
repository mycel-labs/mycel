package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EpochIncentiveList:     []EpochIncentive{},
		ValidatorIncentiveList: []ValidatorIncentive{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in epochIncentive
	epochIncentiveIndexMap := make(map[string]struct{})

	for _, elem := range gs.EpochIncentiveList {
		index := string(EpochIncentiveKey(elem.Epoch))
		if _, ok := epochIncentiveIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for epochIncentive")
		}
		epochIncentiveIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in validatorIncentive
	validatorIncentiveIndexMap := make(map[string]struct{})

	for _, elem := range gs.ValidatorIncentiveList {
		index := string(ValidatorIncentiveKey(elem.Address))
		if _, ok := validatorIncentiveIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for validatorIncentive")
		}
		validatorIncentiveIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
