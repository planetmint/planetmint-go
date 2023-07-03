package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"planetmint-go/x/asset/types"
)

var _ = strconv.Itoa(0)

func CmdNotarizeAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "notarize-asset [hash] [signature] [pub-key]",
		Short: "Broadcast message notarize-asset",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHash := args[0]
			argSignature := args[1]
			argPubKey := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNotarizeAsset(
				clientCtx.GetFromAddress().String(),
				argHash,
				argSignature,
				argPubKey,
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
