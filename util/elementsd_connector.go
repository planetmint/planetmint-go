package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
	"strconv"
	"strings"

	"github.com/planetmint/planetmint-go/config"
)

type ReissueResult struct {
	Txid string `json:"txid"`
	Vin  int    `json:"vin"`
}

func ReissueAsset(reissue_tx string) (txid string, err error) {
	conf := config.GetConfig()
	cmd_args := strings.Split(reissue_tx, " ")
	cmd := exec.Command("/usr/local/bin/elements-cli", "-rpcpassword="+conf.RPCPassword,
		"-rpcuser="+conf.RPCUser, "-rpcport="+strconv.Itoa(conf.RPCPort), "-rpcconnect="+conf.RPCHost,
		cmd_args[0], cmd_args[1], cmd_args[2])

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	errstr := stderr.String()

	if err != nil || len(errstr) > 0 {
		err = errors.New("reissuance of RDDL failed")
	} else {
		var txobj ReissueResult
		err = json.Unmarshal(stdout.Bytes(), &txobj)
		if err == nil {
			txid = txobj.Txid
		}
	}
	return txid, err
}

func DistributeAsset(address string, amount string) (txid string, err error) {
	conf := config.GetConfig()
	cmd := exec.Command("/usr/local/bin/elements-cli", "-rpcpassword="+conf.RPCPassword,
		"-rpcuser="+conf.RPCUser, "-rpcport="+strconv.Itoa(conf.RPCPort), "-rpcconnect="+conf.RPCHost,
		"sendtoaddress", address, amount, "", "", "false", "true", "null", "\"unset\"", "false",
		"\""+conf.ReissuanceAsset+"\"")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	errstr := stderr.String()

	if err != nil || len(errstr) > 0 {
		errormessage := "distribution of RDDL failed for " + address
		err = errors.New(errormessage)
	} else {
		var txobj ReissueResult
		err = json.Unmarshal(stdout.Bytes(), &txobj)
		if err == nil {
			txid = txobj.Txid
		}
	}
	return txid, err
}
