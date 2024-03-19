package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/epochs module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
)

var (
	ErrEpochIdentifierCannotBeEmpty            = sdkerrors.Register(ModuleName, 1101, "epoch identifier cannot be empty")
	ErrEpochDurationCannotBeZero               = sdkerrors.Register(ModuleName, 1102, "epoch duration cannot be zero")
	ErrCurrentEpochCannotBeNegative            = sdkerrors.Register(ModuleName, 1103, "current epoch cannot be negative")
	ErrCurrentEpochStartHeightCannotBeNegative = sdkerrors.Register(ModuleName, 1104, "current epoch start height cannot be negative")
	ErrDuplicatedEpochEntry                    = sdkerrors.Register(ModuleName, 1105, "duplicated epoch entry")
	ErrInvalidParameterType                    = sdkerrors.Register(ModuleName, 1106, "invalid parameter type")
	ErrEmptyDistributionEpochIdentifier        = sdkerrors.Register(ModuleName, 1107, "empty distribution epoch identifier")
)
