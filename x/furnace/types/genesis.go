package types

import (
	"fmt"
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
		BurnAmounts:  []BurnAmount{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in burnAmount
	burnAmountIndexMap := make(map[string]struct{})

	for _, elem := range gs.BurnAmounts {
		index := string(BurnAmountKey(elem.Identifier))
		if _, ok := burnAmountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for burnAmount")
		}
		burnAmountIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
