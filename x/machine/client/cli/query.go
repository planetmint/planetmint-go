package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/planetmint/planetmint-go/x/machine/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group machine queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdActivatedTrustAnchorCount(),
		GetCmdActiveTrustAnchorCount(),
		GetCmdLiquidAssetsByMachineid(),
		GetCmdMachineByAddress(),
		GetCmdMachineByPublicKey(),
		GetCmdQueryParams(),
		GetCmdTrustAnchor(),
	)

	// this line is used by starport scaffolding # 1

	return cmd
}
