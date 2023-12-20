package dao

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/planetmint/planetmint-go/config"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daocli "github.com/planetmint/planetmint-go/x/dao/client/cli"
	machinecli "github.com/planetmint/planetmint-go/x/machine/client/cli"
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
		fmt.Sprintf("%d", queryHeight),
	})
	s.Require().NoError(err)

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
}

func (s *PopSelectionE2ETestSuite) attestMachine(name string, mnemonic string, num int) {
	val := s.network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, err := kb.NewAccount(name, mnemonic, keyring.DefaultBIP39Passphrase, sample.DefaultDerivationPath, hd.Secp256k1)
	s.Require().NoError(err)

	addr, _ := account.GetAddress()

	// sending funds to machine to initialize account on chain
	args := []string{
		val.Moniker,
		addr.String(),
		sample.Amount,
		"--yes",
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
	}
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	// register Ta
	prvKey, pubKey := sample.KeyPair(num)

	ta := sample.TrustAnchor(pubKey)
	taJSON, err := json.Marshal(&ta)
	s.Require().NoError(err)
	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, name),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(taJSON),
	}
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdRegisterTrustAnchor(), args)
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	machine := sample.Machine(name, pubKey, prvKey, addr.String())
	machineJSON, err := json.Marshal(&machine)
	s.Require().NoError(err)

	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, name),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(machineJSON),
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())
}
