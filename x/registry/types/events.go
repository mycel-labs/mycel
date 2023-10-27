package types

// Register second-level-domain event
const (
	EventTypeRegisterDomain = "register-domain"

	AttributeRegisterSecondLevelDomainEventName            = "name"
	AttributeRegisterSecondLevelDomainEventParent          = "parent"
	AttributeRegisterSecondLevelDomainEventExpirationDate  = "expiration-date"
	AttributeRegisterSecondLevelDomainEventRegistrationFee = "registration-fee"
)

// Register top-level-domain event
const (
	EventTypeRegisterTopLevelDomain = "register-top-level-domain"

	AttributeRegisterTopLevelDomainEventName           = "name"
	AttributeRegisterTopLevelDomainEventExpirationDate = "expiration-date"
)

// Update wallet record event
const (
	EventTypeUpdateWalletRecord = "update-wallet-record"

	AttributeUpdateWalletRecordEventDomainName       = "name"
	AttributeUpdateWalletRecordEventDomainParent     = "parent"
	AttributeUpdateWalletRecordEventWalletRecordType = "wallet-record-type"
	AttributeUpdateWalletRecordEventValue            = "value"
)

// Update dns record event
const (
	EventTypeUpdateDnsRecord = "update-dns-record"

	AttributeUpdateDnsRecordEventDomainName    = "name"
	AttributeUpdateDnsRecordEventDomainParent  = "parent"
	AttributeUpdateDnsRecordEventDnsRecordType = "dns-record-type"
	AttributeUpdateDnsRecordEventValue         = "value"
)

// Withdraw fees event
const (
	EventTypeWithdrawRegistrationFee = "withdraw-registration-fees"

	AttributeWithdrawRegistrationFeeEventDomainName = "name"
	AttributeWithdrawRegistrationFeeEventDomainFee  = "fee"
)
