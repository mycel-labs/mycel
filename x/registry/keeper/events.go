package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/x/registry/types"
)

// Register top-level-domain event
func EmitRegisterTopLevelDomainEvent(ctx sdk.Context, domain types.TopLevelDomain, fee types.TopLevelDomainFee) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRegisterTopLevelDomain,
			sdk.NewAttribute(types.AttributeRegisterTopLevelDomainEventName, domain.Name),
			sdk.NewAttribute(types.AttributeRegisterTopLevelDomainEventExpirationDate, domain.ExpirationDate.String()),
			sdk.NewAttribute(types.AttributeRegisterTopLevelDomainEventMaxSubdomainRegistrations, fmt.Sprintf("%d", domain.SubdomainConfig.MaxSubdomainRegistrations)),
			sdk.NewAttribute(types.AttributeRegisterTopLevelDomainEventTotalRegistrationFee, fee.TotalFee.String()),
			sdk.NewAttribute(types.AttributeRegisterTopLevelDomainEventBurnWeight, fee.BurnWeight),
			sdk.NewAttribute(types.AttributeRegisterTopLevelDomainEventRegistrationFeeToBurn, fee.FeeToBurn.String()),
			sdk.NewAttribute(types.AttributeRegisterTopLevelDomainEventRegistrationFeeToTreasury, fee.FeeToTreasury.String()),
		),
	)
}

// Register second-level-domain event
func EmitRegisterSecondLevelDomainEvent(ctx sdk.Context, domain types.SecondLevelDomain, fee sdk.Coin) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeRegisterSecondLevelDomain,
			sdk.NewAttribute(types.AttributeRegisterSecondLevelDomainEventName, domain.Name),
			sdk.NewAttribute(types.AttributeRegisterSecondLevelDomainEventParent, domain.Parent),
			sdk.NewAttribute(types.AttributeRegisterSecondLevelDomainEventExpirationDate, domain.ExpirationDate.String()),
			sdk.NewAttribute(types.AttributeRegisterSecondLevelDomainEventRegistrationFee, fee.String()),
		),
	)
}

// Update wallet record event
func EmitUpdateWalletRecordEvent(ctx sdk.Context, msg types.MsgUpdateWalletRecord) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeUpdateWalletRecord,
			sdk.NewAttribute(types.AttributeUpdateWalletRecordEventDomainName, msg.Name),
			sdk.NewAttribute(types.AttributeUpdateWalletRecordEventDomainParent, msg.Parent),
			sdk.NewAttribute(types.AttributeUpdateWalletRecordEventWalletRecordType, msg.WalletRecordType),
			sdk.NewAttribute(types.AttributeUpdateWalletRecordEventValue, msg.Value),
		),
	)
}

// Update dns record event
func EmitUpdateDnsRecordEvent(ctx sdk.Context, msg types.MsgUpdateDnsRecord) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeUpdateDnsRecord,
			sdk.NewAttribute(types.AttributeUpdateDnsRecordEventDomainName, msg.Name),
			sdk.NewAttribute(types.AttributeUpdateDnsRecordEventDomainParent, msg.Parent),
			sdk.NewAttribute(types.AttributeUpdateDnsRecordEventDnsRecordType, msg.DnsRecordType),
			sdk.NewAttribute(types.AttributeUpdateDnsRecordEventValue, msg.Value),
		),
	)
}

// Update dns record event
func EmitUpdateTextRecordEvent(ctx sdk.Context, msg types.MsgUpdateTextRecord) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeUpdateTextRecord,
			sdk.NewAttribute(types.AttributeUpdateTextRecordEventDomainName, msg.Name),
			sdk.NewAttribute(types.AttributeUpdateTextRecordEventDomainParent, msg.Parent),
			sdk.NewAttribute(types.AttributeUpdateTextRecordEventKey, msg.Key),
			sdk.NewAttribute(types.AttributeUpdateTextRecordEventValue, msg.Value),
		),
	)
}

// Withdraw fees event
func EmitWithdrawRegistrationFeeEvent(ctx sdk.Context, msg types.MsgWithdrawRegistrationFee, fee sdk.Coins) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeWithdrawRegistrationFee,
			sdk.NewAttribute(types.AttributeWithdrawRegistrationFeeEventDomainName, msg.Name),
			sdk.NewAttribute(types.AttributeWithdrawRegistrationFeeEventDomainFee, fee.String()),
		),
	)
}

// Extend top-level-domain expiration date event
func EmitExtendTopLevelDomainExpirationDateEvent(ctx sdk.Context, domain types.TopLevelDomain, fee types.TopLevelDomainFee) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeExtendTopLevelDomainExpirationDate,
			sdk.NewAttribute(types.AttributeExtendTopLevelDomainExpirationDateEventDomainName, domain.Name),
			sdk.NewAttribute(types.AttributeExtendTopLevelDomainExpirationDateEventExpirationDate, domain.ExpirationDate.String()),
			sdk.NewAttribute(types.AttributeExtendTopLevelDomainExpirationDateEventTotalRegistrationFee, fee.TotalFee.String()),
			sdk.NewAttribute(types.AttributeExtendTopLevelDomainExpirationDateEventBurnWeight, fee.BurnWeight),
			sdk.NewAttribute(types.AttributeExtendTopLevelDomainExpirationDateEventRegistrationFeeToBurn, fee.FeeToBurn.String()),
			sdk.NewAttribute(types.AttributeExtendTopLevelDomainExpirationDateEventRegistrationFeeToTreasury, fee.FeeToTreasury.String()),
		),
	)
}

// Update top-level-domain registration policy
func EmitUpdateTopLevelDomainRegistrationPolicyEvent(ctx sdk.Context, domain types.TopLevelDomain) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeUpdateTopLevelDomainRegistrationPolicy,
			sdk.NewAttribute(types.AttributeUpdateTopLevelDomainRegistrationPolicyEventDomainName, domain.Name),
			sdk.NewAttribute(types.AttributeUpdateTopLevelDomainRegistrationPolicyEventRegistrationPolicy, domain.SubdomainConfig.RegistrationPolicy.String()),
		),
	)
}
