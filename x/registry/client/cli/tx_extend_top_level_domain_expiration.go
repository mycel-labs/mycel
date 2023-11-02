package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mycel-domain/mycel/x/registry/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdExtendTopLevelDomainExpiration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extend-top-level-domain-expiration [name] [registration-period-in-year]",
		Short: "Broadcast message extendTopLevelDomainExpiration",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argRegistrationPeriodInYear, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgExtendTopLevelDomainExpiration(
				clientCtx.GetFromAddress().String(),
				argName,
				argRegistrationPeriodInYear,
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
