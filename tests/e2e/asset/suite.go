package asset

import (
	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"

	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	assetcli "github.com/planetmint/planetmint-go/x/asset/client/cli"
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
	s.T().Log("setting up e2e test suite")

	s.network = network.Load(s.T(), s.cfg)
	err := e2etestutil.AttestMachine(s.network, sample.Name, sample.Mnemonic, 0, sample.FeeDenom)
	s.Require().NoError(err)
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
			false,
		},
		{
			"machine not found",
			assettypes.NewMsgNotarizeAsset("plmnt1v5394e8vmfrp4qzdav7xkze0f567w3tsgxf09j", sample.Asset()),
			"error during CheckTx or ReCheckTx: machine not found",
			true,
		},
	}

	for _, tc := range testCases {
		out, err := e2etestutil.BuildSignBroadcastTx(s.T(), addr, tc.msg)
		if tc.expectCheckTxErr {
			s.Require().Error(err)
		} else {
			s.Require().NoError(err)
		}

		txResponse, err := lib.GetTxResponseFromOut(out)
		s.Require().NoError(err)

		s.Require().NoError(s.network.WaitForNextBlock())

		if !tc.expectCheckTxErr {
			assert.Equal(s.T(), int(0), int(txResponse.Code))
			args := []string{sample.Asset()}
			asset, err := clitestutil.ExecTestCLICmd(val.ClientCtx, assetcli.CmdGetByCID(), args)
			s.Require().NoError(err)
			assert.Contains(s.T(), asset.String(), sample.Asset())
		} else {
			assert.Contains(s.T(), txResponse.RawLog, tc.rawLog)
		}
	}
}
