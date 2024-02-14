package cli

import (
	"bytes"
	"context"
	"errors"

	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/testutil"

	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/planetmint/planetmint-go/testutil/network"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// ExecTestCLICmd builds the client context, mocks the output and executes the command.
func ExecTestCLICmd(clientCtx client.Context, cmd *cobra.Command, extraArgs []string) (out testutil.BufferWriter, err error) {
	cmd.SetArgs(extraArgs)

	_, out = testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	if err = cmd.ExecuteContext(ctx); err != nil {
		return
	}

	output, ok := out.(*bytes.Buffer)
	if !ok {
		err = lib.ErrTypeAssertionFailed
		return
	}

	txResponse, err := lib.GetTxResponseFromOut(output)
	if err != nil {
		return
	}

	if txResponse.Code != 0 {
		err = errors.New(txResponse.RawLog)
		return
	}
	return
}

// GetRawLogFromTxOut queries the TxHash of out from the chain and returns the RawLog from the answer.
func GetRawLogFromTxOut(val *network.Validator, out *bytes.Buffer) (rawLog string, err error) {
	txResponse, err := lib.GetTxResponseFromOut(out)
	if err != nil {
		return
	}
	if txResponse.Code != 0 {
		err = errors.New(txResponse.RawLog)
		return
	}
	args := []string{
		txResponse.TxHash,
	}

	output, err := ExecTestCLICmd(val.ClientCtx, authcmd.QueryTxCmd(), args)
	if err != nil {
		return
	}

	out, ok := output.(*bytes.Buffer)
	if !ok {
		err = lib.ErrTypeAssertionFailed
		return
	}

	txRes, err := lib.GetTxResponseFromOut(out)
	if err != nil {
		return
	}
	rawLog = txRes.RawLog
	return
}
