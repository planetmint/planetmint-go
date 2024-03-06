package dao

import (
	"bufio"
	"bytes"
	"os"
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/planetmint/planetmint-go/lib"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/util"
	daocli "github.com/planetmint/planetmint-go/x/dao/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/planetmint/planetmint-go/testutil/moduleobject"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

// E2ETestSuite struct definition of dao suite
type E2ETestSuite struct {
	suite.Suite

	cfg                network.Config
	network            *network.Network
	reissuanceEpochs   int64
	distributionOffset int64
	bobAddr            sdk.AccAddress
	aliceAddr          sdk.AccAddress
}

// NewE2ETestSuite returns configured dao E2ETestSuite
func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

// SetupSuite initializes dao E2ETestSuite
func (s *E2ETestSuite) SetupSuite() {
	// set epochs: make sure to start after initial height of 7
	s.reissuanceEpochs = 25
	s.distributionOffset = 5

	s.T().Log("setting up e2e dao test suite")

	// Setup MintAddress parameter in genesis state
	// use sample.Mnemonic to make mint address deterministic for test
	s.cfg.Mnemonics = []string{sample.Mnemonic}
	valAddr, err := s.createValAccount(s.cfg)
	s.Require().NoError(err)

	// set accounts for alice and bob in genesis state
	var authGenState authtypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[authtypes.ModuleName], &authGenState)

	s.bobAddr = sample.Secp256k1AccAddress()
	s.aliceAddr = sample.Secp256k1AccAddress()

	bob := authtypes.NewBaseAccount(s.bobAddr, nil, 0, 0)
	alice := authtypes.NewBaseAccount(s.aliceAddr, nil, 0, 0)

	accounts, err := authtypes.PackAccounts(authtypes.GenesisAccounts{bob, alice})
	s.Require().NoError(err)

	authGenState.Accounts = append(authGenState.Accounts, accounts...)
	s.cfg.GenesisState[authtypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&authGenState)

	var daoGenState daotypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[daotypes.ModuleName], &daoGenState)
	daoGenState.Params.DistributionOffset = s.distributionOffset
	daoGenState.Params.ReissuanceEpochs = s.reissuanceEpochs
	daoGenState.Params.MintAddress = valAddr.String()
	daoGenState.Params.ClaimAddress = valAddr.String()
	s.cfg.GenesisState[daotypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&daoGenState)

	bbalances := sdk.NewCoins(
		sdk.NewCoin(daoGenState.Params.TokenDenom, math.NewInt(10000)),
	)

	abalances := sdk.NewCoins(
		sdk.NewCoin(daoGenState.Params.TokenDenom, math.NewInt(10000)),
	)

	accountBalances := []banktypes.Balance{
		{Address: s.bobAddr.String(), Coins: bbalances.Sort()},
		{Address: s.aliceAddr.String(), Coins: abalances.Sort()},
	}
	// set the balances in genesis state
	var bankGenState banktypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[banktypes.ModuleName], &bankGenState)
	bankGenState.Balances = append(bankGenState.Balances, accountBalances...)
	s.cfg.GenesisState[banktypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&bankGenState)

	s.network = network.Load(s.T(), s.cfg)

	// create account for redeem claim test case
	account, err := e2etestutil.CreateAccount(s.network, sample.Name, sample.Mnemonic)
	s.Require().NoError(err)
	err = e2etestutil.FundAccount(s.network, account, sample.FeeDenom)
	s.Require().NoError(err)
}

// TearDownSuite clean up after testing
func (s *E2ETestSuite) TearDownSuite() {
	util.TerminationWaitGroup.Wait()
	s.T().Log("tearing down e2e dao test suite")
}

func (s *E2ETestSuite) TestMintToken() {
	val := s.network.Validators[0]

	mintRequest := moduleobject.MintRequest(s.aliceAddr.String(), 1000, "hash")
	msg1 := daotypes.NewMsgMintToken(val.Address.String(), &mintRequest)
	out, err := e2etestutil.BuildSignBroadcastTx(s.T(), val.Address, msg1)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.dao.MsgMintToken")

	// assert that alice has actually received the minted tokens 10000 (initial supply) + 1000 (minted) = 11000 (total)
	output, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		s.aliceAddr.String(),
	})
	out, ok := output.(*bytes.Buffer)
	if !ok {
		err = lib.ErrTypeAssertionFailed
		s.Require().NoError(err)
	}
	assert.Contains(s.T(), out.String(), "plmnt")
	assert.Contains(s.T(), out.String(), "11000")
	s.Require().NoError(err)

	// send mint token request from non mint address
	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	msg1 = daotypes.NewMsgMintToken(addr.String(), &mintRequest)
	out, err = lib.BroadcastTxWithFileLock(addr, msg1)
	s.Require().NoError(err)

	txResponse, err := lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(2), int(txResponse.Code))
}

func (s *E2ETestSuite) createValAccount(cfg network.Config) (address sdk.AccAddress, err error) {
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

func (s *E2ETestSuite) TestReissuance() {
	val := s.network.Validators[0]

	var err error
	latestHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)

	var wait int64
	for {
		latestHeight, err = s.network.WaitForHeight(latestHeight + 1)
		s.Require().NoError(err)

		// wait + for sending the reissuance result, i.e.:
		// 0:  block 25: initializing RDDL reissuance broadcast tx succeeded
		// 1:  block 26: sending the reissuance result broadcast tx succeeded
		// 2:  block 27: confirmation
		wait = 2
		if latestHeight%s.reissuanceEpochs == wait {
			break
		}
	}

	// - because we waited on the reissuance result, see above
	intValue := strconv.FormatInt(latestHeight-wait, 10)
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetReissuance(), []string{intValue})
	s.Require().NoError(err)
}
