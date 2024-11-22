package lib

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sigs.k8s.io/yaml"
)

func readSequenceFromFile(seqFile *os.File) (sequence uint64, err error) {
	scanner := bufio.NewScanner(seqFile)

	// read the first line
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, fmt.Errorf("reading sequence file: %w", err)
		}
		return 0, fmt.Errorf("sequence file empty: %s", seqFile.Name())
	}

	sequenceString := scanner.Text()

	// check for additional lines
	if scanner.Scan() {
		return 0, fmt.Errorf("malformed sequence file %s: contains multiple lines", seqFile.Name())
	}

	sequence, err = strconv.ParseUint(sequenceString, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing sequence number: %w", err)
	}

	return
}

func getSequenceFromChain(clientCtx client.Context) (sequence uint64, err error) {
	account, err := clientCtx.AccountRetriever.GetAccount(clientCtx, clientCtx.FromAddress)
	if err != nil {
		return 0, fmt.Errorf("retrieving account: %w", err)
	}
	return account.GetSequence(), nil
}

func getOrCreateSequenceDirectory() (path string, err error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("getting current user: %w", err)
	}

	path = filepath.Join(GetConfig().RootDir, "sequence")

	// expand tilde to user's home directory
	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(usr.HomeDir, path[2:])
	}

	// create directory if it doesn't exist
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return "", fmt.Errorf("creating sequence directory: %w", err)
	}

	return
}

func openSequenceFile(fromAddress sdk.AccAddress) (file *os.File, err error) {
	path, err := getOrCreateSequenceDirectory()
	if err != nil {
		return
	}

	addrHex := hex.EncodeToString(fromAddress)
	filename := filepath.Join(path, addrHex)

	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("opening sequence file: %w", err)
	}

	return
}

// ParseTxResponse converts output buffer to TxResponse
func ParseTxResponse(out *bytes.Buffer) (txResponse sdk.TxResponse, err error) {
	// Convert string numbers to actual numbers
	numericRegex := regexp.MustCompile(`"([0-9]+?)"`)
	jsonStr := numericRegex.ReplaceAllString(out.String(), "${1}")

	// We might have YAML here, so we need to convert to JSON first, because TxResponse struct lacks `yaml:"height,omitempty"`, etc.
	// Since JSON is a subset of YAML, passing JSON through YAMLToJSON is a no-op and the result is the byte array of the JSON again.
	jsonBytes, err := yaml.YAMLToJSON([]byte(jsonStr))
	if err != nil {
		return sdk.TxResponse{}, fmt.Errorf("converting YAML to JSON: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, &txResponse); err != nil {
		return sdk.TxResponse{}, fmt.Errorf("unmarshaling transaction response: %w", err)
	}

	return
}
