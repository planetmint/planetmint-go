package dao

import (
	"fmt"
	"planetmint-go/testutil/network"
	"planetmint-go/testutil/sample"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clitestutil "planetmint-go/testutil/cli"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
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

	// set the balances in the genesis state
	var bankGenState banktypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[banktypes.ModuleName], &bankGenState)

	bbalances := sdk.NewCoins(
		sdk.NewCoin("rddl", math.NewInt(10000)),
	)

	abalances := sdk.NewCoins(
		sdk.NewCoin("rddl", math.NewInt(5000)),
	)

	accountBalances := []banktypes.Balance{
		{Address: bobAddr.String(), Coins: bbalances.Sort()},
		{Address: aliceAddr.String(), Coins: abalances.Sort()},
	}
	bankGenState.Balances = append(bankGenState.Balances, accountBalances...)
	s.cfg.GenesisState[banktypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&bankGenState)

	s.cfg.MinGasPrices = fmt.Sprintf("0.000006%s", "node0token")
	s.network = network.New(s.T(), s.cfg)
}

// TearDownSuite clean up after testing
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *E2ETestSuite) TestDistributeCollectedFees() {
	val := s.network.Validators[0]

	// sending funds to alice
	args := []string{
		val.Moniker,
		aliceAddr.String(),
		"1000stake",
		"--yes",
		// fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		fmt.Sprintf("--%s=%s", flags.FlagFees, "10node0token"),
	}
	_, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)

	s.network.WaitForNextBlock()
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		aliceAddr.String(),
	})
	s.Require().NoError(err)
	s.T().Log(out)

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		bobAddr.String(),
	})
	s.Require().NoError(err)
	s.T().Log(out)
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)

	s.network.WaitForNextBlock()

	s.network.WaitForNextBlock()
	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		aliceAddr.String(),
	})
	s.Require().NoError(err)
	s.T().Log(out)

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		bobAddr.String(),
	})
	s.Require().NoError(err)
	s.T().Log(out)
}