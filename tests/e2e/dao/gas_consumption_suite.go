package dao

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GasConsumptionE2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewGasConsumptionE2ETestSuite(cfg network.Config) *GasConsumptionE2ETestSuite {
	return &GasConsumptionE2ETestSuite{cfg: cfg}
}

func (s *GasConsumptionE2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")
	conf := config.GetConfig()
	conf.FeeDenom = "stake"
	s.network = network.New(s.T(), s.cfg)
}

func (s *GasConsumptionE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suites")
}

func (s *GasConsumptionE2ETestSuite) TestConsumption() {
	val := s.network.Validators[0]

	account, err := e2etestutil.CreateAccount(s.network, sample.Name, sample.Mnemonic)
	s.Require().NoError(err)

	addr, err := account.GetAddress()
	s.Require().NoError(err)

	// send huge tx but as val and with no gas kv costs
	msgs := createMsgs(val.Address, addr, 10)

	out, err := lib.BroadcastTxWithFileLock(val.Address, msgs...)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	_, err = clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().NoError(err)

	// exceed gas limit with too many msgs as non validator
	msgs = createMsgs(addr, val.Address, 10)

	out, err = lib.BroadcastTxWithFileLock(addr, msgs...)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	_, err = clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().Error(err)
	assert.Contains(s.T(), err.Error(), "out of gas")
}

func createMsgs(from sdk.AccAddress, to sdk.AccAddress, n int) (msgs []sdk.Msg) {
	coins := sdk.NewCoins(sdk.NewInt64Coin("stake", 10))
	for i := 0; i < n; i++ {
		msg := banktypes.NewMsgSend(from, to, coins)
		msgs = append(msgs, msg)
	}
	return
}