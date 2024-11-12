package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func GetCmdDistribution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "distribution [height]",
		Short: "Query for distributions by height",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			requestHeight, err := cast.ToInt64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			latestHeight, err := rpc.GetChainHeight(clientCtx)
			if err != nil {
				return err
			}

			if requestHeight > latestHeight {
				err = fmt.Errorf("height %d must be less than or equal to the current blockchain height %d",
					requestHeight, latestHeight)
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetDistributionRequest{
				Height: requestHeight,
			}

			res, err := queryClient.GetDistribution(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
