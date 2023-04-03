package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mycel/x/incentives/types"
)

func CmdListDelegetorIncentive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-delegetor-incentive",
		Short: "list all delegetorIncentive",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDelegetorIncentiveRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DelegetorIncentiveAll(context.Background(), params)
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

func CmdShowDelegetorIncentive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-delegetor-incentive [address]",
		Short: "shows a delegetorIncentive",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryGetDelegetorIncentiveRequest{
				Address: argAddress,
			}

			res, err := queryClient.DelegetorIncentive(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
