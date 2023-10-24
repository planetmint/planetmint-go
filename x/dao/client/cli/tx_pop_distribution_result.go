package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPopDistributionResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pop-distribution-result [last-pop] [dao-tx] [investor-tx] [pop-tx]",
		Short: "Broadcast message PopDistributionResult",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argLastPop, err := cast.ToInt64E(args[0])
			if err != nil {
				return err
			}
			argDaoTx := args[1]
			argInvestorTx := args[2]
			argPopTx := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPopDistributionResult(
				clientCtx.GetFromAddress().String(),
				argLastPop,
				argDaoTx,
				argInvestorTx,
				argPopTx,
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
