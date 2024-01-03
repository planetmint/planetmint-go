package lib

import (
	"bytes"
	"encoding/json"
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sigs.k8s.io/yaml"
)

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
