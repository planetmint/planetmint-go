package util

import (
	"errors"
	"strings"

	"github.com/planetmint/planetmint-go/config"
	elements "github.com/rddl-network/elements-rpc"
)

func ReissueAsset(reissueTx string) (txID string, err error) {
	conf := config.GetConfig()
	url := conf.GetRPCURL()
	cmdArgs := strings.Split(reissueTx, " ")

	result, err := elements.ReissueAsset(url, []string{cmdArgs[1], cmdArgs[2]})
	if err != nil {
		errstr := err.Error()
		err = errors.New("reissuance of RDDL failed: " + errstr)
		return
	}
	txID = result.TxID
	return
}

func DistributeAsset(address string, amount string) (txID string, err error) {
	conf := config.GetConfig()
	url := conf.GetRPCURL()

	txID, err = elements.SendToAddress(url, []string{
		address,
		`"` + amount + `"`,
		`""`,
		`""`,
		"false",
		"true",
		"null",
		`"unset"`,
		"false",
		`"` + conf.ReissuanceAsset + `"`,
	})
	if err != nil {
		errormessage := "distribution of RDDL failed for " + address + ": " + err.Error()
		err = errors.New(errormessage)
		return
	}
	return
}
