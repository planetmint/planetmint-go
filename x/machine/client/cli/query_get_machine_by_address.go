package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetMachineByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-machine-by-address [address]",
		Short: "Query get-machine-by-address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetMachineByAddressRequest{

				Address: reqAddress,
			}

			res, err := queryClient.GetMachineByAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
