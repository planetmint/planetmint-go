package asset

import (
	"encoding/hex"
	"fmt"
	"planetmint-go/testutil"
	"planetmint-go/testutil/sample"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"

	assettypes "planetmint-go/x/asset/types"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

// TestNotarizeAssetREST notarizes asset over REST endpoint
func (s *E2ETestSuite) TestNotarizeAssetREST() {
	val := s.network.Validators[0]

	// Create Msg
	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)

	addr, err := k.GetAddress()
	s.Require().NoError(err)

	xskKey, _ := hdkeychain.NewKeyFromString(xPrvKey)
	privKey, _ := xskKey.ECPrivKey()
	byte_key := privKey.Serialize()
	sk := hex.EncodeToString(byte_key)
	cid, signatureHex := sample.Asset(sk)

	testCases := []struct {
		name             string
		msg              assettypes.MsgNotarizeAsset
		rawLog           string
		expectCheckTxErr bool
	}{
		{
			"machine not found",
			assettypes.MsgNotarizeAsset{
				Creator:   addr.String(),
				Hash:      cid,
				Signature: signatureHex,
				PubKey:    "human pubkey",
			},
			"machine not found",
			true,
		},
		{
			"invalid signature hex string",
			assettypes.MsgNotarizeAsset{
				Creator:   addr.String(),
				Hash:      cid,
				Signature: "invalid signature",
				PubKey:    xPubKey,
			},
			"invalid signature hex string",
			false,
		},
		{
			"invalid signature",
			assettypes.MsgNotarizeAsset{
				Creator:   addr.String(),
				Hash:      cid,
				Signature: hex.EncodeToString([]byte("invalid signature")),
				PubKey:    xPubKey,
			},
			"invalid signature",
			false,
		},
		{
			"valid notarization",
			assettypes.MsgNotarizeAsset{
				Creator:   addr.String(),
				Hash:      cid,
				Signature: signatureHex,
				PubKey:    xPubKey,
			},
			"planetmintgo.asset.MsgNotarizeAsset",
			false,
		},
	}

	for _, tc := range testCases {
		// Prepare Tx
		txBytes, err := testutil.PrepareTx(val, &tc.msg, sample.Name)
		s.Require().NoError(err)

		// Broadcast Tx
		broadcastTxResponse, err := testutil.BroadcastTx(val, txBytes)
		s.Require().NoError(err)

		s.Require().NoError(s.network.WaitForNextBlock())

		tx, err := testutil.GetRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", val.APIAddress, broadcastTxResponse.TxResponse.TxHash))
		s.Require().NoError(err)

		if !tc.expectCheckTxErr {
			var txRes txtypes.GetTxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(tx, &txRes)
			s.Require().NoError(err)
			s.Require().Contains(txRes.TxResponse.RawLog, tc.rawLog)
		} else {
			s.Require().Contains(broadcastTxResponse.TxResponse.RawLog, tc.rawLog)
		}
	}
}
