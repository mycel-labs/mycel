package types

import (
	"errors"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DefaultIndex is the default global index

func NewGenesisState(epochs []EpochInfo) *GenesisState {
	return &GenesisState{Epochs: epochs}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Epochs: []EpochInfo{
			{
				Identifier:              WeeklyEpochId,
				StartTime:               time.Time{},
				Duration:                time.Hour * 24 * 7,
				CurrentEpoch:            0,
				CurrentEpochStartHeight: 0,
				CurrentEpochStartTime:   time.Time{},
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
			return sdkerrors.Wrapf(errors.New(epoch.Identifier), ErrDuplicatedEpochEntry.Error())
		}
		if err := epoch.Validate(); err != nil {
			return err
		}
		epochIdentifiers[epoch.Identifier] = true
	}
	return nil
}
