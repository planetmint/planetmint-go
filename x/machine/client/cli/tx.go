package cli

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      types.ModuleName + " transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdAttestMachine())
	cmd.AddCommand(CmdRegisterTrustAnchor())
	cmd.AddCommand(CmdNotarizeLiquidAsset())
	cmd.AddCommand(CmdUpdateParams())
	cmd.AddCommand(CmdMintProduction())
	cmd.AddCommand(CmdBurnConsumption())
	// this line is used by starport scaffolding # 1

	return cmd
}
