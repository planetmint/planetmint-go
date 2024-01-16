package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetTrustAnchorStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trust-anchor-status [machine-id]",
		Short: "Query for trust anchor status by machine ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqMachineid := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetTrustAnchorStatusRequest{

				Machineid: reqMachineid,
			}

			res, err := queryClient.GetTrustAnchorStatus(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
