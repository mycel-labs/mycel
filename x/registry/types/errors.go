package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// top-level-domain sentinel errors
var (
	ErrInvalidTopLevelDomainName               = errorsmod.Register(ModuleName, 1000, "invalid top-level-domain name")
	ErrTopLevelDomainNotFound                  = errorsmod.Register(ModuleName, 1001, "top-level-domain not found")
	ErrTopLevelDomainExpired                   = errorsmod.Register(ModuleName, 1002, "top-level-domain expired")
	ErrTopLevelDomainAlreadyTaken              = errorsmod.Register(ModuleName, 1003, "top-level-domain already taken")
	ErrTopLevelDomainNotRegistrable            = errorsmod.Register(ModuleName, 1004, "top-level-domain not registrable")
	ErrTopLevelDomainMaxSubdomainCountReached  = errorsmod.Register(ModuleName, 1005, "top-level-domain max subdomain count reached")
	ErrTopLevelDomainInvalidRegistrationPeriod = errorsmod.Register(ModuleName, 1006, "top-level-domain invalid registration period")
	ErrTopLevelDomainNotEditable               = errorsmod.Register(ModuleName, 1007, "top-level-domain not editable")
)

// second-level-domain sentinel errors
var (
	ErrInvalidSecondLevelDomainName               = errorsmod.Register(ModuleName, 1100, "invalid second-level-domain name")
	ErrInvalidSecondLevelDomainParent             = errorsmod.Register(ModuleName, 1101, "invalid second-level-domain parent")
	ErrSecondLevelDomainNotFound                  = errorsmod.Register(ModuleName, 1102, "second-level-domain not found")
	ErrSecondLevelDomainExpired                   = errorsmod.Register(ModuleName, 1103, "second-level-domain expired")
	ErrSecondLevelDomainAlreadyTaken              = errorsmod.Register(ModuleName, 1104, "second-level-domain already taken")
	ErrSecondLevelDomainParentDoesNotExist        = errorsmod.Register(ModuleName, 1105, "second-level-domain parent does not exist")
	ErrSecondLevelDomainInvalidRegistrationPeriod = errorsmod.Register(ModuleName, 1106, "second-level-domain invalid registration period")
	ErrSecondLevelDomainNotEditable               = errorsmod.Register(ModuleName, 1107, "second-level-domain not editable")
	ErrSecondLevelDomainNotRegistrable            = errorsmod.Register(ModuleName, 1108, "second-level-domain not registrable")
)

// record sentinel errors
var (
	ErrInvalidWalletAddress    = errorsmod.Register(ModuleName, 1200, "invalid wallet address")
	ErrInvalidWalletRecordType = errorsmod.Register(ModuleName, 1201, "invalid wallet record type")
	ErrInvalidDnsRecordValue   = errorsmod.Register(ModuleName, 1202, "invalid dns record value")
	ErrInvalidDnsRecordType    = errorsmod.Register(ModuleName, 1203, "invalid dns record type")
	ErrInvalidTextRecordKey    = errorsmod.Register(ModuleName, 1204, "invalid text record key, this key is reserved")
)

// withdraw sentinel errors
var (
	ErrNoEnoughBalance              = errorsmod.Register(ModuleName, 1300, "no enough balance")
	ErrNoPermissionToWithdraw       = errorsmod.Register(ModuleName, 1301, "no permission to withdraw")
	ErrNoWithdrawalAmountToWithdraw = errorsmod.Register(ModuleName, 1302, "no registration fee to withdraw")
	ErrInvalidDenom                 = errorsmod.Register(ModuleName, 1303, "invalid denom")
)

// policy sentinel errors
var (
	ErrInvalidRegistrationPolicy = errorsmod.Register(ModuleName, 1400, "invalid registration policy")
)
