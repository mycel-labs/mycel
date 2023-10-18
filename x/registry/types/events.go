package types

// Register second-level-domain event
const (
	EventTypeRegsterDomain = "register-domain"

	AttributeRegisterSecondLevelDomainEventName            = "name"
	AttributeRegisterSecondLevelDomainEventParent          = "parent"
	AttributeRegisterSecondLevelDomainEventExpirationDate  = "expiration-date"
	AttributeRegisterSecondLevelDomainEventRegistrationFee = "registration-fee"
)

// Register top-level-domain event
const (
	EventTypeRegsterTopLevelDomain = "register-top-level-domain"

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
	EventTypeWithdrawRegistrationFees = "withdraw-registration-fees"

	AttributeWithdrawRegistrationFeesEventDomainName = "name"
	AttributeWithdrawRegistrationFeesEventDomainFees = "fees"
)
