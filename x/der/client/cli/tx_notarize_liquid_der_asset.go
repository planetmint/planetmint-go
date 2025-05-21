package cli

import (
	"strconv"

	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/planetmint/planetmint-go/x/der/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdNotarizeLiquidDerAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "notarize-liquid-der-asset [der-asset]",
		Short: "Broadcast message notarizeLiquidDerAsset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDerAsset := new(types.LiquidDerAsset)
			err = json.Unmarshal([]byte(args[0]), argDerAsset)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNotarizeLiquidDerAsset(
				clientCtx.GetFromAddress().String(),
				argDerAsset,
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
