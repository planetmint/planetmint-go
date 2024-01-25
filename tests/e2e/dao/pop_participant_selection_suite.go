package dao

import (
	"fmt"
	"log"
	"math"
	"strconv"

	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/testutil"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daocli "github.com/planetmint/planetmint-go/x/dao/client/cli"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

var machines = []struct {
	name     string
	mnemonic string
	address  string
}{
	{
		name:     "R2D2",
		mnemonic: "number judge garbage lock village slush business upset suspect green wrestle puzzle foil tragic drum stereo ticket teach upper bone inject monkey deny portion",
		address:  "plmnt1kp93kns6hs2066d8qw0uz84fw3vlthewt2ck6p",
	},
	{
		name:     "C3PO",
		mnemonic: "letter plate husband impulse grid lake panel seminar try powder virtual run spice siege mutual enhance ripple country two boring have convince symptom fuel",
		address:  "plmnt15wrx9eqegjtlvvx80huau7rkn3f44rdj969xrx",
	},
}

type PopSelectionE2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewPopSelectionE2ETestSuite(cfg network.Config) *PopSelectionE2ETestSuite {
	return &PopSelectionE2ETestSuite{cfg: cfg}
}

func (s *PopSelectionE2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")
	conf := config.GetConfig()
	conf.FeeDenom = sample.FeeDenom
	conf.MqttResponseTimeout = 200

	s.network = network.New(s.T(), s.cfg)

	// trigger one participant selection per test
	conf.PopEpochs = 10
	conf.ReissuanceEpochs = 60
	conf.DistributionOffset = 2
}

// TearDownSuite clean up after testing
func (s *PopSelectionE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *PopSelectionE2ETestSuite) perpareLocalTest() testutil.BufferWriter {
	val := s.network.Validators[0]
	conf := config.GetConfig()

	latestHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)

	wait := int(math.Ceil(float64(conf.PopEpochs) / 2.0))
	for {
		latestHeight, err = s.network.WaitForHeight(latestHeight + 1)
		s.Require().NoError(err)

		if latestHeight%int64(conf.PopEpochs) == int64(wait) {
			break
		}
	}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetChallenge(), []string{
		strconv.Itoa(int(latestHeight) - wait),
	})
	s.Require().NoError(err)
	return out
}

type yamlChallenge struct {
	Initiator  string `yaml:"initiator"`
	Challenger string `yaml:"challenger"`
	Challengee string `yaml:"challengee"`
	Height     string `yaml:"height"`
	Success    bool   `yaml:"success"`
	Finished   bool   `yaml:"finished"`
}

func (s *PopSelectionE2ETestSuite) sendPoPResult(storedChallenge []byte, success bool) {
	val := s.network.Validators[0]
	var wrapper struct {
		Challenge yamlChallenge `yaml:"challenge"`
	}

	err := yaml.Unmarshal(storedChallenge, &wrapper)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	tmpChallenge := wrapper.Challenge
	var challenge daotypes.Challenge
	challenge.Challengee = tmpChallenge.Challengee
	challenge.Challenger = tmpChallenge.Challenger
	challenge.Initiator = tmpChallenge.Initiator
	challenge.Height, err = strconv.ParseInt(tmpChallenge.Height, 10, 64)
	s.Require().NoError(err)
	challenge.Finished = true
	challenge.Success = success

	msg := daotypes.NewMsgReportPopResult(val.Address.String(), &challenge)
	_, err = e2etestutil.BuildSignBroadcastTx(s.T(), val.Address, msg)
	s.Require().NoError(err)
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionNoActors() {
	out := s.perpareLocalTest()

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), true)
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionOneActors() {
	err := e2etestutil.AttestMachine(s.network, machines[0].name, machines[0].mnemonic, 0)
	s.Require().NoError(err)

	out := s.perpareLocalTest()

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), true)
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionTwoActors() {
	err := e2etestutil.AttestMachine(s.network, machines[1].name, machines[1].mnemonic, 1)
	s.Require().NoError(err)

	out := s.perpareLocalTest()

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), true)
}

func (s *PopSelectionE2ETestSuite) VerifyTokens(token string) {
	val := s.network.Validators[0]
	conf := config.GetConfig()
	// check balance for crddl
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetCmdQueryTotalSupply(), []string{
		fmt.Sprintf("--%s=%s", bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), conf.ClaimDenom)
	assert.Equal(s.T(), "amount: \"17979452050\"\ndenom: "+token+"\n", out.String()) // Total supply 2 * 7990867578 (total supply) + 1 * 1997716894 (challenger) = 17979452050

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		machines[0].address,
		fmt.Sprintf("--%s=%s", bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), token)
	assert.Equal(s.T(), "amount: \"5993150682\"\ndenom: "+token+"\n", out.String()) // 3 * 1997716894 = 5993150682

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		machines[1].address,
		fmt.Sprintf("--%s=%s", bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), token)
	assert.Equal(s.T(), "amount: \"11986301368\"\ndenom: "+token+"\n", out.String()) // 2 * 5993150684 = 11986301368
}

func (s *PopSelectionE2ETestSuite) TestTokenDistribution1() {
	conf := config.GetConfig()

	out := s.perpareLocalTest()

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), false)

	out = s.perpareLocalTest()

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), true)

	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	s.VerifyTokens(conf.StagedDenom)

	// send Reissuance and DistributionResult implicitly
	latestHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)
	for {
		latestHeight, err := s.network.WaitForHeight(latestHeight + 1)
		s.Require().NoError(err)
		// s.Require().NoError(s.network.WaitForNextBlock())
		if latestHeight%int64(conf.ReissuanceEpochs) == int64(conf.DistributionOffset) {
			break
		}
	}
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	s.VerifyTokens(conf.ClaimDenom)
}
