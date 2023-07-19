package testutil

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/testutil/network"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// GetRequest defines a wrapper around an HTTP GET request with a provided URL.
// An error is returned if the request or reading the body fails.
func GetRequest(url string) ([]byte, error) {
	res, err := http.Get(url) //nolint:gosec // only used for testing
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// PostRequest defines a wrapper around an HTTP POST request with a provided URL and data.
// An error is returned if the request or reading the body fails.
func PostRequest(url, contentType string, data []byte) ([]byte, error) {
	res, err := http.Post(url, contentType, bytes.NewBuffer(data)) //nolint:gosec // only used	for testing
	if err != nil {
		return nil, fmt.Errorf("error while sending post request: %w", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	bz, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return bz, nil
}

func PrepareTx(val *network.Validator, msg sdk.Msg, signer string) ([]byte, error) {
	k, err := val.ClientCtx.Keyring.Key(signer)
	if err != nil {
		return nil, err
	}

	addr, err := k.GetAddress()
	if err != nil {
		return nil, err
	}

	reqAccountInfo := fmt.Sprintf("%s/cosmos/auth/v1beta1/account_info/%s", val.APIAddress, addr.String())
	respAccountInfo, err := GetRequest(reqAccountInfo)
	if err != nil {
		return nil, err
	}

	var resAccountInfo authtypes.QueryAccountInfoResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(respAccountInfo, &resAccountInfo)
	if err != nil {
		return nil, err
	}

	txBuilder := val.ClientCtx.TxConfig.NewTxBuilder()
	txBuilder.SetMsgs(msg)
	txBuilder.SetGasLimit(200000)
	txBuilder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin("stake", 2)})
	txBuilder.SetTimeoutHeight(0)

	pk, err := k.GetPubKey()
	if err != nil {
		return nil, err
	}

	sk := k.GetLocal().PrivKey

	var priv cryptotypes.PrivKey
	err = val.ClientCtx.Codec.UnpackAny(sk, &priv)
	if err != nil {
		return nil, err
	}

	sigV2 := signing.SignatureV2{
		PubKey: pk,
		Data: &signing.SingleSignatureData{
			SignMode:  val.ClientCtx.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: resAccountInfo.Info.Sequence,
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	signerData := xauthsigning.SignerData{
		ChainID:       val.ClientCtx.ChainID,
		AccountNumber: resAccountInfo.Info.AccountNumber,
		Sequence:      resAccountInfo.Info.Sequence,
	}
	sigV2, err = tx.SignWithPrivKey(
		val.ClientCtx.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, priv, val.ClientCtx.TxConfig, resAccountInfo.Info.Sequence,
	)
	if err != nil {
		return nil, err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	txBytes, err := val.ClientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, nil
}

func BroadcastTx(val *network.Validator, txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	broadcastTxUrl := fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", val.APIAddress)
	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}

	broadCastTxBody, err := val.ClientCtx.Codec.MarshalJSON(&req)
	if err != nil {
		return nil, err
	}
	broadCastTxResponse, err := PostRequest(broadcastTxUrl, "application/json", broadCastTxBody)
	if err != nil {
		return nil, err
	}

	var bctRes txtypes.BroadcastTxResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(broadCastTxResponse, &bctRes)
	if err != nil {
		return nil, err
	}

	return &bctRes, nil
}
