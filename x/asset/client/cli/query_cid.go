package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/asset/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetByCID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cid [cid]",
		Short: "Query for assets by CID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqCid := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetNotarizedAssetRequest{

				Cid: reqCid,
			}

			res, err := queryClient.GetNotarizedAsset(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
