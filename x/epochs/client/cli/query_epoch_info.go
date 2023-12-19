package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/mycel-domain/mycel/x/epochs/types"
)

func CmdListEpochInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-epoch-info",
		Short: "list all epochInfo",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEpochInfoRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EpochInfoAll(context.Background(), params)
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

func CmdShowEpochInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-epoch-info [identifier]",
		Short: "shows a epochInfo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIdentifier := args[0]

			params := &types.QueryGetEpochInfoRequest{
				Identifier: argIdentifier,
			}

			res, err := queryClient.EpochInfo(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
