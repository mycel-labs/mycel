package types

import (
	"errors"
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
)

// StartInitialEpoch sets the epoch info fields to ther start values
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
		return errorsmod.Wrapf(errors.New(fmt.Sprintf("%d", e.CurrentEpoch)), ErrCurrentEpochCannotBeNegative.Error())
	}
	if e.CurrentEpochStartHeight < 0 {
		return errorsmod.Wrapf(errors.New(fmt.Sprintf("%d", e.CurrentEpoch)), ErrCurrentEpochStartHeightCannotBeNegative.Error())
	}
	return nil
}
