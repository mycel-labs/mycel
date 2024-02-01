package cli

import (
	"fmt"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/mycel-domain/mycel/x/registry/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitTopLevelDomainProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-top-level-domain-proposal [name] [registration-period-in-year]",
		Short: "Broadcast message submit-top-level-domain-proposal",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argRegistrationPeriodInYear, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			if argRegistrationPeriodInYear < 1 || argRegistrationPeriodInYear > 4 {
				return errorsmod.Wrapf(types.ErrTopLevelDomainInvalidRegistrationPeriod, "%d year(s)", argRegistrationPeriodInYear)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, err := cli.ReadGovPropFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitTopLevelDomainProposal(
				clientCtx.GetFromAddress().String(),
				argName,
				argRegistrationPeriodInYear,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			if err := proposal.SetMsgs([]sdk.Msg{msg}); err != nil {
				return fmt.Errorf("failed to create submit top-level-domain proposal message: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), proposal)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
