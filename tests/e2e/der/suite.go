package der

import (
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	dercli "github.com/planetmint/planetmint-go/x/der/client/cli"
	dertypes "github.com/planetmint/planetmint-go/x/der/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// E2ETestSuite struct definition of machine suite
type E2ETestSuite struct {
	suite.Suite

	cfg      network.Config
	network  *network.Network
	feeDenom string
}

// NewE2ETestSuite returns configured machine E2ETestSuite
func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

// SetupSuite initializes machine E2ETestSuite
func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e machine test suite")

	s.feeDenom = sample.FeeDenom
	s.network = network.Load(s.T(), s.cfg)

	// create machine account for attestation
	account, err := e2etestutil.CreateAccount(s.network, sample.Name, sample.Mnemonic)
	s.Require().NoError(err)
	err = e2etestutil.FundAccount(s.network, account, s.feeDenom)
	s.Require().NoError(err)
}

// TearDownSuite clean up after testing
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e machine test suite")
}

// TestRegisterDER attests a DER and queries the attested DER from the chain
func (s *E2ETestSuite) TestRegisterDER() {
	val := s.network.Validators[0]

	der := dertypes.DER{
		ZigbeeID:      "0123456789123456",
		DirigeraID:    "1123456789123456",
		DirigeraMAC:   "",
		PlmntAddress:  val.Address.String(),
		LiquidAddress: "liquidder",
	}

	msg1 := dertypes.NewMsgRegisterDER(val.Address.String(), &der)
	out, err := e2etestutil.BuildSignBroadcastTx(s.T(), val.Address, msg1)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.der.MsgRegisterDER")

	// Check if DER can be resolved
	output, err := clitestutil.ExecTestCLICmd(val.ClientCtx, dercli.CmdDer(), []string{
		der.ZigbeeID,
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), output.String(), "0123456789123456")

	// Check if the NFT got created
	output, err = clitestutil.ExecTestCLICmd(val.ClientCtx, dercli.CmdNft(), []string{
		der.ZigbeeID,
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), output.String(), "0123456789123456")
}
