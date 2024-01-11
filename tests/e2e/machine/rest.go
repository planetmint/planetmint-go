package machine

import (
	"fmt"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/testutil"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
)

// RestE2ETestSuite struct definition of machine suite
type RestE2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// NewRestE2ETestSuite returns configured machine RestE2ETestSuite
func NewRestE2ETestSuite(cfg network.Config) *RestE2ETestSuite {
	return &RestE2ETestSuite{cfg: cfg}
}

// SetupSuite initializes machine E2ETestSuite
func (s *RestE2ETestSuite) SetupSuite() {
	conf := config.GetConfig()
	conf.FeeDenom = "stake"

	s.T().Log("setting up e2e test suite")

	s.network = network.New(s.T())
	// create machine account for attestation
	account, err := e2etestutil.CreateAccount(s.network, sample.Name, sample.Mnemonic)
	s.Require().NoError(err)
	err = e2etestutil.FundAccount(s.network, account)
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
		Creator:     val.Address.String(),
		TrustAnchor: &ta,
	}
	out, err := e2etestutil.BuildSignBroadcastTx(s.T(), val.Address, &taMsg)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.machine.MsgRegisterTrustAnchor")

	// Create Attest Machine TX
	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	msg := machinetypes.MsgAttestMachine{
		Creator: addr.String(),
		Machine: &machine,
	}
	out, err = e2etestutil.BuildSignBroadcastTx(s.T(), addr, &msg)
	s.Require().NoError(err)

	// give machine attestation some time to issue the liquid asset
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	rawLog, err = clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.machine.MsgAttestMachine")

	queryMachineURL := fmt.Sprintf("%s/planetmint/machine/get_machine_by_public_key/%s", baseURL, pubKey)
	queryMachineRes, err := testutil.GetRequest(queryMachineURL)
	s.Require().NoError(err)

	var qmRes machinetypes.QueryGetMachineByPublicKeyResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(queryMachineRes, &qmRes)
	s.Require().NoError(err)
	s.Require().Equal(&machine, qmRes.Machine)
}
