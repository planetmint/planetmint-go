package dao

import (
	"math"
	"strconv"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/testutil"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daocli "github.com/planetmint/planetmint-go/x/dao/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

func (s *PopSelectionE2ETestSuite) TestPopSelectionNoActors() {
	out := s.perpareLocalTest()

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionOneActors() {
	err := e2etestutil.AttestMachine(s.network, machines[0].name, machines[0].mnemonic, 0)
	s.Require().NoError(err)

	out := s.perpareLocalTest()

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionTwoActors() {
	err := e2etestutil.AttestMachine(s.network, machines[1].name, machines[1].mnemonic, 1)
	s.Require().NoError(err)

	out := s.perpareLocalTest()

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
}
