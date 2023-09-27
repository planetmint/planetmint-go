package asset

import (
	"encoding/json"
	"fmt"

	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"

	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	assetcli "github.com/planetmint/planetmint-go/x/asset/client/cli"
	machinecli "github.com/planetmint/planetmint-go/x/machine/client/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	pubKey string
	prvKey string
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

	s.network = network.New(s.T())
	val := s.network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, err := kb.NewAccount(sample.Name, sample.Mnemonic, keyring.DefaultBIP39Passphrase, sample.DefaultDerivationPath, hd.Secp256k1)
	s.Require().NoError(err)

	addr, _ := account.GetAddress()

	// sending funds to machine to initialize account on chain
	args := []string{
		val.Moniker,
		addr.String(),
		sample.Amount,
		"--yes",
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
	}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)

	txResponse, err := clitestutil.GetTxResponseFromOut(out)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxResponse(val, txResponse)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "cosmos.bank.v1beta1.MsgSend")

	s.Require().NoError(s.network.WaitForNextBlock())

	prvKey, pubKey = sample.KeyPair()

	ta := sample.TrustAnchor(pubKey)
	taJSON, err := json.Marshal(&ta)
	s.Require().NoError(err)
	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(taJSON),
	}
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdRegisterTrustAnchor(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	machineJSON, err := json.Marshal(&machine)
	s.Require().NoError(err)

	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, sample.Name),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(machineJSON),
	}

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	s.Require().NoError(err)

	txResponse, err = clitestutil.GetTxResponseFromOut(out)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err = clitestutil.GetRawLogFromTxResponse(val, txResponse)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.machine.MsgAttestMachine")
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
	cid := sample.Asset()

	testCases := []struct {
		name             string
		args             []string
		rawLog           string
		expectCheckTxErr bool
	}{
		{
			"valid notarization",
			[]string{
				cid,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
				"--yes",
			},
			"[]",
			true,
		},
	}

	for _, tc := range testCases {
		out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, assetcli.CmdNotarizeAsset(), tc.args)
		s.Require().NoError(err)

		txResponse, err := clitestutil.GetTxResponseFromOut(out)
		s.Require().NoError(err)

		s.Require().NoError(s.network.WaitForNextBlock())
		rawLog, err := clitestutil.GetRawLogFromTxResponse(val, txResponse)

		if !tc.expectCheckTxErr {
			s.Require().NoError(err)
			assert.Contains(s.T(), rawLog, tc.rawLog)
		} else {
			assert.Contains(s.T(), txResponse.RawLog, tc.rawLog)
		}
	}
}
