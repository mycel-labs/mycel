package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/mycel-domain/mycel/x/resolver/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQueryWalletRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-wallet-record [domain-name] [network-name]",
		Short: "Query queryWalletRecord",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDomainName := args[0]
			reqNetworkName := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQueryWalletRecordRequest{

				DomainName:  reqDomainName,
				NetworkName: reqNetworkName,
			}

			res, err := queryClient.QueryWalletRecord(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
