package cli

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/planetmint/planetmint-go/x/dao/types"
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

	cmd.AddCommand(CmdReportPopResult())
	cmd.AddCommand(CmdReissueRDDLProposal())
	cmd.AddCommand(CmdMintToken())
	cmd.AddCommand(CmdReissueRDDLResult())
	cmd.AddCommand(CmdDistributionResult())
	cmd.AddCommand(CmdDistributionRequest())
	cmd.AddCommand(CmdUpdateParams())
	cmd.AddCommand(CmdInitPop())
	cmd.AddCommand(CmdCreateRedeemClaim())
	cmd.AddCommand(CmdUpdateRedeemClaim())
	cmd.AddCommand(CmdConfirmRedeemClaim())
	// this line is used by starport scaffolding # 1

	return cmd
}
