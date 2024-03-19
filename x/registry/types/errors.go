package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/registry module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
)

// top-level-domain sentinel errors
var (
	ErrInvalidTopLevelDomainName               = sdkerrors.Register(ModuleName, 0o00, "invalid top-level-domain name")
	ErrTopLevelDomainNotFound                  = sdkerrors.Register(ModuleName, 0o01, "top-level-domain not found")
	ErrTopLevelDomainExpired                   = sdkerrors.Register(ModuleName, 0o02, "top-level-domain expired")
	ErrTopLevelDomainAlreadyTaken              = sdkerrors.Register(ModuleName, 0o03, "top-level-domain already taken")
	ErrTopLevelDomainNotRegistrable            = sdkerrors.Register(ModuleName, 0o04, "top-level-domain not registrable")
	ErrTopLevelDomainMaxSubdomainCountReached  = sdkerrors.Register(ModuleName, 0o05, "top-level-domain max subdomain count reached")
	ErrTopLevelDomainInvalidRegistrationPeriod = sdkerrors.Register(ModuleName, 0o06, "top-level-domain invalid registration period")
	ErrTopLevelDomainNotEditable               = sdkerrors.Register(ModuleName, 0o07, "top-level-domain not editable")
)

// second-level-domain sentinel errors
var (
	ErrInvalidSecondLevelDomainName               = sdkerrors.Register(ModuleName, 100, "invalid second-level-domain name")
	ErrInvalidSecondLevelDomainParent             = sdkerrors.Register(ModuleName, 101, "invalid second-level-domain parent")
	ErrSecondLevelDomainNotFound                  = sdkerrors.Register(ModuleName, 102, "second-level-domain not found")
	ErrSecondLevelDomainExpired                   = sdkerrors.Register(ModuleName, 103, "second-level-domain expired")
	ErrSecondLevelDomainAlreadyTaken              = sdkerrors.Register(ModuleName, 104, "second-level-domain already taken")
	ErrSecondLevelDomainParentDoesNotExist        = sdkerrors.Register(ModuleName, 105, "second-level-domain parent does not exist")
	ErrSecondLevelDomainInvalidRegistrationPeriod = sdkerrors.Register(ModuleName, 106, "second-level-domain invalid registration period")
	ErrSecondLevelDomainNotEditable               = sdkerrors.Register(ModuleName, 107, "second-level-domain not editable")
	ErrSecondLevelDomainNotRegistrable            = sdkerrors.Register(ModuleName, 108, "second-level-domain not registrable")
)

// record sentinel errors
var (
	ErrInvalidWalletAddress    = sdkerrors.Register(ModuleName, 200, "invalid wallet address")
	ErrInvalidWalletRecordType = sdkerrors.Register(ModuleName, 201, "invalid wallet record type")
	ErrInvalidDnsRecordValue   = sdkerrors.Register(ModuleName, 202, "invalid dns record value")
	ErrInvalidDnsRecordType    = sdkerrors.Register(ModuleName, 203, "invalid dns record type")
	ErrInvalidTextRecordKey    = sdkerrors.Register(ModuleName, 204, "invalid text record key, this key is reserved")
)

// withdraw sentinel errors
var (
	ErrNoEnoughBalance              = sdkerrors.Register(ModuleName, 300, "no enough balance")
	ErrNoPermissionToWithdraw       = sdkerrors.Register(ModuleName, 301, "no permission to withdraw")
	ErrNoWithdrawalAmountToWithdraw = sdkerrors.Register(ModuleName, 302, "no registration fee to withdraw")
	ErrInvalidDenom                 = sdkerrors.Register(ModuleName, 303, "invalid denom")
)

// policy sentinel errors
var (
	ErrInvalidRegistrationPolicy = sdkerrors.Register(ModuleName, 400, "invalid registration policy")
	ErrNotAllowedRegisterDomain  = sdkerrors.Register(ModuleName, 401, "not allowed to regsiter under private domain")
)
