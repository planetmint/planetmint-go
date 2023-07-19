package asset

import (
	"encoding/hex"
	"planetmint-go/testutil"
	"planetmint-go/testutil/sample"

	assettypes "planetmint-go/x/asset/types"
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

	msg := assettypes.MsgNotarizeAsset{
		Creator:   addr.String(),
		Hash:      cidHash,
		Signature: signature,
		PubKey:    hex.EncodeToString(privKey.PubKey().Bytes()),
	}

	// Prepare Tx
	txBytes, err := testutil.PrepareTx(val, &msg, "machine")
	s.Require().NoError(err)

	// Broadcast Tx
	broadcastTxResponse, err := testutil.BroadcastTx(val, txBytes)
	s.Require().NoError(err)
	s.Require().Equal(uint32(0), broadcastTxResponse.TxResponse.Code)
}
