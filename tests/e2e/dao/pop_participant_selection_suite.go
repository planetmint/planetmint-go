package dao

import (
	"strconv"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daocli "github.com/planetmint/planetmint-go/x/dao/client/cli"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
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
	cfg := config.GetConfig()
	cfg.FeeDenom = "stake"

	s.network = network.New(s.T(), s.cfg)

	// create 2 machines accounts
	for i, machine := range machines {
		s.attestMachine(machine.name, machine.mnemonic, i)
	}
}

// TearDownSuite clean up after testing
func (s *PopSelectionE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *PopSelectionE2ETestSuite) TestPopSelection() {
	val := s.network.Validators[0]

	// set PopEpochs to 1 in Order to trigger some participant selections
	cfg := config.GetConfig()
	cfg.PopEpochs = 1

	// wait for some blocks so challenges get stored
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	// check if machines are selected as challanger/challengee
	height, _ := s.network.LatestHeight()
	queryHeight := height - 1
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetChallenge(), []string{
		strconv.FormatInt(queryHeight, 10),
	})
	s.Require().NoError(err)

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
}

func (s *PopSelectionE2ETestSuite) attestMachine(name string, mnemonic string, num int) {
	val := s.network.Validators[0]

	account, err := e2etestutil.CreateAccount(s.network, name, mnemonic)
	s.Require().NoError(err)
	err = e2etestutil.FundAccount(s.network, account)
	s.Require().NoError(err)

	// register Ta
	prvKey, pubKey := sample.KeyPair(num)

	ta := sample.TrustAnchor(pubKey)
	registerMsg := machinetypes.NewMsgRegisterTrustAnchor(val.Address.String(), &ta)
	_, err = lib.BroadcastTxWithFileLock(val.Address, registerMsg)
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	addr, err := account.GetAddress()
	s.Require().NoError(err)

	// name and address of private key with which to sign
	clientCtx := val.ClientCtx.
		WithFromAddress(addr).
		WithFromName(name)
	libConfig := lib.GetConfig()
	libConfig.SetClientCtx(clientCtx)

	machine := sample.Machine(name, pubKey, prvKey, addr.String())
	attestMsg := machinetypes.NewMsgAttestMachine(addr.String(), &machine)
	_, err = lib.BroadcastTxWithFileLock(addr, attestMsg)
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	// reset clientCtx to validator ctx
	libConfig.SetClientCtx(val.ClientCtx)
}
