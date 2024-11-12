package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/planetmint/planetmint-go/x/dao/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group dao queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdChallenge(),
		GetCmdChallenges(),
		GetCmdDistribution(),
		GetCmdListRedeemClaim(),
		GetCmdMintRequests(),
		GetCmdQueryParams(),
		GetCmdRedeemClaimByLiquidTxHash(),
		GetCmdReissuance(),
		GetCmdReissuances(),
		GetCmdShowRedeemClaim(),
	)

	// this line is used by starport scaffolding # 1

	return cmd
}
