package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/furnace module sentinel errors
var (
	ErrSample                    = errorsmod.Register(ModuleName, 1100, "sample error")
	ErrInvalidRegistrationPeriod = errorsmod.Register(ModuleName, 1101, "invalid registration period")
)
