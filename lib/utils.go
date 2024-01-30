package lib

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
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

func getSequenceFromFile(seqFile *os.File) (sequence uint64, err error) {
	var sequenceString string
	lineCount := int64(0)
	scanner := bufio.NewScanner(seqFile)
	for scanner.Scan() {
		sequenceString = scanner.Text()
		lineCount++
	}
	err = scanner.Err()
	if err != nil {
		return
	}
	if lineCount == 0 {
		err = errors.New("Sequence file empty " + seqFile.Name() + ": no lines")
		return
	} else if lineCount != 1 {
		err = errors.New("Malformed " + seqFile.Name() + ": wrong number of lines")
		return
	}
	sequence, err = strconv.ParseUint(sequenceString, 10, 64)
	if err != nil {
		return
	}
	return
}

func getSequenceFromChain(clientCtx client.Context) (sequence uint64, err error) {
	// get sequence number from chain
	account, err := clientCtx.AccountRetriever.GetAccount(clientCtx, clientCtx.FromAddress)
	if err != nil {
		return
	}
	sequence = account.GetSequence()
	return
}

func createSequenceDirectory() (path string, err error) {
	usr, err := user.Current()
	if err != nil {
		return
	}
	homeDir := usr.HomeDir
	path = filepath.Join(GetConfig().RootDir, "sequence")
	// expand tilde to user's home directory
	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(homeDir, path[2:])
	}
	_, err = os.Stat(path)
	// directory already exists
	if !os.IsNotExist(err) {
		return
	}
	err = os.Mkdir(path, os.ModePerm)
	return
}

func openSequenceFile(fromAddress sdk.AccAddress) (file *os.File, err error) {
	path, err := createSequenceDirectory()
	if err != nil {
		return
	}

	addrHex := hex.EncodeToString(fromAddress)
	filename := filepath.Join(path, addrHex)

	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	return
}

// GetTxResponseFromOut converts strings to numbers and unmarshalles out into TxResponse struct
func GetTxResponseFromOut(out *bytes.Buffer) (txResponse sdk.TxResponse, err error) {
	m := regexp.MustCompile(`"([0-9]+?)"`)
	str := m.ReplaceAllString(out.String(), "${1}")

	// We might have YAML here, so we need to convert to JSON first, because TxResponse struct lacks `yaml:"height,omitempty"`, etc.
	// Since JSON is a subset of YAML, passing JSON through YAMLToJSON is a no-op and the result is the byte array of the JSON again.
	j, err := yaml.YAMLToJSON([]byte(str))
	if err != nil {
		return
	}

	err = json.Unmarshal(j, &txResponse)
	return
}
