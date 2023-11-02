package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/furnace module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidRegistrationPeriod = sdkerrors.Register(ModuleName, 1101, "invalid registration period")
)
