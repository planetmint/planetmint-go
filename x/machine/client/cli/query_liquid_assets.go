package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetLiquidAssetsByMachineid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquid-assets [machine-id]",
		Short: "Query for liquid assets by machine ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqMachineID := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetLiquidAssetsByMachineidRequest{

				MachineID: reqMachineID,
			}

			res, err := queryClient.GetLiquidAssetsByMachineid(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
