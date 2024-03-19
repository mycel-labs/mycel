package types

import (
	errorsmod "cosmossdk.io/errors"
)

const (
	WeeklyEpochId = "weekly"
	DailyEpochId  = "daily"
)

func ValidateEpochIdentifierInterface(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return errorsmod.Wrapf(ErrInvalidParameterType, "%T", i)
	}

	return ValidateEpochIdentifierString(v)
}

func ValidateEpochIdentifierString(s string) error {
	if s == "" {
		return errorsmod.Wrapf(ErrEmptyDistributionEpochIdentifier, "%v", s)
	}
	return nil
}
