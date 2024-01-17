package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetMintRequests() *cobra.Command {
	// Group mint-requests queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "mint-requests",
		Short:                      "Query for mint requests subcommand",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdGetAddress())
	cmd.AddCommand(CmdGetHash())

	return cmd
}

func CmdGetAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address [address]",
		Short: "Query for mint requests by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryMintRequestsByAddressRequest{

				Address: reqAddress,
			}

			res, err := queryClient.MintRequestsByAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash [hash]",
		Short: "Query for mint requests by hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqHash := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetMintRequestsByHashRequest{

				Hash: reqHash,
			}

			res, err := queryClient.GetMintRequestsByHash(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
