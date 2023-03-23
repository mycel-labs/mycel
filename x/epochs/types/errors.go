package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/epochs module sentinel errors
var (
	ErrSample                                  = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrEpochIdentifierCannotBeEmpty            = sdkerrors.Register(ModuleName, 1101, "epoch identifier cannot be empty")
	ErrEpochDurationCannotBeZero               = sdkerrors.Register(ModuleName, 1102, "epoch duration cannot be zero")
	ErrCurrentEpochCannotBeNegative            = sdkerrors.Register(ModuleName, 1103, "current epoch cannot be negative")
	ErrCurrentEpochStartHeightCannotBeNegative = sdkerrors.Register(ModuleName, 1104, "current epoch start height cannot be negative")
	ErrDuplicatedEpochEntry                    = sdkerrors.Register(ModuleName, 1105, "duplicated epoch entry")
)
