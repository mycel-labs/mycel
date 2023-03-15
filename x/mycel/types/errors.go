package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mycel module sentinel errors
var (
	ErrSample                  = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrorDomainIsAlreadyTaken  = sdkerrors.Register(ModuleName, 1101, "domain is already taken")
	ErrorDomainNameIsInvalid   = sdkerrors.Register(ModuleName, 1102, "name is invalid")
	ErrorDomainParentIsInvalid = sdkerrors.Register(ModuleName, 1103, "parent is invalid")
)
