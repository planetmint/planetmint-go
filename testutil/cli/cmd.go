package cli

import (
	"context"
	"encoding/json"
	"planetmint-go/testutil"
	"regexp"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
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

// GetTxResponseFromOut converts strings to numbers and unmarshalles out into TxResponse struct
func GetTxResponseFromOut(out testutil.BufferWriter) (sdk.TxResponse, error) {
	var txResponse sdk.TxResponse

	m := regexp.MustCompile(`"([0-9]+?)"`)
	str := m.ReplaceAllString(out.String(), "${1}")

	// We might have YAML here, so we need to convert to JSON first, because TxResponse struct lacks `yaml:"height,omitempty"`, etc.
	// Since JSON is a subset of YAML, passing JSON through YAMLToJSON is a no-op and the result is the byte array of the JSON again.
	j, err := yaml.YAMLToJSON([]byte(str))
	if err != nil {
		return txResponse, err
	}

	err = json.Unmarshal(j, &txResponse)
	if err != nil {
		return txResponse, err
	}

	return txResponse, nil
}

// GetRawLogFromTxResponse queries the TxHash of txResponse from the chain and returns the RawLog from the answer.
func GetRawLogFromTxResponse(val *network.Validator, txResponse sdk.TxResponse) (string, error) {
	args := []string{
		txResponse.TxHash,
	}

	out, err := ExecTestCLICmd(val.ClientCtx, authcmd.QueryTxCmd(), args)
	if err != nil {
		return "", err
	}

	txRes, err := GetTxResponseFromOut(out)
	if err != nil {
		return "", err
	}

	return txRes.RawLog, nil
}
