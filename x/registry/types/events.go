package types

// Register domain event
const (
	EventTypeRegsterDomain = "register-domain"

	AttributeRegisterDomainEventName           = "name"
	AttributeRegisterDomainEventParent         = "parent"
	AttributeRegisterDomainEventExpirationDate = "expiration-date"
	AttributeRegisterDomainLevel               = "domain-level"
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
