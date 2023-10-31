package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/mycel module sentinel errors
var (
	ErrSample                       = errorsmod.Register(ModuleName, 1100, "sample error")
	ErrDomainIsAlreadyTaken         = errorsmod.Register(ModuleName, 1101, "domain is already taken")
	ErrInvalidDomainName            = errorsmod.Register(ModuleName, 1102, "invalid name")
	ErrInvalidDomainParent          = errorsmod.Register(ModuleName, 1103, "invalid parent")
	ErrDomainNotFound               = errorsmod.Register(ModuleName, 1104, "domain not found")
	ErrInvalidWalletAddress         = errorsmod.Register(ModuleName, 1105, "invalid wallet address")
	ErrInvalidWalletRecordType      = errorsmod.Register(ModuleName, 1106, "invalid wallet record type")
	ErrInvalidDnsRecordValue        = errorsmod.Register(ModuleName, 1107, "invalid dns record value")
	ErrInvalidDnsRecordType         = errorsmod.Register(ModuleName, 1108, "invalid dns record type")
	ErrDomainNotEditable            = errorsmod.Register(ModuleName, 1109, "role not pemitted to edit the domain")
	ErrParentDomainDoesNotExist     = errorsmod.Register(ModuleName, 1110, "parent domain does not exist")
	ErrParentDomainMustBeEmpty      = errorsmod.Register(ModuleName, 1111, "parent domain must be empty")
	ErrDomainNotRegistrable         = errorsmod.Register(ModuleName, 1112, "domain is not registrable")
	ErrMaxSubdomainCountReached     = errorsmod.Register(ModuleName, 1113, "max subdomain count reached")
	ErrInvalidRegistrationPeriod    = errorsmod.Register(ModuleName, 1114, "invalid registration period")
	ErrDomainExpired                = errorsmod.Register(ModuleName, 1115, "domain expired")
	ErrNoEnoughBalance              = errorsmod.Register(ModuleName, 1116, "no enough balance")
	ErrNoPermissionToWithdraw       = errorsmod.Register(ModuleName, 1117, "no permission to withdraw")
	ErrNoWithdrawalAmountToWithdraw = errorsmod.Register(ModuleName, 1118, "no registration fee to withdraw")
	ErrInvalidDenom                 = errorsmod.Register(ModuleName, 1119, "invalid denom")
)
