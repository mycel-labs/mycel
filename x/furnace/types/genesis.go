package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

func GetDefaultEpochBurnConfig() EpochBurnConfig {
	return EpochBurnConfig{
		EpochIdentifier:        epochstypes.DailyEpochId,
		CurrentBurnAmountIndex: 1,
		DefaultTotalEpochs:     120,
	}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EpochBurnConfig: GetDefaultEpochBurnConfig(),
		BurnAmounts: []BurnAmount{
			{
				Index:                 0,
				BurnStarted:           false,
				TotalEpochs:           120,
				CurrentEpoch:          0,
				TotalBurnAmount:       sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0)),
				CumulativeBurntAmount: sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0)),
			},
		},
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
		index := string(BurnAmountKey(elem.Index))
		if _, ok := burnAmountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for burnAmount")
		}
		burnAmountIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
