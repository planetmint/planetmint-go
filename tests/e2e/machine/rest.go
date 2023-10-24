package machine

import (
	"fmt"

	"github.com/planetmint/planetmint-go/testutil"
	"github.com/planetmint/planetmint-go/testutil/sample"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

func (s *E2ETestSuite) TestAttestMachineREST() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress

	// Query Sequence Number
	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)

	addr, err := k.GetAddress()
	s.Require().NoError(err)

	prvKey, pubKey := sample.KeyPair(1)

	// Register TA
	ta := sample.TrustAnchor(pubKey)
	taMsg := machinetypes.MsgRegisterTrustAnchor{
		Creator:     addr.String(),
		TrustAnchor: &ta,
	}
	txBytes, err := testutil.PrepareTx(val, &taMsg, sample.Name)
	s.Require().NoError(err)

	_, err = testutil.BroadcastTx(val, txBytes)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	// Create Attest Machine TX
	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	msg := machinetypes.MsgAttestMachine{
		Creator: addr.String(),
		Machine: &machine,
	}

	txBytes, err = testutil.PrepareTx(val, &msg, sample.Name)
	s.Require().NoError(err)

	broadcastTxResponse, err := testutil.BroadcastTx(val, txBytes)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	tx, err := testutil.GetRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", val.APIAddress, broadcastTxResponse.TxResponse.TxHash))
	s.Require().NoError(err)

	var txRes txtypes.GetTxResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(tx, &txRes)
	s.Require().NoError(err)
	s.Require().Equal(uint32(0), txRes.TxResponse.Code)

	queryMachineUrl := fmt.Sprintf("%s/planetmint/machine/get_machine_by_public_key/%s", baseURL, pubKey)
	queryMachineRes, err := testutil.GetRequest(queryMachineUrl)
	s.Require().NoError(err)

	var qmRes machinetypes.QueryGetMachineByPublicKeyResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(queryMachineRes, &qmRes)
	s.Require().NoError(err)
	s.Require().Equal(&machine, qmRes.Machine)
}
