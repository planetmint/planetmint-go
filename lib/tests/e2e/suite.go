package machine

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/lib/trustwallet"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"

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
	s.T().Log("setting up e2e lib test suite")

	s.network = network.New(s.T())
}

// TearDownSuite clean up after testing
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e lib test suite")
}

func (s *E2ETestSuite) TestBankSendBroadcastTxWithFileLock() {
	val := s.network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, err := kb.NewAccount(sample.Name, sample.Mnemonic, keyring.DefaultBIP39Passphrase, sample.DefaultDerivationPath, hd.Secp256k1)
	s.Require().NoError(err)

	addr, err := account.GetAddress()
	s.Require().NoError(err)

	// incorrect denom
	coin := sdk.NewCoins(sdk.NewInt64Coin("foobar", 1000))
	msg := banktypes.NewMsgSend(val.Address, addr, coin)

	out, err := lib.BroadcastTxWithFileLock(val.Address, msg)
	s.Require().NoError(err)

	txResponse, err := lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	assert.Equal(s.T(), "received wrong fee denom; got: plmnt required: stake: invalid coins", txResponse.RawLog)

	libConfig := lib.GetConfig()
	libConfig.SetFeeDenom("stake")

	// incorrect coin
	out, err = lib.BroadcastTxWithFileLock(val.Address, msg)
	s.Require().NoError(err)

	txResponse, err = lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	assert.Equal(s.T(), "[]", txResponse.RawLog)

	s.Require().NoError(s.network.WaitForNextBlock())
	_, err = clitestutil.GetRawLogFromTxOut(val, out)
	assert.Equal(s.T(), "failed to execute message; message index: 0: spendable balance  is smaller than 1000foobar: insufficient funds", err.Error())

	// valid transaction
	coin = sdk.NewCoins(sdk.NewInt64Coin("stake", 1000))
	msg = banktypes.NewMsgSend(val.Address, addr, coin)

	out, err = lib.BroadcastTxWithFileLock(val.Address, msg)
	s.Require().NoError(err)

	txResponse, err = lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	assert.Equal(s.T(), "[]", txResponse.RawLog)
}

func (s *E2ETestSuite) TestLoadKeys() {
	s.T().SkipNow()
	_, err := setKeys()
	if err == nil {
		_, err = loadKeys()
		s.Require().NoError(err)
	}
	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestSuite) TestOccSigning() {
	s.T().SkipNow()
	val := s.network.Validators[0]

	keys, err := loadKeys()
	s.Require().NoError(err)

	addr := sdk.MustAccAddressFromBech32(keys.PlanetmintAddress)

	coin := sdk.NewCoins(sdk.NewInt64Coin("stake", 10))
	msg := banktypes.NewMsgSend(addr, val.Address, coin)

	libConfig := lib.GetConfig()
	libConfig.SetSerialPort("/dev/ttyACM0")

	out, err := lib.BroadcastTxWithFileLock(addr, msg)
	s.Require().NoError(err)

	txResponse, err := lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal("[]", txResponse.RawLog)
	s.Require().Equal(uint32(0), txResponse.Code)
}

// set sample mnemonic on trust wallet
func setKeys() (string, error) {
	connector, err := trustwallet.NewTrustWalletConnector("/dev/ttyACM0")
	if err != nil {
		return "", err
	}
	return connector.RecoverFromMnemonic(sample.Mnemonic)
}

func loadKeys() (*trustwallet.PlanetMintKeys, error) {
	connector, err := trustwallet.NewTrustWalletConnector("/dev/ttyACM0")
	if err != nil {
		return nil, err
	}
	return connector.GetPlanetmintKeys()
}
