package machine

import (
	"encoding/json"
	"fmt"
	clitestutil "planetmint-go/testutil/cli"
	"planetmint-go/testutil/network"
	"planetmint-go/testutil/sample"
	machinecli "planetmint-go/x/machine/client/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// E2ETestSuite struct definition of machine suite
type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// NewE2ETestSuite returns configured machine E2ETestSuite
func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

// SetupSuite initializes machine E2ETestSuite
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
}

// TearDownSuite clean up after testing
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

// TestAttestMachine attests machine and query attested machine from chain
func (s *E2ETestSuite) TestAttestMachine() {
	val := s.network.Validators[0]

	machine := sample.Machine(sample.Name, sample.PubKey)
	machineJSON, err := json.Marshal(&machine)
	s.Require().NoError(err)

	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, sample.Name),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(machineJSON),
	}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	s.Require().NoError(err)

	txResponse, err := clitestutil.GetTxResponseFromOut(out)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxResponse(val, txResponse)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.machine.MsgAttestMachine")

	args = []string{
		sample.PubKey,
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdGetMachineByPublicKey(), args)
	s.Require().NoError(err)
}
