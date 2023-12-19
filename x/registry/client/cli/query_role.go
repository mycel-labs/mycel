package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/mycel-domain/mycel/x/registry/types"
)

var _ = strconv.Itoa(0)

func CmdRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "role [domain-name] [address]",
		Short: "Query role",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDomainName := args[0]
			reqAddress := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryRoleRequest{
				DomainName: reqDomainName,
				Address:    reqAddress,
			}

			res, err := queryClient.Role(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
