package types

import (
	"fmt"
	"time"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	startTime := time.Time{}
	return &GenesisState{
		Epochs: []EpochInfo{
			{
				Identifier:              WeeklyEpochId,
				StartTime:               time.Time{},
				Duration:                time.Hour * 24 * 7,
				CurrentEpoch:            0,
				CurrentEpochStartHeight: 0,
				CurrentEpochStartTime:   startTime,
				EpochCountingStarted:    false,
			},
			{
				Identifier:              DailyEpochId,
				StartTime:               time.Time{},
				Duration:                time.Hour * 24,
				CurrentEpoch:            0,
				CurrentEpochStartHeight: 0,
				CurrentEpochStartTime:   startTime,
				EpochCountingStarted:    false,
			},
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in epochInfo
	epochInfoIndexMap := make(map[string]struct{})

	for _, elem := range gs.Epochs {
		index := string(EpochInfoKey(elem.Identifier))
		if _, ok := epochInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for epochInfo")
		}
		epochInfoIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
