package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
		fmt.Printf("Error starting command: %s\n", errstr)
		err = errors.New("Reissuance of RDDL failed.")
	} else {
		var txobj ReissueResult
		json.Unmarshal(stdout.Bytes(), &txobj)
		txid = txobj.Txid
	}
	return txid, err
}
