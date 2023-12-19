package dao

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/planetmint/planetmint-go/config"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	machinecli "github.com/planetmint/planetmint-go/x/machine/client/cli"
	"github.com/stretchr/testify/suite"
)

var (
	R2D2Addr sdk.AccAddress
	C3POAddr sdk.AccAddress
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
	// create 2 machines accounts
	s.T().Log("setting up e2e test suite")
	s.network = network.New(s.T(), s.cfg)

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

	// wait for PopInit

	// out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdGetMachineByPublicKey(), args)
	// s.Require().NoError(err)

	for _, machine := range machines {
		args := []string{
			machine.address,
		}

		out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdGetMachineByAddress(), args)
		s.Require().NoError(err)

		fmt.Println(out)
	}

	// check if machines are selected as challanger/challengee
	cfg := config.GetConfig()
	cfg.PopEpochs = 1

	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *PopSelectionE2ETestSuite) attestMachine(name string, mnemonic string, num int) {
	val := s.network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, err := kb.NewAccount(name, mnemonic, keyring.DefaultBIP39Passphrase, sample.DefaultDerivationPath, hd.Secp256k1)
	s.Require().NoError(err)

	R2D2Addr, _ := account.GetAddress()

	// sending funds to machine to initialize account on chain
	args := []string{
		val.Moniker,
		R2D2Addr.String(),
		sample.Amount,
		"--yes",
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
	}
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)
	fmt.Println(out)
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
	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdRegisterTrustAnchor(), args)
	s.Require().NoError(err)
	fmt.Println(out)
	s.Require().NoError(s.network.WaitForNextBlock())

	machine := sample.Machine(name, pubKey, prvKey, R2D2Addr.String())
	machineJSON, err := json.Marshal(&machine)
	s.Require().NoError(err)

	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, name),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sample.Fees),
		"--yes",
		string(machineJSON),
	}

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	s.Require().NoError(err)
	fmt.Println(out)
	s.Require().NoError(s.network.WaitForNextBlock())
}
