package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/planetmint/planetmint-go/x/dao/types"
)

func GetCmdListRedeemClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-redeem-claim",
		Short: "list all redeem-claim",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRedeemClaimRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RedeemClaimAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdShowRedeemClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-redeem-claim [beneficiary] [id]",
		Short: "shows a redeem-claim",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argBeneficiary := args[0]
			argID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetRedeemClaimRequest{
				Beneficiary: argBeneficiary,
				Id:          argID,
			}

			res, err := queryClient.RedeemClaim(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

var _ = strconv.Itoa(0)

func GetCmdRedeemClaimByLiquidTxHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem-claim-by-liquid-tx-hash [liquid-tx-hash]",
		Short: "Query redeem-claim-by-liquid-tx-hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqLiquidTxHash := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryRedeemClaimByLiquidTxHashRequest{

				LiquidTxHash: reqLiquidTxHash,
			}

			res, err := queryClient.RedeemClaimByLiquidTxHash(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
