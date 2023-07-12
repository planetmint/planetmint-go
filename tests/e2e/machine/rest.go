package machine

import (
	"fmt"
	"net/url"
	"planetmint-go/testutil"
	machinetypes "planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client/tx"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (s *E2ETestSuite) TestAttestMachineREST() {
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

	// Create Attest Machine TX
	machine := machinetypes.Machine{
		Name:             "machine",
		Ticker:           "machine_ticker",
		Issued:           1,
		Amount:           1000,
		Precision:        8,
		IssuerPlanetmint: pubKey,
		IssuerLiquid:     pubKey,
		MachineId:        pubKey,
		Metadata: &machinetypes.Metadata{
			AdditionalDataCID: "CID",
			Gps:               "{\"Latitude\":\"-48.876667\",\"Longitude\":\"-123.393333\"}",
		},
	}

	txBuilder := val.ClientCtx.TxConfig.NewTxBuilder()

	msg := machinetypes.MsgAttestMachine{
		Creator: addr.String(),
		Machine: &machine,
	}
	err = txBuilder.SetMsgs(&msg)
	s.Require().NoError(err)

	txBuilder.SetGasLimit(200000)
	txBuilder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin("stake", 2)})
	txBuilder.SetTimeoutHeight(0)

	pk, err := k.GetPubKey()
	s.Require().NoError(err)
	sk := k.GetLocal().PrivKey

	var priv cryptotypes.PrivKey
	err = val.ClientCtx.Codec.UnpackAny(sk, &priv)
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

	marshalledReq, err := val.ClientCtx.Codec.MarshalJSON(&req)
	s.Require().NoError(err)

	s.Require().NoError(err)
	r, err := testutil.PostRequest(broadcastTxUrl, "application/json", marshalledReq)
	s.Require().NoError(err)

	s.T().Log("RESULT:")
	s.T().Log(string(r))

	s.T().Log(string(pubKey))
	urlPubKey := url.QueryEscape(pubKey)
	s.T().Log(string(urlPubKey))

	queryMachineUrl := fmt.Sprintf("%s/planetmint-go/machine/get_machine_by_public_key/%s", baseURL, urlPubKey)
	s.T().Log(string(queryMachineUrl))
	queryMachineRes, err := testutil.GetRequest(queryMachineUrl)
	s.Require().NoError(err)
	s.T().Log(string(queryMachineRes))

	// Encode TX
	reqEncodeTx := fmt.Sprintf("%s/cosmos/tx/v1beta1/encode", baseURL)

	msgs := make([]*types.Any, 1)
	msgs[0], err = cdctypes.NewAnyWithValue(&msg)
	s.Require().NoError(err)

	signerInfos := make([]*txtypes.SignerInfo, 1)
	signerInfos[0] = &txtypes.SignerInfo{}
	signerInfos[0].PublicKey = &types.Any{
		TypeUrl: "/cosmos.crypto.secp256k1.PubKey",
		Value:   k.PubKey.Value,
	}
	signerInfos[0].ModeInfo = &txtypes.ModeInfo{
		Sum: &txtypes.ModeInfo_Single_{
			Single: &txtypes.ModeInfo_Single{
				Mode: signing.SignMode_SIGN_MODE_DIRECT,
			},
		},
	}
	signerInfos[0].Sequence = resAccountInfo.Info.Sequence

	reqEncodeTxBody := txtypes.TxEncodeRequest{
		Tx: &txtypes.Tx{
			Body: &txtypes.TxBody{
				Messages:                    msgs,
				Memo:                        "",
				TimeoutHeight:               0,
				ExtensionOptions:            []*types.Any{},
				NonCriticalExtensionOptions: []*types.Any{},
			},
			AuthInfo: &txtypes.AuthInfo{
				SignerInfos: signerInfos,
				Tip:         nil,
			},
			Signatures: [][]byte{},
		},
	}
	reqEncodeTxBodyBytes, err := reqEncodeTxBody.Marshal()
	s.Require().NoError(err)
	// var resEncodeTX TXEncodeRequest
	respEncodeTx, err := testutil.PostRequest(reqEncodeTx, "application/json", reqEncodeTxBodyBytes)
	s.Require().NoError(err)

	s.T().Log(string(respEncodeTx))
	// Send encoded TX to /cosmos/tx/v1beta1/txs

	// Query Machine by Pubkey
}
