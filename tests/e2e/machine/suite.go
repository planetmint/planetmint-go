package machine

import (
	"encoding/json"
	"fmt"

	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	machinecli "github.com/planetmint/planetmint-go/x/machine/client/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxOut(val, out)
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

	// register Ta
	prvKey, pubKey := sample.KeyPair()

	ta := sample.TrustAnchor(pubKey)
	taJSON, err := json.Marshal(&ta)
	s.Require().NoError(err)
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, sample.Name),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(taJSON),
	}
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdRegisterTrustAnchor(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.machine.MsgRegisterTrustAnchor")

	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	machineJSON, err := json.Marshal(&machine)
	s.Require().NoError(err)

	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, sample.Name),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(machineJSON),
	}

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err = clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.machine.MsgAttestMachine")

	args = []string{
		pubKey,
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdGetMachineByPublicKey(), args)
	s.Require().NoError(err)
}

func (s *E2ETestSuite) TestInvalidAttestMachine() {
	val := s.network.Validators[0]

	// already used in REST test case
	prvKey, pubKey := sample.KeyPair(1)

	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	machineJSON, err := json.Marshal(&machine)
	s.Require().NoError(err)

	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, sample.Name),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(machineJSON),
	}

	out, _ := clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	txResponse, err := clitestutil.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(txResponse.Code), int(4))

	unregisteredPubKey, unregisteredPrivKey := sample.KeyPair(2)
	machine = sample.Machine(sample.Name, unregisteredPubKey, unregisteredPrivKey, addr.String())
	machineJSON, err = json.Marshal(&machine)
	s.Require().NoError(err)

	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, sample.Name),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(machineJSON),
	}

	out, _ = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	txResponse, err = clitestutil.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(txResponse.Code), int(3))
}

func (s *E2ETestSuite) TestMachineAllowanceAttestation() {
	// create address for machine
	val := s.network.Validators[0]
	kb := val.ClientCtx.Keyring

	account, _, err := kb.NewMnemonic("AllowanceMachine", keyring.English, sample.DefaultDerivationPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	addr, err := account.GetAddress()
	s.Require().NoError(err)

	// register TA
	prvKey, pubKey := sample.KeyPair(3)

	ta := sample.TrustAnchor(pubKey)
	taJSON, err := json.Marshal(&ta)
	s.Require().NoError(err)
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Moniker),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(taJSON),
	}
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdRegisterTrustAnchor(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	// create allowance for machine
	args = []string{
		val.Moniker,   // granter
		addr.String(), // grantee
		fmt.Sprintf("--%s=%s", feegrant.FlagAllowedMsgs, "/planetmintgo.machine.MsgAttestMachine"),
		fmt.Sprintf("--%s=%s", feegrant.FlagSpendLimit, "2stake"),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, feegrant.NewCmdFeeGrant(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	// attest machine with fee granter without funding the machine account first
	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	machineJSON, err := json.Marshal(&machine)
	s.Require().NoError(err)

	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		fmt.Sprintf("--%s=%s", flags.FlagFeeGranter, val.Address.String()),
		"--yes",
		string(machineJSON),
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	args = []string{
		pubKey,
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdGetMachineByPublicKey(), args)
	s.Require().NoError(err)
}
