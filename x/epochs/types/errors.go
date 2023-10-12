package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/epochs module sentinel errors
var (
	ErrSample                                  = errorsmod.Register(ModuleName, 1100, "sample error")
	ErrEpochIdentifierCannotBeEmpty            = errorsmod.Register(ModuleName, 1101, "epoch identifier cannot be empty")
	ErrEpochDurationCannotBeZero               = errorsmod.Register(ModuleName, 1102, "epoch duration cannot be zero")
	ErrCurrentEpochCannotBeNegative            = errorsmod.Register(ModuleName, 1103, "current epoch cannot be negative")
	ErrCurrentEpochStartHeightCannotBeNegative = errorsmod.Register(ModuleName, 1104, "current epoch start height cannot be negative")
	ErrDuplicatedEpochEntry                    = errorsmod.Register(ModuleName, 1105, "duplicated epoch entry")
	ErrInvalidParameterType                    = errorsmod.Register(ModuleName, 1106, "invalid parameter type")
	ErrEmptyDistributionEpochIdentifier        = errorsmod.Register(ModuleName, 1107, "empty distribution epoch identifier")
)
