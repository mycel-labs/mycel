package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/github.com/mycel-domain/mycel module sentinel errors
var (
	ErrSample                   = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrDomainIsAlreadyTaken     = sdkerrors.Register(ModuleName, 1101, "domain is already taken")
	ErrInvalidDomainName        = sdkerrors.Register(ModuleName, 1102, "invalid name")
	ErrInvalidDomainParent      = sdkerrors.Register(ModuleName, 1103, "invalid parent")
	ErrDomainNotFound           = sdkerrors.Register(ModuleName, 1104, "domain not found")
	ErrInvalidWalletAddress     = sdkerrors.Register(ModuleName, 1105, "invalid wallet address")
	ErrInvalidWalletRecordType  = sdkerrors.Register(ModuleName, 1106, "invalid wallet record type")
	ErrDomainNotOwned           = sdkerrors.Register(ModuleName, 1107, "domain not owned by msg creator")
	ErrParentDomainDoesNotExist = sdkerrors.Register(ModuleName, 1108, "parent domain does not exist")
	ErrParentDomainMustBeEmpty  = sdkerrors.Register(ModuleName, 1109, "parent domain must be empty")
)
