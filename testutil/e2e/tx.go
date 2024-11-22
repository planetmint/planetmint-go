package e2e

import (
	"bytes"
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/lib"
)

// BuildSignBroadcastTx builds, signs and broadcasts transaction to the network.
func BuildSignBroadcastTx(t *testing.T, addr sdk.AccAddress, msgs ...sdk.Msg) (out *bytes.Buffer, err error) {
	out, err = lib.BroadcastTxWithFileLock(addr, msgs...)
	if err != nil {
		t.Log("broadcast tx failed: " + err.Error())
		return
	}
	txResponse, err := lib.ParseTxResponse(out)
	if err != nil {
		t.Log("getting tx response from out failed: " + err.Error())
		return
	}
	if txResponse.Code != 0 {
		err = errors.New(txResponse.RawLog)
		return
	}
	return
}
