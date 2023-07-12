package asset

import (
	"encoding/json"
	"fmt"
	"planetmint-go/testutil/network"

	clitestutil "planetmint-go/testutil/cli"
	machinecli "planetmint-go/x/machine/client/cli"

	machinetypes "planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/stretchr/testify/suite"
)

// Queryable pubkey for TestNotarizeAsset
const pubKey = "AjKN6HiWucu1EBwzX0ACnkvomJiLRwq79oPxoLMY1zRw"
const mnemonic = "helmet hedgehog lab actor weekend elbow pelican valid obtain hungry rocket decade tower gallery fit practice cart cherry giggle hair snack glance bulb farm"

// Struct definition of machine E2ETestSuite
type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// Returns new machine E2ETestSuite
func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

// Sets up new machine E2ETestSuite
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

// Tear down machine E2ETestSuite
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *E2ETestSuite) TestNotarizeAsset() {
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
}
