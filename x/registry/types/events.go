package types

// Register domain event
const (
	EventTypeRegsterDomain = "register_domain"

	AttributeRegisterDomainEventName                     = "name"
	AttributeRegisterDomainEventParent                   = "parent"
	AttributeRegisterDomainEventRegistrationPeriodInYear = "regstration-period-in-year"
	AttributeRegisterDomainEventExpireationDate          = "expireation-date"
	AttributeRegisterDomainLevel                         = "domain-level"
)

// Update wallet record event
const (
	EventTypeUpdateWalletRecord = "update_wallet_record"

	AttributeUpdateWalletRecordEventWalletRecordType = "wallet-record-type"
	AttributeUpdateWalletRecordEventValue            = "value"
)
