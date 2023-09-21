package types

import (
	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
	"time"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

func GetDefaultEpochBurnConfig() EpochBurnConfig {
	return EpochBurnConfig{
		EpochIdentifier:             epochstypes.DailyEpochId,
		CurrentBurnAmountIdentifier: 1,
		Duration:                    time.Hour * 24 * 30 * 4,
	}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EpochBurnConfig: GetDefaultEpochBurnConfig(),
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
