package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func GetCmdTrustAnchor() *cobra.Command {
	// Group trust-anchor queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "trust-anchor",
		Short:                      "Query for trust anchor subcommand",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdGetStatus())

	return cmd
}

func CmdGetStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status [machine-id]",
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
