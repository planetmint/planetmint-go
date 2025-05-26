package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/der/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft [zigbee-id]",
		Short: "Query nft",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqZigbeeID := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryNftRequest{

				ZigbeeID: reqZigbeeID,
			}

			res, err := queryClient.Nft(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
