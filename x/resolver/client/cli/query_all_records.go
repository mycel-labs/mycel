package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/mycel-domain/mycel/x/resolver/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAllRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-records [domainName] [domainParent]",
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
