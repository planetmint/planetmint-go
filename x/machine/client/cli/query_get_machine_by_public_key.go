package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"planetmint-go/x/machine/types"
)

var _ = strconv.Itoa(0)

func CmdGetMachineByPublicKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-machine-by-public-key [public-key]",
		Short: "Query get-machine-by-public-key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPublicKey := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetMachineByPublicKeyRequest{

				PublicKey: reqPublicKey,
			}

			res, err := queryClient.GetMachineByPublicKey(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
