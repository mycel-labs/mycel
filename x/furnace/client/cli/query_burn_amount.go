package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/mycel-domain/mycel/x/furnace/types"
	"github.com/spf13/cast"
)

func CmdListBurnAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-burn-amount",
		Short: "list all burnAmount",
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

			params := &types.QueryAllBurnAmountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.BurnAmountAll(cmd.Context(), params)
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

func CmdShowBurnAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-burn-amount [index]",
		Short: "shows a burnAmount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argIndex, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetBurnAmountRequest{
				Index: argIndex,
			}

			res, err := queryClient.BurnAmount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
