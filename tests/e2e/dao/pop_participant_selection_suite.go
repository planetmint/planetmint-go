package dao

import (
	"strconv"

	"github.com/planetmint/planetmint-go/config"
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
	// set PopEpochs to 1 in Order to trigger some participant selections
	conf.PopEpochs = 1

	s.network = network.New(s.T(), s.cfg)
}

// TearDownSuite clean up after testing
func (s *PopSelectionE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionNoActors() {
	val := s.network.Validators[0]
	// wait for some blocks so challenges get stored
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	// check if a PoP without challenger and challengee passes
	height, _ := s.network.LatestHeight()
	queryHeight := height - 1
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetChallenge(), []string{
		strconv.FormatInt(queryHeight, 10),
	})
	s.Require().NoError(err)

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionOneActors() {
	val := s.network.Validators[0]
	////////////////////////////////////////////////////
	// create 1 machinesaccounts
	// ensure that a single machine isn't added to a PoP with only one participant
	err := e2etestutil.AttestMachine(s.network, machines[0].name, machines[0].mnemonic, 0)
	s.Require().NoError(err)

	// wait for some blocks so challenges get stored
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	// check if the next PoP went through with only the challenger selected
	height, _ := s.network.LatestHeight()
	queryHeight := height - 1
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetChallenge(), []string{
		strconv.FormatInt(queryHeight, 10),
	})
	s.Require().NoError(err)

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)

}

func (s *PopSelectionE2ETestSuite) TestPopSelectionTwoActors() {
	val := s.network.Validators[0]
	////////////////////////////////////////////////////
	// create 2nd machine
	err := e2etestutil.AttestMachine(s.network, machines[1].name, machines[1].mnemonic, 1)
	s.Require().NoError(err)

	// wait for some blocks so challenges get stored
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	// check if the next PoP went through with only the challenger selected
	height, _ := s.network.LatestHeight()
	queryHeight := height - 1
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetChallenge(), []string{
		strconv.FormatInt(queryHeight, 10),
	})
	s.Require().NoError(err)

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
}
