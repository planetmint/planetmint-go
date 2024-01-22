package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/spf13/cobra"
)

func CmdCreateRedeemClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-redeem-claim [beneficiary] [amount]",
		Short: "Create a new redeem-claim",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexBeneficiary := args[0]

			// Get value arguments
			argAmount, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRedeemClaim(
				clientCtx.GetFromAddress().String(),
				indexBeneficiary,
				argAmount,
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

func CmdUpdateRedeemClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-redeem-claim [beneficiary] [id] [liquid-tx-hash]",
		Short: "Update a redeem-claim",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexBeneficiary := args[0]
			indexId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			// Get value arguments
			argLiquidTxHash := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRedeemClaim(
				clientCtx.GetFromAddress().String(),
				indexBeneficiary,
				argLiquidTxHash,
				indexId,
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
