package types

// Register top-level-domain event
const (
	EventTypeRegisterTopLevelDomain = "register-top-level-domain"

	AttributeRegisterTopLevelDomainEventName                      = "name"
	AttributeRegisterTopLevelDomainEventExpirationDate            = "expiration-date"
	AttributeRegisterTopLevelDomainEventMaxSubdomainRegistrations = "max-subdomain-registrations"
	AttributeRegisterTopLevelDomainEventTotalRegistrationFee      = "total-registration-fee"
	AttributeRegisterTopLevelDomainEventBurnWeight                = "burn-weight"
	AttributeRegisterTopLevelDomainEventRegistrationFeeToBurn     = "registration-fee-to-burn"
	AttributeRegisterTopLevelDomainEventRegistrationFeeToTreasury = "registration-fee-to-treasury"
)

// Register second-level-domain event
const (
	EventTypeRegisterSecondLevelDomain = "register-second-leve-domain"

	AttributeRegisterSecondLevelDomainEventName            = "name"
	AttributeRegisterSecondLevelDomainEventParent          = "parent"
	AttributeRegisterSecondLevelDomainEventExpirationDate  = "expiration-date"
	AttributeRegisterSecondLevelDomainEventRegistrationFee = "registration-fee"
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

// Extend top-level-domain expiration date event
const (
	EventTypeExtendTopLevelDomainExpirationDate = "extend-top-level-domain-expiration-date"

	AttributeExtendTopLevelDomainExpirationDateEventDomainName                = "name"
	AttributeExtendTopLevelDomainExpirationDateEventExpirationDate            = "expiration-date"
	AttributeExtendTopLevelDomainExpirationDateEventTotalRegistrationFee      = "total-registration-fee"
	AttributeExtendTopLevelDomainExpirationDateEventBurnWeight                = "burn-weight"
	AttributeExtendTopLevelDomainExpirationDateEventRegistrationFeeToBurn     = "registration-fee-to-burn"
	AttributeExtendTopLevelDomainExpirationDateEventRegistrationFeeToTreasury = "registration-fee-to-treasury"
)
