package lib

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sigs.k8s.io/yaml"
)

func openSequenceFile(fromAddress sdk.AccAddress) (file *os.File, err error) {
	usr, err := user.Current()
	if err != nil {
		return
	}
	homeDir := usr.HomeDir

	addrHex := hex.EncodeToString(fromAddress)
	filename := filepath.Join(GetConfig().RootDir, addrHex+".sequence")

	// Expand tilde to user's home directory.
	if filename == "~" {
		filename = homeDir
	} else if strings.HasPrefix(filename, "~/") {
		filename = filepath.Join(homeDir, filename[2:])
	}

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
