package cli

import (
	"strconv"

	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"planetmint-go/x/machine/types"
)

var _ = strconv.Itoa(0)

func CmdAttestMachine() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attest-machine [machine]",
		Short: "Broadcast message attest-machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMachine := new(types.Machine)
			err = json.Unmarshal([]byte(args[0]), argMachine)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAttestMachine(
				clientCtx.GetFromAddress().String(),
				argMachine,
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
