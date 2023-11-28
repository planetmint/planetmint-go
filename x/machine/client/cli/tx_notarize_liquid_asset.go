package cli

import (
	"strconv"

	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdNotarizeLiquidAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "notarize-liquid-asset [notarization]",
		Short: "Broadcast message NotarizeLiquidAsset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNotarization := new(types.LiquidAsset)
			err = json.Unmarshal([]byte(args[0]), argNotarization)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNotarizeLiquidAsset(
				clientCtx.GetFromAddress().String(),
				argNotarization,
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
