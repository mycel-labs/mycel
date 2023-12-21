package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func CmdListDomainOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-domain-ownership",
		Short: "list all domainOwnership",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDomainOwnershipRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DomainOwnershipAll(context.Background(), params)
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

func CmdShowDomainOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-domain-ownership [owner]",
		Short: "shows a domainOwnership",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argOwner := args[0]

			params := &types.QueryGetDomainOwnershipRequest{
				Owner: argOwner,
			}

			res, err := queryClient.DomainOwnership(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
