package cli

import (
	"strconv"

	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMintToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-token [mint-request]",
		Short: "Broadcast message mint-token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMintRequest := new(types.MintRequest)
			err = json.Unmarshal([]byte(args[0]), argMintRequest)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintToken(
				clientCtx.GetFromAddress().String(),
				argMintRequest,
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
