package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/mycel-domain/mycel/x/registry/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateDnsRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-dns-record [name] [parent] [dns-record-type] [value]",
		Short: "Broadcast message updateDnsRecord",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argParent := args[1]
			argDnsRecordType := args[2]
			argValue := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDnsRecord(
				clientCtx.GetFromAddress().String(),
				argName,
				argParent,
				argDnsRecordType,
				argValue,
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
