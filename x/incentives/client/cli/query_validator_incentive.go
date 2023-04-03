package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mycel/x/incentives/types"
)

func CmdListValidatorIncentive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-validator-incentive",
		Short: "list all validatorIncentive",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllValidatorIncentiveRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ValidatorIncentiveAll(context.Background(), params)
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

func CmdShowValidatorIncentive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-validator-incentive [address]",
		Short: "shows a validatorIncentive",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryGetValidatorIncentiveRequest{
				Address: argAddress,
			}

			res, err := queryClient.ValidatorIncentive(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
