package asset

import (
	"fmt"
	"planetmint-go/testutil"
	"planetmint-go/testutil/sample"

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

	cidHash, signature := sample.Asset(prvKey)

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
				PubKey:    pubKey,
			},
			"invalid signature",
		},
		{
			"valid notarization",
			assettypes.MsgNotarizeAsset{
				Creator:   addr.String(),
				Hash:      cidHash,
				Signature: signature,
				PubKey:    pubKey,
			},
			"planetmintgo.asset.MsgNotarizeAsset",
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

		var txRes txtypes.GetTxResponse
		err = val.ClientCtx.Codec.UnmarshalJSON(tx, &txRes)
		s.Require().NoError(err)
		s.Require().Contains(txRes.TxResponse.RawLog, tc.rawLog)
	}
}
