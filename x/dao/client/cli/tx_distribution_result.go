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

func CmdDistributionResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "distribution-result [last-pop] [dao-txid] [investor-txid] [pop-txid] [early-investor-txid] [strategic-txid]",
		Short: "Broadcast message DistributionResult",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argLastPop, err := cast.ToInt64E(args[0])
			if err != nil {
				return err
			}
			argDaoTxid := args[1]
			argInvestorTxid := args[2]
			argPopTxid := args[3]
			argEarlyInvestorTxid := args[4]
			argStrategicTxid := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDistributionResult(
				clientCtx.GetFromAddress().String(),
				argLastPop,
				argDaoTxid,
				argInvestorTxid,
				argPopTxid,
				argEarlyInvestorTxid,
				argStrategicTxid,
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
