package machine

import (
	"fmt"
	"planetmint-go/testutil"
	"planetmint-go/testutil/sample"
	machinetypes "planetmint-go/x/machine/types"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

func (s *E2ETestSuite) TestAttestMachineREST() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress

	// Query Sequence Number
	k, err := val.ClientCtx.Keyring.Key("machine")
	s.Require().NoError(err)

	addr, err := k.GetAddress()
	s.Require().NoError(err)

	// Create Attest Machine TX
	machine := sample.Machine("machine", pubKey)
	msg := machinetypes.MsgAttestMachine{
		Creator: addr.String(),
		Machine: &machine,
	}

	txBytes, err := testutil.PrepareTx(val, &msg, "machine")
	s.Require().NoError(err)

	broadcastTxResponse, err := testutil.BroadcastTx(val, txBytes)
	s.Require().NoError(err)

	s.network.WaitForNextBlock()
	tx, err := testutil.GetRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", val.APIAddress, broadcastTxResponse.TxResponse.TxHash))
	s.Require().NoError(err)

	var txRes txtypes.GetTxResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(tx, &txRes)
	s.Require().NoError(err)
	s.Require().Equal(uint32(0), txRes.TxResponse.Code)

	queryMachineUrl := fmt.Sprintf("%s/planetmint-go/machine/get_machine_by_public_key/%s", baseURL, pubKey)
	queryMachineRes, err := testutil.GetRequest(queryMachineUrl)
	s.Require().NoError(err)

	var qmRes machinetypes.QueryGetMachineByPublicKeyResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(queryMachineRes, &qmRes)
	s.Require().NoError(err)
	s.Require().Equal(&machine, qmRes.Machine)
}
