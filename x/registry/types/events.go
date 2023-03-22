package types

// Register domain event
const (
	EventTypeRegsterDomain = "register-domain"

	AttributeRegisterDomainEventName                     = "name"
	AttributeRegisterDomainEventParent                   = "parent"
	AttributeRegisterDomainEventRegistrationPeriodInYear = "regstration-period-in-year"
	AttributeRegisterDomainEventExpirationDate           = "expiration-date"
	AttributeRegisterDomainLevel                         = "domain-level"
)

// Update wallet record event
const (
	EventTypeUpdateWalletRecord = "update-wallet-record"

	AttributeUpdateWalletRecordEventDomainName       = "name"
	AttributeUpdateWalletRecordEventDomainParent     = "parent"
	AttributeUpdateWalletRecordEventWalletRecordType = "wallet-record-type"
	AttributeUpdateWalletRecordEventValue            = "value"
)
