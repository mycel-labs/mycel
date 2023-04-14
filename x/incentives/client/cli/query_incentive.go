package cli

import (
	"context"

	"github.com/mycel-domain/mycel/x/incentives/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListIncentive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-incentive",
		Short: "list all incentive",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllIncentiveRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.IncentiveAll(context.Background(), params)
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

func CmdShowIncentive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-incentive [epoch]",
		Short: "shows a incentive",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argEpoch, err := cast.ToInt64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetIncentiveRequest{
				Epoch: argEpoch,
			}

			res, err := queryClient.Incentive(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
