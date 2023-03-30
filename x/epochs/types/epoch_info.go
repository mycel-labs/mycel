package types

import (
	"errors"
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// StartInitialEpoch sets the epoch info fields to their start values
func (ei *EpochInfo) StartInitialEpoch() {
	ei.EpochCountingStarted = true
	ei.CurrentEpoch = 1
	ei.CurrentEpochStartTime = ei.StartTime
}

// EndEpoch increments the epoch counter and resets the epoch start time
func (ei *EpochInfo) EndEpoch() {
	ei.CurrentEpoch++
	ei.CurrentEpochStartTime = ei.CurrentEpochStartTime.Add(ei.Duration)
}

// Validate performs a stateless validation of the epoch info fields
func (ei EpochInfo) Validate() error {
	if strings.TrimSpace(ei.Identifier) == "" {
		return ErrEpochIdentifierCannotBeEmpty
	}
	if ei.Duration == 0 {
		return ErrEpochDurationCannotBeZero
	}
	if ei.CurrentEpoch < 0 {
		return sdkerrors.Wrapf(errors.New(fmt.Sprintf("%d", ei.CurrentEpoch)), ErrCurrentEpochCannotBeNegative.Error())
	}
	if ei.CurrentEpochStartHeight < 0 {
		return sdkerrors.Wrapf(errors.New(fmt.Sprintf("%d", ei.CurrentEpoch)), ErrCurrentEpochStartHeightCannotBeNegative.Error())
	}
	return nil
}
