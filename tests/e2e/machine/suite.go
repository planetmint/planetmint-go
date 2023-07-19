package machine

import (
	"encoding/json"
	"fmt"
	clitestutil "planetmint-go/testutil/cli"
	"planetmint-go/testutil/network"
	machinecli "planetmint-go/x/machine/client/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/stretchr/testify/suite"

	machinetypes "planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Queryable pubkey for TestAttestMachine
const pubKey = "AjKN6HiWucu1EBwzX0ACnkvomJiLRwq79oPxoLMY1zRw"
const mnemonic = "helmet hedgehog lab actor weekend elbow pelican valid obtain hungry rocket decade tower gallery fit practice cart cherry giggle hair snack glance bulb farm"

// E2ETestSuite struct definition of machine suite
type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// NewE2ETestSuite returns configured machine E2ETestSuite
func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

// SetupSuite initializes machine E2ETestSuite
func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	s.network = network.New(s.T())
	val := s.network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, err := kb.NewAccount("machine", mnemonic, keyring.DefaultBIP39Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)
	s.Require().NoError(err)

	addr, _ := account.GetAddress()

	// sending funds to machine to initialize account on chain
	args := []string{
		"node0",
		addr.String(),
		"1000stake",
		"--yes",
		fmt.Sprintf("--%s=%s", flags.FlagFees, "2stake"),
	}
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
}

// TearDownSuite clean up after testing
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

// TestAttestMachine attests machine and query attested machine from chain
func (s *E2ETestSuite) TestAttestMachine() {
	val := s.network.Validators[0]

	machine := machinetypes.Machine{
		Name:             "machine",
		Ticker:           "machine_ticker",
		Issued:           1,
		Amount:           1000,
		Precision:        8,
		IssuerPlanetmint: pubKey,
		IssuerLiquid:     pubKey,
		MachineId:        pubKey,
		Metadata: &machinetypes.Metadata{
			AdditionalDataCID: "CID",
			Gps:               "{\"Latitude\":\"-48.876667\",\"Longitude\":\"-123.393333\"}",
		},
	}
	machineJSON, err := json.Marshal(&machine)
	s.Require().NoError(err)

	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, "machine"),
		fmt.Sprintf("--%s=%s", flags.FlagFees, "2stake"),
		"--yes",
		string(machineJSON),
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	args = []string{
		pubKey,
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdGetMachineByPublicKey(), args)
	s.Require().NoError(err)
}
