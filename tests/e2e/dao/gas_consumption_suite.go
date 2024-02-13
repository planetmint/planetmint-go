package dao

import (
	"bufio"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/planetmint/planetmint-go/lib"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GasConsumptionE2ETestSuite struct {
	suite.Suite

	cfg        network.Config
	network    *network.Network
	minterAddr sdk.AccAddress
}

func NewGasConsumptionE2ETestSuite(cfg network.Config) *GasConsumptionE2ETestSuite {
	return &GasConsumptionE2ETestSuite{cfg: cfg}
}

func (s *GasConsumptionE2ETestSuite) createValAccount(cfg network.Config) (address sdk.AccAddress, err error) {
	buf := bufio.NewReader(os.Stdin)

	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, s.T().TempDir(), buf, cfg.Codec, cfg.KeyringOptions...)
	if err != nil {
		return nil, err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(cfg.SigningAlgo, keyringAlgos)
	if err != nil {
		return nil, err
	}

	mnemonic := cfg.Mnemonics[0]

	record, err := kb.NewAccount("node0", mnemonic, keyring.DefaultBIP39Passphrase, sdk.GetConfig().GetFullBIP44Path(), algo)
	if err != nil {
		return nil, err
	}

	addr, err := record.GetAddress()
	if err != nil {
		return nil, err
	}

	return addr, nil
}

func (s *GasConsumptionE2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	s.cfg.Mnemonics = []string{sample.Mnemonic}
	addr, err := s.createValAccount(s.cfg)
	s.Require().NoError(err)

	// set accounts for alice and bob in genesis state
	var authGenState authtypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[authtypes.ModuleName], &authGenState)
	s.minterAddr = addr
	minter := authtypes.NewBaseAccount(s.minterAddr, nil, 0, 0)
	accounts, err := authtypes.PackAccounts(authtypes.GenesisAccounts{minter})
	s.Require().NoError(err)
	authGenState.Accounts = append(authGenState.Accounts, accounts...)
	s.cfg.GenesisState[authtypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&authGenState)

	var daoGenState daotypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[daotypes.ModuleName], &daoGenState)
	daoGenState.Params.FeeDenom = sample.FeeDenom
	daoGenState.Params.MintAddress = s.minterAddr.String()
	s.cfg.GenesisState[daotypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&daoGenState)

	s.network = network.New(s.T(), s.cfg)
	account, err := e2etestutil.CreateAccount(s.network, sample.Name, sample.Mnemonic)
	s.Require().NoError(err)
	err = e2etestutil.FundAccount(s.network, account, daoGenState.Params.FeeDenom)
	s.Require().NoError(err)
}

func (s *GasConsumptionE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suites")
}

func (s *GasConsumptionE2ETestSuite) TestValidatorConsumption() {
	val := s.network.Validators[0]

	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	// send huge tx but as val and with no gas kv costs
	msgs := createMsgs(val.Address, addr, 10)

	out, err := lib.BroadcastTxWithFileLock(val.Address, msgs...)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	_, err = clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)
}

func (s *GasConsumptionE2ETestSuite) TestNonValidatorConsumptionOverflow() {
	val := s.network.Validators[0]

	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	// exceed gas limit with too many msgs as non validator
	msgs := createMsgs(addr, val.Address, 10)

	out, err := lib.BroadcastTxWithFileLock(addr, msgs...)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	_, err = clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().Error(err)
	assert.Equal(s.T(), "out of gas in location: Has; gasWanted: 200000, gasUsed: 200701: out of gas", err.Error())
}

func createMsgs(from sdk.AccAddress, to sdk.AccAddress, n int) (msgs []sdk.Msg) {
	coins := sdk.NewCoins(sdk.NewInt64Coin("stake", 10))
	for i := 0; i < n; i++ {
		msg := banktypes.NewMsgSend(from, to, coins)
		msgs = append(msgs, msg)
	}
	return
}

func (s *GasConsumptionE2ETestSuite) TestNetworkBasedTxGasLimit() {
	var gasAmountAboveGlobalGasLimit uint64 = 200000000
	libConfig := lib.GetConfig()
	libConfig.SetTxGas(gasAmountAboveGlobalGasLimit)

	var msgs []sdk.Msg

	for i := 0; i < 1000; i++ {
		mintRequest := sample.MintRequest(s.minterAddr.String(), 1, "hash"+strconv.Itoa(i))
		msg := daotypes.NewMsgMintToken(s.minterAddr.String(), &mintRequest)
		msgs = append(msgs, msg)
	}

	_, err := e2etestutil.BuildSignBroadcastTx(s.T(), s.minterAddr, msgs...)
	s.Require().Error(err)
	s.Assert().Contains(err.Error(), "out of gas in location: txSize; gasWanted: "+strconv.FormatUint(gasAmountAboveGlobalGasLimit, 10)+", gasUsed:")
	s.Assert().Contains(err.Error(), " out of gas")

	s.Require().NoError(s.network.WaitForNextBlock())
}
