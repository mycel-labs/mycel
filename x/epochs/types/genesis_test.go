package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type GenesisTestSuite struct {
	suite.Suite
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) TestValidateGenesis() {
	testCases := []struct {
		err      error
		genState *GenesisState
		expPass  bool
	}{
		{
			nil,
			DefaultGenesis(),
			true,
		},
		{
			nil,
			&GenesisState{
				Epochs: []EpochInfo{},
			},
			true,
		},
		{
			nil,
			&GenesisState{
				Epochs: []EpochInfo{
					{
						Identifier:              DailyEpochId,
						StartTime:               time.Time{},
						Duration:                time.Hour * 24 * 7,
						CurrentEpoch:            0,
						CurrentEpochStartHeight: 0,
						CurrentEpochStartTime:   time.Time{},
						EpochCountingStarted:    false,
					},
				},
			},
			true,
		},
		{
			ErrDuplicatedEpochEntry,
			&GenesisState{
				Epochs: []EpochInfo{
					{
						Identifier:              DailyEpochId,
						StartTime:               time.Time{},
						Duration:                time.Hour * 24 * 7,
						CurrentEpoch:            0,
						CurrentEpochStartHeight: 0,
						CurrentEpochStartTime:   time.Time{},
						EpochCountingStarted:    false,
					},
					{
						Identifier:              DailyEpochId,
						StartTime:               time.Time{},
						Duration:                time.Hour * 24 * 7,
						CurrentEpoch:            0,
						CurrentEpochStartHeight: 0,
						CurrentEpochStartTime:   time.Time{},
						EpochCountingStarted:    false,
					},
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.genState.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.err)
		} else {
			suite.Require().Error(err, tc.err)
		}
	}
}
