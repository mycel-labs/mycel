package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mycel-domain/mycel/x/registry/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetRegistrationFees() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-registration-fees [domain] [fees-by-name] [fees-by-length] [default-fees]",
		Short: "Broadcast message set-registration-fees",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDomain := args[0]
			argFeesByName := args[1]
			argFeesByLength := args[2]
			argDefaultFees := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetRegistrationFees(
				clientCtx.GetFromAddress().String(),
				argDomain,
				argFeesByName,
				argFeesByLength,
				argDefaultFees,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
