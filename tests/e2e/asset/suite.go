package asset

import (
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"

	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	assettypes "github.com/planetmint/planetmint-go/x/asset/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// E2ETestSuite struct definition of asset suite
type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// NewE2ETestSuite returns configured asset E2ETestSuite
func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

// SetupSuite initializes asset E2ETestSuite
func (s *E2ETestSuite) SetupSuite() {
	conf := config.GetConfig()
	conf.FeeDenom = "stake"

	s.T().Log("setting up e2e test suite")

	s.network = network.New(s.T())
	err := e2etestutil.AttestMachine(s.network, sample.Name, sample.Mnemonic, 0)
	s.Require().NoError(err)

	val := s.network.Validators[0]
	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)

	addr, _ := k.GetAddress()

	clientCtx := val.ClientCtx.
		WithFromAddress(addr).
		WithFromName(sample.Name)
	libConfig := lib.GetConfig()
	libConfig.SetClientCtx(clientCtx)
}

// TearDownSuite clean up after testing
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

// TestNotarizeAsset notarizes asset over cli
func (s *E2ETestSuite) TestNotarizeAsset() {
	val := s.network.Validators[0]
	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)

	addr, _ := k.GetAddress()

	testCases := []struct {
		name             string
		msg              *assettypes.MsgNotarizeAsset
		rawLog           string
		expectCheckTxErr bool
	}{
		{
			"valid notarization",
			assettypes.NewMsgNotarizeAsset(addr.String(), sample.Asset()),
			"[]",
			true,
		},
	}

	for _, tc := range testCases {
		out, err := e2etestutil.BuildSignBroadcastTx(s.T(), addr, tc.msg)
		s.Require().NoError(err)

		txResponse, err := lib.GetTxResponseFromOut(out)
		s.Require().NoError(err)

		s.Require().NoError(s.network.WaitForNextBlock())
		rawLog, err := clitestutil.GetRawLogFromTxOut(val, out)
		s.Require().NoError(err)

		if !tc.expectCheckTxErr {
			assert.Contains(s.T(), rawLog, tc.rawLog)
		} else {
			assert.Contains(s.T(), txResponse.RawLog, tc.rawLog)
		}
	}
}
