package cli

import (
	"context"
	"planetmint-go/testutil"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// ExecTestCLICmd builds the client context, mocks the output and executes the command.
func ExecTestCLICmd(clientCtx client.Context, cmd *cobra.Command, extraArgs []string) (testutil.BufferWriter, error) {
	cmd.SetArgs(extraArgs)

	_, out := testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	if err := cmd.ExecuteContext(ctx); err != nil {
		return out, err
	}

	return out, nil
}
