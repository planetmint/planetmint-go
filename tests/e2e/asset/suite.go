package asset

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"planetmint-go/testutil/network"
	"planetmint-go/testutil/sample"

	clitestutil "planetmint-go/testutil/cli"
	assetcli "planetmint-go/x/asset/client/cli"
	machinecli "planetmint-go/x/machine/client/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
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

	s.network = network.New(s.T())
	val := s.network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, err := kb.NewAccount(sample.Name, sample.Mnemonic, keyring.DefaultBIP39Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)
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

	machine := sample.Machine(sample.Name, sample.PubKey)
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

// Needed to export private key from Keyring
type unsafeExporter interface {
	ExportPrivateKeyObject(uid string) (types.PrivKey, error)
}

// TestNotarizeAsset notarizes asset over cli
func (s *E2ETestSuite) TestNotarizeAsset() {
	val := s.network.Validators[0]

	privKey, err := val.ClientCtx.Keyring.(unsafeExporter).ExportPrivateKeyObject(sample.Name)
	s.Require().NoError(err)

	sk := hex.EncodeToString(privKey.Bytes())

	cidHash, signature := sample.Asset(sk)

	testCases := []struct {
		name   string
		args   []string
		rawLog string
	}{
		{
			"machine not found",
			[]string{
				cidHash,
				signature,
				"pubkey",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sample.Name),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
				"--yes",
			},
			"machine not found",
		},
		{
			"invalid signature",
			[]string{
				"cid",
				"signature",
				hex.EncodeToString(privKey.PubKey().Bytes()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sample.Name),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
				"--yes",
			},
			"invalid signature",
		},
		{
			"valid notarization",
			[]string{
				cidHash,
				signature,
				hex.EncodeToString(privKey.PubKey().Bytes()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sample.Name),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
				"--yes",
			},
			"planetmintgo.asset.MsgNotarizeAsset",
		},
	}

	for _, tc := range testCases {
		out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, assetcli.CmdNotarizeAsset(), tc.args)
		s.Require().NoError(err)

		txResponse, err := clitestutil.GetTxResponseFromOut(out)
		s.Require().NoError(err)

		s.Require().NoError(s.network.WaitForNextBlock())
		rawLog, err := clitestutil.GetRawLogFromTxResponse(val, txResponse)
		s.Require().NoError(err)

		assert.Contains(s.T(), rawLog, tc.rawLog)
	}
}
