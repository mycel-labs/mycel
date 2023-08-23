package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/mycel-domain/mycel/x/registry/types"
	"github.com/spf13/cobra"
)

func CmdListSecondLevelDomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-second-level-domain",
		Short: "list all secondLevelDomains",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSecondLevelDomainRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SecondLevelDomainAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowSecondLevelDomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-second-level-domain [name] [parent]",
		Short: "shows a secondLevelDomain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argName := args[0]
			argParent := args[1]

			params := &types.QueryGetSecondLevelDomainRequest{
				Name:   argName,
				Parent: argParent,
			}

			res, err := queryClient.SecondLevelDomain(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
