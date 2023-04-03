package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EpochIncentiveList: []EpochIncentive{},
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
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
