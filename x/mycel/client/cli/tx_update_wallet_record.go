package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"mycel/x/mycel/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateWalletRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-wallet-record [name] [parent] [wallet-record-type]",
		Short: "Broadcast message updateWalletRecord",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argParent := args[1]
			argWalletRecordType := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateWalletRecord(
				clientCtx.GetFromAddress().String(),
				argName,
				argParent,
				argWalletRecordType,
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
