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

func CmdReissueRDDLProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reissue-rddl-proposal [proposer] [tx] [blockheight]",
		Short: "Broadcast message reissueRDDLProposal",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProposer := args[0]
			argTx := args[1]
			argBlockHeight, err := cast.ToInt64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgReissueRDDLProposal(
				clientCtx.GetFromAddress().String(),
				argProposer,
				argTx,
				argBlockHeight,
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
