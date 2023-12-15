package dao

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/planetmint/planetmint-go/config"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daocli "github.com/planetmint/planetmint-go/x/dao/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

var (
	bobAddr   sdk.AccAddress
	aliceAddr sdk.AccAddress
)

// E2ETestSuite struct definition of dao suite
type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// NewE2ETestSuite returns configured dao E2ETestSuite
func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

// SetupSuite initializes dao E2ETestSuite
func (s *E2ETestSuite) SetupSuite() {
	// set FeeDenom to node0token because the sending account is initialized with no plmnt tokens
	conf := config.GetConfig()
	conf.FeeDenom = "node0token"
	// set epochs: make sure to start after initial height of 7
	conf.ReIssuanceEpochs = 25
	conf.SetPlanetmintConfig(conf)

	s.T().Log("setting up e2e test suite")

	// set accounts for alice and bob in genesis state
	var authGenState authtypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[authtypes.ModuleName], &authGenState)

	bobAddr = sample.Secp256k1AccAddress()
	aliceAddr = sample.Secp256k1AccAddress()

	bob := authtypes.NewBaseAccount(bobAddr, nil, 0, 0)
	alice := authtypes.NewBaseAccount(aliceAddr, nil, 0, 0)

	accounts, err := authtypes.PackAccounts(authtypes.GenesisAccounts{bob, alice})
	s.Require().NoError(err)

	authGenState.Accounts = append(authGenState.Accounts, accounts...)
	s.cfg.GenesisState[authtypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&authGenState)

	// set the balances in genesis state
	var bankGenState banktypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[banktypes.ModuleName], &bankGenState)

	bbalances := sdk.NewCoins(
		sdk.NewCoin(conf.TokenDenom, math.NewInt(10000)),
		sdk.NewCoin(conf.StakeDenom, math.NewInt(10000)),
	)

	abalances := sdk.NewCoins(
		sdk.NewCoin(conf.TokenDenom, math.NewInt(10000)),
		sdk.NewCoin(conf.StakeDenom, math.NewInt(5000)),
	)

	accountBalances := []banktypes.Balance{
		{Address: bobAddr.String(), Coins: bbalances.Sort()},
		{Address: aliceAddr.String(), Coins: abalances.Sort()},
	}
	bankGenState.Balances = append(bankGenState.Balances, accountBalances...)
	s.cfg.GenesisState[banktypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&bankGenState)

	// Setup MintAddress parameter in genesis state
	// use sample.Mnemonic to make mint address deterministic for test
	s.cfg.Mnemonics = []string{sample.Mnemonic}

	// set MintAddress in GenesisState
	var daoGenState daotypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[daotypes.ModuleName], &daoGenState)
	valAddr, err := s.createValAccount(s.cfg)
	s.Require().NoError(err)

	daoGenState.Params.MintAddress = valAddr.String()
	s.cfg.GenesisState[daotypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&daoGenState)

	s.cfg.MinGasPrices = fmt.Sprintf("0.000006%s", conf.FeeDenom)
	s.network = network.New(s.T(), s.cfg)
}

// TearDownSuite clean up after testing
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *E2ETestSuite) TestDistributeCollectedFees() {
	conf := config.GetConfig()
	val := s.network.Validators[0]

	// sending funds to alice and pay some fees to be distributed
	args := []string{
		val.Moniker,
		aliceAddr.String(),
		"1000stake",
		"--yes",
		fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("10%s", conf.FeeDenom)),
	}
	_, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	// assert that alice has 6 of 20 paid fee tokens based on 5000 stake of 15000 total stake
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		aliceAddr.String(),
	})
	assert.Contains(s.T(), out.String(), "node0token")
	assert.Contains(s.T(), out.String(), "6")
	s.Require().NoError(err)

	// assert that bob has 13 of 20 paid fee tokens based on 10000 stake of 15000 total stake
	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		bobAddr.String(),
	})
	assert.Contains(s.T(), out.String(), "node0token")
	assert.Contains(s.T(), out.String(), "13")
	s.Require().NoError(err)
}

func (s *E2ETestSuite) TestMintToken() {
	conf := config.GetConfig()
	val := s.network.Validators[0]

	mintRequest := sample.MintRequest(aliceAddr.String(), 1000, "hash")
	mrJSON, err := json.Marshal(&mintRequest)
	s.Require().NoError(err)

	// send mint token request from mint address
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Moniker),
		fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("10%s", conf.FeeDenom)),
		"--yes",
		string(mrJSON),
	}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdMintToken(), args)
	s.Require().NoError(err)

	txResponse, err := clitestutil.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(0), int(txResponse.Code))

	s.Require().NoError(s.network.WaitForNextBlock())
	rawLog, err := clitestutil.GetRawLogFromTxResponse(val, txResponse)
	s.Require().NoError(err)

	assert.Contains(s.T(), rawLog, "planetmintgo.dao.MsgMintToken")

	// assert that alice has actually received the minted tokens 10000 (initial supply) + 1000 (minted) = 11000 (total)
	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		aliceAddr.String(),
	})
	assert.Contains(s.T(), out.String(), "plmnt")
	assert.Contains(s.T(), out.String(), "11000")
	s.Require().NoError(err)

	// send mint token request from non mint address
	kb := val.ClientCtx.Keyring
	account, err := kb.NewAccount(sample.Name, sample.Mnemonic, keyring.DefaultBIP39Passphrase, sample.DefaultDerivationPath, hd.Secp256k1)
	s.Require().NoError(err)

	addr, _ := account.GetAddress()

	// sending funds to account to initialize on chain
	args = []string{
		val.Moniker,
		addr.String(),
		sample.Amount,
		"--yes",
		fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("10%s", conf.FeeDenom)),
	}
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
		fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("10%s", conf.FeeDenom)),
		"--yes",
		string(mrJSON),
	}

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdMintToken(), args)
	s.Require().NoError(err)

	txResponse, err = clitestutil.GetTxResponseFromOut(out)
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
	conf := config.GetConfig()
	val := s.network.Validators[0]

	var err error
	latestHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)

	var wait int
	for {
		latestHeight, err = s.network.WaitForHeight(latestHeight + 1)
		s.Require().NoError(err)

		// wait + for sending the re-issuance result, i.e.:
		//   block 25: initializing RDDL re-issuance broadcast tx succeeded
		//   block 26: sending the re-issuance result broadcast tx succeeded
		wait = 2
		if latestHeight%int64(conf.ReIssuanceEpochs+wait) == 0 {
			break
		}
	}

	// - because we waited on the re-issuance result, see above
	intValue := strconv.FormatInt(latestHeight-int64(wait), 10)
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetReissuance(), []string{intValue})
	s.Require().NoError(err)
}

func (s *E2ETestSuite) TestPoPResult() {
	conf := config.GetConfig()
	conf.PopEpochs = 1
	val := s.network.Validators[0]

	// send PoP results
	challenges := make([]daotypes.Challenge, 5)
	for i := range challenges {
		blockHeight := (i + 1) * config.GetConfig().PopEpochs
		challenges[i].Height = int64(blockHeight)
		challenges[i].Initiator = val.Address.String()
		challenges[i].Challenger = aliceAddr.String()
		challenges[i].Challengee = bobAddr.String()
		challenges[i].Success = true
		challenges[i].Finished = true

		chJSON, err := json.Marshal(&challenges[i])
		s.Require().NoError(err)

		args := []string{
			fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Moniker),
			fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("10%s", conf.FeeDenom)),
			"--yes",
			string(chJSON),
		}

		_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdReportPopResult(), args)
		s.Require().NoError(err)
		s.Require().NoError(s.network.WaitForNextBlock())
	}

	// check balance for stagedcrddl
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetCmdQueryTotalSupply(), []string{
		fmt.Sprintf("--%s=%s", bank.FlagDenom, conf.StagedDenom),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), conf.StagedDenom)
	assert.Contains(s.T(), out.String(), "7990867578") // Total supply 5 * 7990867578 = 39954337890

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		aliceAddr.String(),
		fmt.Sprintf("--%s=%s", bank.FlagDenom, conf.StagedDenom),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), conf.StagedDenom)
	assert.Contains(s.T(), out.String(), "1997716894") // 5 * 1997716894 = 9988584470

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		bobAddr.String(),
		fmt.Sprintf("--%s=%s", bank.FlagDenom, conf.StagedDenom),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), conf.StagedDenom)
	assert.Contains(s.T(), out.String(), "5993150684") // 5 * 5993150684 = 29965753420

	// send ReissuanceProposal
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdReissueRDDLProposal(), []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Moniker),
		fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("10%s", conf.FeeDenom)),
		"--yes",
		aliceAddr.String(),
		"reissueasset 7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9 2996.07000000",
		strconv.FormatInt(challenges[4].Height, 10),
		strconv.FormatInt(challenges[0].Height, 10),
		strconv.FormatInt(challenges[2].Height, 10),
	})
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	// send ReissuanceResult
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdReissueRDDLResult(), []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Moniker),
		fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("10%s", conf.FeeDenom)),
		"--yes",
		aliceAddr.String(),
		"TxID",
		strconv.FormatInt(challenges[4].Height, 10),
	})
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	// send DistributionRequest
	distributionOrder := daotypes.DistributionOrder{
		Proposer:     aliceAddr.String(),
		FirstPop:     challenges[0].Height,
		LastPop:      challenges[2].Height,
		DaoTxID:      "DaoTxID",
		PopTxID:      "PoPTxID",
		InvestorTxID: "InvestorTxID",
	}
	doJSON, err := json.Marshal(&distributionOrder)
	s.Require().NoError(err)

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdDistributionRequest(), []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Moniker),
		fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("10%s", conf.FeeDenom)),
		"--yes",
		string(doJSON),
	})
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	// send DistributionResult
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdDistributionResult(), []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Moniker),
		fmt.Sprintf("--%s=%s", flags.FlagFees, fmt.Sprintf("10%s", conf.FeeDenom)),
		"--yes",
		strconv.FormatInt(challenges[2].Height, 10),
		"DaoTxID",
		"InvestorTxID",
		"PoPTxID",
	})
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	// check balance for crddl
	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetCmdQueryTotalSupply(), []string{
		fmt.Sprintf("--%s=%s", bank.FlagDenom, conf.ClaimDenom),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), conf.ClaimDenom)
	assert.Contains(s.T(), out.String(), "0") // Total supply 3 * 5993150684 + 3 * 1997716894 = 23972602734

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		aliceAddr.String(),
		fmt.Sprintf("--%s=%s", bank.FlagDenom, conf.ClaimDenom),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), conf.ClaimDenom)
	assert.Contains(s.T(), out.String(), "0") // 3 * 1997716894 = 5993150682

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		bobAddr.String(),
		fmt.Sprintf("--%s=%s", bank.FlagDenom, conf.ClaimDenom),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), conf.ClaimDenom)
	assert.Contains(s.T(), out.String(), "0") // 3 * 5993150684 = 17979452052
}
