package registry

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/mycel-domain/mycel/api/mycel/registry/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "TopLevelDomain",
					Use:       "top-level-domain [name]",
					Short:     "Query information about a top-level domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
					},
				},

				{
					RpcMethod: "TopLevelDomainAll",
					Use:       "top-level-domain-all",
					Short:     "Query all top-level domains",
				},

				{
					RpcMethod: "SecondLevelDomain",
					Use:       "second-level-domain [name] [parent]",
					Short:     "Query information about a second-level domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "parent"},
					},
				},

				{
					RpcMethod: "SecondLevelDomainAll",
					Use:       "second-level-domain-all",
					Short:     "Query all second-level domains",
				},

				{
					RpcMethod: "DomainOwnership",
					Use:       "domain-ownership [owner]",
					Short:     "Query domain ownership information",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "owner"},
					},
				},

				{
					RpcMethod: "DomainOwnershipAll",
					Use:       "domain-ownership-all",
					Short:     "Query all domain ownership information",
				},

				{
					RpcMethod: "DomainRegistrationFee",
					Use:       "domain-registration-fee [name] [parent] [registration-period-in-year] [registerer]",
					Short:     "Query domain registration fee",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "parent"},
						{ProtoField: "registration_period_in_year"},
						{ProtoField: "registerer"},
					},
				},
				{
					RpcMethod: "Role",
					Use:       "role [domain-name] [address]",
					Short:     "Query the role of an address in a domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "domain_name"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "WalletRecord",
					Use:       "wallet-record [domain-name] [domain-parent] [wallet-record-type]",
					Short:     "Query wallet record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "domain_name"},
						{ProtoField: "domain_parent"},
						{ProtoField: "wallet_record_type"},
					},
				},
				{
					RpcMethod: "DnsRecord",
					Use:       "dns-record [domain-name] [domain-parent] [dns-record-type]",
					Short:     "Query DNS record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "domain_name"},
						{ProtoField: "domain_parent"},
						{ProtoField: "dns_record_type"},
					},
				},
				{
					RpcMethod: "AllRecords",
					Use:       "all-records [domain-name] [domain-parent]",
					Short:     "Query all records",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "domain_name"},
						{ProtoField: "domain_parent"},
					},
				},
				{
					RpcMethod: "TextRecord",
					Use:       "text-record [domain-name] [domain-parent] [key]",
					Short:     "Query text record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "domain_name"},
						{ProtoField: "domain_parent"},
						{ProtoField: "key"},
					},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateWalletRecord",
					Use:       "update-wallet-record [creator] [name] [parent] [wallet-record-type] [value]",
					Short:     "Update a wallet record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "creator"},
						{ProtoField: "name"},
						{ProtoField: "parent"},
						{ProtoField: "wallet_record_type"},
						{ProtoField: "value"},
					},
				},
				{
					RpcMethod: "UpdateDnsRecord",
					Use:       "update-dns-record [creator] [name] [parent] [dns-record-type] [value]",
					Short:     "Update a DNS record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "parent"},
						{ProtoField: "dns_record_type"},
						{ProtoField: "value"},
					},
				},
				{
					RpcMethod: "RegisterSecondLevelDomain",
					Use:       "register-second-level-domain [creator] [name] [parent] [registration-period-in-year]",
					Short:     "Register a second-level domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "parent"},
						{ProtoField: "registration_period_in_year"},
					},
				},
				{
					RpcMethod: "RegisterTopLevelDomain",
					Use:       "register-top-level-domain [creator] [name] [registration-period-in-year]",
					Short:     "Register a top-level domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "registration_period_in_year"},
					},
				},
				{
					RpcMethod: "WithdrawRegistrationFee",
					Use:       "withdraw-registration-fee [creator] [name]",
					Short:     "Withdraw registration fee",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
					},
				},
				{
					RpcMethod: "ExtendTopLevelDomainExpirationDate",
					Use:       "extend-top-level-expiration-date [creator] [name] [extension-period-in-year]",
					Short:     "Extend the expiration date of a top-level domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "extension_period_in_year"},
					},
				},
				{
					RpcMethod: "UpdateTextRecord",
					Use:       "update-text-record [creator] [name] [parent] [key] [value]",
					Short:     "Update a text record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "parent"},
						{ProtoField: "key"},
						{ProtoField: "value"},
					},
				},
				{
					RpcMethod: "UpdateTopLevelDomainRegistrationPolicy",
					Use:       "update-top-level-domain-registration-policy [creator] [name] [registration-policy]",
					Short:     "Update top-level domain registration policy",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "registration_policy"},
					},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
