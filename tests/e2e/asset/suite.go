package asset

import (
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"

	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	assettypes "github.com/planetmint/planetmint-go/x/asset/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
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
	conf := config.GetConfig()
	conf.FeeDenom = "stake"

	s.T().Log("setting up e2e test suite")

	s.network = network.New(s.T())
	val := s.network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, err := kb.NewAccount(sample.Name, sample.Mnemonic, keyring.DefaultBIP39Passphrase, sample.DefaultDerivationPath, hd.Secp256k1)
	s.Require().NoError(err)

	addr, _ := account.GetAddress()

	// sending funds to machine to initialize account on chain
	coin := sdk.NewCoins(sdk.NewInt64Coin("stake", 1000))
	msg1 := banktypes.NewMsgSend(val.Address, addr, coin)
	out, err := lib.BroadcastTxWithFileLock(val.Address, msg1)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "cosmos.bank.v1beta1.MsgSend")

	s.Require().NoError(s.network.WaitForNextBlock())

	prvKey, pubKey = sample.KeyPair()

	ta := sample.TrustAnchor(pubKey)
	msg2 := machinetypes.NewMsgRegisterTrustAnchor(val.Address.String(), &ta)
	out, err = lib.BroadcastTxWithFileLock(val.Address, msg2)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	_, err = clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	// name and address of private key with which to sign
	clientCtx := val.ClientCtx.
		WithFromAddress(addr).
		WithFromName(sample.Name)
	libConfig := lib.GetConfig()
	libConfig.SetClientCtx(clientCtx)

	machine := sample.Machine(sample.Name, pubKey, prvKey, addr.String())
	msg3 := machinetypes.NewMsgAttestMachine(addr.String(), &machine)
	out, err = lib.BroadcastTxWithFileLock(addr, msg3)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err = clitestutil.GetRawLogFromTxOut(val, out)
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
		out, err := lib.BroadcastTxWithFileLock(addr, tc.msg)
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
