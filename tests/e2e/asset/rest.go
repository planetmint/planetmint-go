package asset

import (
	"encoding/hex"
	"fmt"
	"planetmint-go/testutil"
	"planetmint-go/testutil/sample"

	assettypes "planetmint-go/x/asset/types"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (s *E2ETestSuite) TestNotarizeAssetREST() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress

	// Query Sequence Number
	k, err := val.ClientCtx.Keyring.Key("machine")
	s.Require().NoError(err)

	addr, err := k.GetAddress()
	s.Require().NoError(err)

	reqAccountInfo := fmt.Sprintf("%s/cosmos/auth/v1beta1/account_info/%s", baseURL, addr.String())
	respAccountInfo, err := testutil.GetRequest(reqAccountInfo)
	s.Require().NoError(err)

	var resAccountInfo authtypes.QueryAccountInfoResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(respAccountInfo, &resAccountInfo)
	s.Require().NoError(err)

	privKey, err := val.ClientCtx.Keyring.(unsafeExporter).ExportPrivateKeyObject("machine")
	s.Require().NoError(err)

	sk := hex.EncodeToString(privKey.Bytes())

	cidHash, signature := sample.Asset(sk)

	msg := assettypes.MsgNotarizeAsset{
		Creator:   addr.String(),
		Hash:      cidHash,
		Signature: signature,
		PubKey:    hex.EncodeToString(privKey.PubKey().Bytes()),
	}

	txBuilder := val.ClientCtx.TxConfig.NewTxBuilder()
	err = txBuilder.SetMsgs(&msg)

	txBuilder.SetGasLimit(200000)
	txBuilder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin("stake", 2)})
	txBuilder.SetTimeoutHeight(0)

	pk, err := k.GetPubKey()
	s.Require().NoError(err)
	secretk := k.GetLocal().PrivKey

	var priv cryptotypes.PrivKey
	err = val.ClientCtx.Codec.UnpackAny(secretk, &priv)
	s.Require().NoError(err)

	sigV2 := signing.SignatureV2{
		PubKey: pk,
		Data: &signing.SingleSignatureData{
			SignMode:  val.ClientCtx.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: resAccountInfo.Info.Sequence,
	}

	err = txBuilder.SetSignatures(sigV2)
	s.Require().NoError(err)

	signerData := xauthsigning.SignerData{
		ChainID:       val.ClientCtx.ChainID,
		AccountNumber: resAccountInfo.Info.AccountNumber,
		Sequence:      resAccountInfo.Info.Sequence,
	}
	sigV2, err = tx.SignWithPrivKey(
		val.ClientCtx.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, priv, val.ClientCtx.TxConfig, resAccountInfo.Info.Sequence,
	)
	s.Require().NoError(err)

	err = txBuilder.SetSignatures(sigV2)
	s.Require().NoError(err)

	txBytes, err := val.ClientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
	s.Require().NoError(err)

	broadcastTxUrl := fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", baseURL)
	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}

	broadCastTxBody, err := val.ClientCtx.Codec.MarshalJSON(&req)
	s.Require().NoError(err)

	s.Require().NoError(err)
	broadCastTxResponse, err := testutil.PostRequest(broadcastTxUrl, "application/json", broadCastTxBody)
	s.Require().NoError(err)

	var bctRes txtypes.BroadcastTxResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(broadCastTxResponse, &bctRes)
	s.Require().NoError(err)
	s.Require().Equal(uint32(0), bctRes.TxResponse.Code)
}
