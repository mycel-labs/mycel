package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func CmdListTopLevelDomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-top-level-domain",
		Short: "list all topLevelDomain",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTopLevelDomainRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TopLevelDomainAll(cmd.Context(), params)
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

func CmdShowTopLevelDomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-top-level-domain [name]",
		Short: "shows a topLevelDomain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argName := args[0]

			params := &types.QueryGetTopLevelDomainRequest{
				Name: argName,
			}

			res, err := queryClient.TopLevelDomain(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
