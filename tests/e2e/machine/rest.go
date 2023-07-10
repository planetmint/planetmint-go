package machine

import (
	"fmt"
	"planetmint-go/testutil"
	machinetypes "planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/codec/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
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

	s.T().Log(string(respAccountInfo))

	var resAccountInfo authtypes.QueryAccountInfoResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(respAccountInfo, &resAccountInfo)
	s.Require().NoError(err)

	s.T().Log(resAccountInfo.Info.AccountNumber)
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
	// machineJSON, err := json.Marshal(&machine)
	// s.Require().NoError(err)

	msg := machinetypes.MsgAttestMachine{
		Creator: string(addr),
		Machine: &machine,
	}

	msg.Type()

	// Encode TX
	reqEncodeTx := fmt.Sprintf("%s/cosmos/tx/encode", baseURL)

	signerInfos := make([]txtypes.SignerInfo, 1)
	signerInfos[1].PublicKey = &types.Any{
		TypeUrl: "/cosmos.crypto.secp256k1.PubKey",
		Value:   k.PubKey.Value,
	}
	signerInfos[1].ModeInfo = &txtypes.ModeInfo{
		Sum: &txtypes.ModeInfo_Single_{
			Single: &txtypes.ModeInfo_Single{
				Mode: signing.SignMode_SIGN_MODE_DIRECT,
			},
		},
	}
	signerInfos[1].Sequence = resAccountInfo.Info.Sequence

	reqEncodeTxBody := txtypes.TxEncodeRequest{
		Tx: &txtypes.Tx{
			Body: &txtypes.TxBody{
				// Messages: make([]*types.Any, 0),
			},
			AuthInfo: &txtypes.AuthInfo{
				SignerInfos: signerInfos,
				Tip:         nil,
			},
			Signatures: [][]byte{},
		},
	}
	s.Require().NoError(err)
	// var resEncodeTX TXEncodeRequest
	respEncodeTx, err := testutil.PostRequest(reqEncodeTx, "application/json", []byte(reqEncodeTxBody))
	s.Require().NoError(err)

	s.T().Log(respEncodeTx)
	// Send encoded TX to /cosmos/tx/v1beta1/txs

	// Query Machine by Pubkey
}
