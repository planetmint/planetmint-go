package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/der/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "der [zigbee-id]",
		Short: "Query der",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqZigbeeID := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDerRequest{

				ZigbeeID: reqZigbeeID,
			}

			res, err := queryClient.Der(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
