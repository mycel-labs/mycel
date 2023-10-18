package types

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
)

// StartInitialEpoch sets the epoch info fields to their start values
func (e *EpochInfo) StartInitialEpoch() {
	e.EpochCountingStarted = true
	e.CurrentEpoch = 1
	e.CurrentEpochStartTime = e.StartTime
}

// EndEpoch increments the epoch counter and resets the epoch start time
func (e *EpochInfo) EndEpoch() {
	e.CurrentEpoch++
	e.CurrentEpochStartTime = e.CurrentEpochStartTime.Add(e.Duration)
}

// Validate performs a stateless validation of the epoch info fields
func (e EpochInfo) Validate() error {
	if strings.TrimSpace(e.Identifier) == "" {
		return ErrEpochIdentifierCannotBeEmpty
	}
	if e.Duration == 0 {
		return ErrEpochDurationCannotBeZero
	}
	if e.CurrentEpoch < 0 {
		return errorsmod.Wrapf(ErrCurrentEpochCannotBeNegative, "%d", e.CurrentEpoch)
	}
	if e.CurrentEpochStartHeight < 0 {
		return errorsmod.Wrapf(ErrCurrentEpochStartHeightCannotBeNegative, "%d", e.CurrentEpoch)
	}
	return nil
}
