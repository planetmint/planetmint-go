package machine

import (
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	machinecli "github.com/planetmint/planetmint-go/x/machine/client/cli"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
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
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

// TestAttestMachine attests machine and query attested machine from chain
func (s *E2ETestSuite) TestAttestMachine() {
	val := s.network.Validators[0]

	// register Ta
	prvKey, pubKey := sample.KeyPair()

	ta := sample.TrustAnchor(pubKey)
	msg1 := machinetypes.NewMsgRegisterTrustAnchor(val.Address.String(), &ta)
	out, err := e2etestutil.BuildSignBroadcastTx(s.T(), val.Address, msg1)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.machine.MsgRegisterTrustAnchor")

	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	msg2 := machinetypes.NewMsgAttestMachine(addr.String(), &machine)
	out, err = e2etestutil.BuildSignBroadcastTx(s.T(), addr, msg2)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err = clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.machine.MsgAttestMachine")

	args := []string{
		pubKey,
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdGetMachineByPublicKey(), args)
	s.Require().NoError(err)
}

func (s *E2ETestSuite) TestInvalidAttestMachine() {
	val := s.network.Validators[0]

	// already used in previous test case
	prvKey, pubKey := sample.KeyPair()

	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	s.Require().NoError(err)

	msg := machinetypes.NewMsgAttestMachine(addr.String(), &machine)
	out, _ := lib.BroadcastTxWithFileLock(addr, msg)
	txResponse, err := lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(txResponse.Code), int(4))

	unregisteredPubKey, unregisteredPrivKey := sample.KeyPair(2)
	machine = sample.Machine(sample.Name, unregisteredPubKey, unregisteredPrivKey, addr.String())
	s.Require().NoError(err)

	msg = machinetypes.NewMsgAttestMachine(addr.String(), &machine)
	out, _ = lib.BroadcastTxWithFileLock(addr, msg)
	txResponse, err = lib.GetTxResponseFromOut(out)
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
	msg1 := machinetypes.NewMsgRegisterTrustAnchor(val.Address.String(), &ta)
	_, err = e2etestutil.BuildSignBroadcastTx(s.T(), val.Address, msg1)
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	// create allowance for machine
	allowedMsgs := []string{"/planetmintgo.machine.MsgAttestMachine"}
	limit := sdk.NewCoins(sdk.NewInt64Coin("stake", 2))
	basic := feegrant.BasicAllowance{
		SpendLimit: limit,
	}
	var grant feegrant.FeeAllowanceI
	grant = &basic
	grant, err = feegrant.NewAllowedMsgAllowance(grant, allowedMsgs)
	s.Require().NoError(err)

	msg2, err := feegrant.NewMsgGrantAllowance(grant, val.Address, addr)
	s.Require().NoError(err)
	_, err = e2etestutil.BuildSignBroadcastTx(s.T(), val.Address, msg2)
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	// attest machine with fee granter without funding the machine account first
	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	s.Require().NoError(err)

	// name and address of private key with which to sign
	clientCtx := val.ClientCtx.
		WithFeeGranterAddress(val.Address)
	libConfig := lib.GetConfig()
	libConfig.SetClientCtx(clientCtx)

	msg3 := machinetypes.NewMsgAttestMachine(addr.String(), &machine)
	_, err = e2etestutil.BuildSignBroadcastTx(s.T(), addr, msg3)
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	// reset clientCtx to validator ctx
	libConfig.SetClientCtx(val.ClientCtx)

	args := []string{
		pubKey,
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdGetMachineByPublicKey(), args)
	s.Require().NoError(err)
}
