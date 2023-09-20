package types

import (
	"errors"
	"time"

	errorsmod "cosmossdk.io/errors"
)

// DefaultIndex is the default global index

func NewGenesisState(epochs []EpochInfo) *GenesisState {
	return &GenesisState{Epochs: epochs}
}

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
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in epochInfo
	epochIdentifiers := make(map[string]bool)

	for _, epoch := range gs.Epochs {
		if epochIdentifiers[epoch.Identifier] {
			return errorsmod.Wrapf(errors.New(epoch.Identifier), ErrDuplicatedEpochEntry.Error())
		}
		if err := epoch.Validate(); err != nil {
			return err
		}
		epochIdentifiers[epoch.Identifier] = true
	}
	return nil
}
