package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateRedeemClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-redeem-claim [beneficiary] [liquid-tx-hash] [amount] [issued]",
		Short: "Create a new redeem-claim",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexBeneficiary := args[0]
			indexLiquidTxHash := args[1]

			// Get value arguments
			argAmount := args[2]
			argIssued, err := cast.ToBoolE(args[3])
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
				indexLiquidTxHash,
				argAmount,
				argIssued,
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
		Use:   "update-redeem-claim [beneficiary] [liquid-tx-hash] [amount] [issued]",
		Short: "Update a redeem-claim",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexBeneficiary := args[0]
			indexLiquidTxHash := args[1]

			// Get value arguments
			argAmount := args[2]
			argIssued, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRedeemClaim(
				clientCtx.GetFromAddress().String(),
				indexBeneficiary,
				indexLiquidTxHash,
				argAmount,
				argIssued,
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

func CmdDeleteRedeemClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-redeem-claim [beneficiary] [liquid-tx-hash]",
		Short: "Delete a redeem-claim",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexBeneficiary := args[0]
			indexLiquidTxHash := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRedeemClaim(
				clientCtx.GetFromAddress().String(),
				indexBeneficiary,
				indexLiquidTxHash,
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
