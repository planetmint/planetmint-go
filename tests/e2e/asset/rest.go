package asset

import (
	"encoding/hex"
	"fmt"
	"planetmint-go/testutil"
	"planetmint-go/testutil/sample"

	assettypes "planetmint-go/x/asset/types"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

func (s *E2ETestSuite) TestNotarizeAssetREST() {
	val := s.network.Validators[0]

	// Create Msg
	k, err := val.ClientCtx.Keyring.Key("machine")
	s.Require().NoError(err)

	addr, err := k.GetAddress()
	s.Require().NoError(err)

	privKey, err := val.ClientCtx.Keyring.(unsafeExporter).ExportPrivateKeyObject("machine")
	s.Require().NoError(err)

	sk := hex.EncodeToString(privKey.Bytes())
	cidHash, signature := sample.Asset(sk)

	testCases := []struct {
		name   string
		msg    assettypes.MsgNotarizeAsset
		rawLog string
	}{
		{
			"machine not found",
			assettypes.MsgNotarizeAsset{
				Creator:   addr.String(),
				Hash:      cidHash,
				Signature: signature,
				PubKey:    "human pubkey",
			},
			"machine not found",
		},
		{
			"invalid signature",
			assettypes.MsgNotarizeAsset{
				Creator:   addr.String(),
				Hash:      cidHash,
				Signature: "invalid signature",
				PubKey:    hex.EncodeToString(privKey.PubKey().Bytes()),
			},
			"invalid signature",
		},
		{
			"valid notarization",
			assettypes.MsgNotarizeAsset{
				Creator:   addr.String(),
				Hash:      cidHash,
				Signature: signature,
				PubKey:    hex.EncodeToString(privKey.PubKey().Bytes()),
			},
			"planetmintgo.asset.MsgNotarizeAsset",
		},
	}

	for _, tc := range testCases {
		// Prepare Tx
		txBytes, err := testutil.PrepareTx(val, &tc.msg, "machine")
		s.Require().NoError(err)

		// Broadcast Tx
		broadcastTxResponse, err := testutil.BroadcastTx(val, txBytes)
		s.Require().NoError(err)

		s.network.WaitForNextBlock()

		tx, err := testutil.GetRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", val.APIAddress, broadcastTxResponse.TxResponse.TxHash))
		s.Require().NoError(err)

		var txRes txtypes.GetTxResponse
		err = val.ClientCtx.Codec.UnmarshalJSON(tx, &txRes)
		s.Require().NoError(err)
		s.Require().Contains(txRes.TxResponse.RawLog, tc.rawLog)
	}
}
