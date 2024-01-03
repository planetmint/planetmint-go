package machine

import (
	"fmt"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/testutil"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/stretchr/testify/suite"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

type RestE2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewRestE2ETestSuite(cfg network.Config) *RestE2ETestSuite {
	return &RestE2ETestSuite{cfg: cfg}
}

func (s *RestE2ETestSuite) SetupSuite() {
	cfg := config.GetConfig()
	cfg.FeeDenom = "stake"

	s.T().Log("setting up e2e test suite")

	s.network = network.New(s.T())
	// create machine account for attestation
	err := CreateAccount(s.network, sample.Name, sample.Mnemonic)
	s.Require().NoError(err)
}

// TearDownSuite clean up after testing
func (s *RestE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *RestE2ETestSuite) TestAttestMachineREST() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress

	// Query Sequence Number
	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)

	addr, err := k.GetAddress()
	s.Require().NoError(err)

	prvKey, pubKey := sample.KeyPair()

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

	queryMachineURL := fmt.Sprintf("%s/planetmint/machine/get_machine_by_public_key/%s", baseURL, pubKey)
	queryMachineRes, err := testutil.GetRequest(queryMachineURL)
	s.Require().NoError(err)

	var qmRes machinetypes.QueryGetMachineByPublicKeyResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(queryMachineRes, &qmRes)
	s.Require().NoError(err)
	s.Require().Equal(&machine, qmRes.Machine)
}
