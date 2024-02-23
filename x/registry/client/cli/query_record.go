package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/mycel-domain/mycel/x/registry/types"
)

var _ = strconv.Itoa(0)

func CmdAllRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-records [name] [parent]",
		Short: "Query allRecords",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqName := args[0]
			reqParent := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRecordsRequest{
				DomainName:   reqName,
				DomainParent: reqParent,
			}

			res, err := queryClient.AllRecords(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdDnsRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dns-record [name] [parent] [dns-record-type]",
		Short: "Query DNS record",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDomainName := args[0]
			reqDomainParent := args[1]
			reqDnsRecordType := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDnsRecordRequest{
				DomainName:    reqDomainName,
				DomainParent:  reqDomainParent,
				DnsRecordType: reqDnsRecordType,
			}

			res, err := queryClient.DnsRecord(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdTextRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "text-record [name] [parent] [key]",
		Short: "Query textRecord",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDomainName := args[0]
			reqDomainParent := args[1]
			reqKey := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTextRecordRequest{
				DomainName:   reqDomainName,
				DomainParent: reqDomainParent,
				Key:          reqKey,
			}

			res, err := queryClient.TextRecord(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryWalletRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wallet-record [name] [parent] [wallet-record-type]",
		Short: "Query wallet record",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDomainName := args[0]
			reqDomainParent := args[1]
			reqWalletRecordType := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryWalletRecordRequest{
				DomainName:       reqDomainName,
				DomainParent:     reqDomainParent,
				WalletRecordType: reqWalletRecordType,
			}

			res, err := queryClient.WalletRecord(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
