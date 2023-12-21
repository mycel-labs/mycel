package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/mycel-domain/mycel/x/resolver/types"
)

var _ = strconv.Itoa(0)

func CmdTextRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "text-record [domain-name] [domain-parent] [key]",
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
