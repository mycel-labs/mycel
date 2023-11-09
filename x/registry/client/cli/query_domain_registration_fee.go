package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/mycel-domain/mycel/x/registry/types"
)

var _ = strconv.Itoa(0)

func CmdDomainRegistrationFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "domain-registration-fee [name] [parent] [registration-period-in-year]",
		Short: "query domain registration fee",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqName := args[0]
			reqParent := args[1]
			reqRegistrationPeriodInYear, err := cast.ToUint64E(args[2])

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDomainRegistrationFeeRequest{
				Name:                     reqName,
				Parent:                   reqParent,
				RegistrationPeriodInYear: reqRegistrationPeriodInYear,
			}

			res, err := queryClient.DomainRegistrationFee(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
