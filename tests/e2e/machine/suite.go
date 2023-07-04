package machine

import (
	"fmt"
	clitestutil "planetmint-go/testutil/cli"
	"planetmint-go/testutil/network"
	machine "planetmint-go/x/machine/client/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	s.network = network.New(s.T())

	kb := s.network.Validators[0].ClientCtx.Keyring
	_, _, err := kb.NewMnemonic("alice", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	_, _, err = kb.NewMnemonic("machine", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *E2ETestSuite) TestAttestMachine() {
	val := s.network.Validators[0]

	account, err := val.ClientCtx.Keyring.Key("alice")
	s.Require().NoError(err)

	addr, _ := account.GetAddress()
	s.T().Log(fmt.Sprintf("address: %s", addr))

	account, err = val.ClientCtx.Keyring.Key("machine")
	s.Require().NoError(err)

	addr, _ = account.GetAddress()
	s.T().Log(fmt.Sprintf("address: %s", addr))

	// cmd := machine.CmdAttestMachine()
	// args := []string{"{\"name\": \"machine\", \"ticker\": \"machine_ticker\", \"issued\": 1, \"amount\": 1000, \"precision\": 8, \"issuerPlanetmint\": \"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB\", \"issuerLiquid\": \"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB\", \"machineId\": \"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB\", \"metadata\": {\"additionalDataCID\": \"CID\", \"gps\": \"{'Latitude':'-48.876667','Longitude':'-123.393333'}\"}}"}

	// testCmd := &cobra.Command{}
	// testCmd.SetArgs(args)

	// err = cmd.RunE(testCmd, args)
	// s.Require().NoError(err)

	args := []string{
		"{\"name\": \"machine\", \"ticker\": \"machine_ticker\", \"issued\": 1, \"amount\": 1000, \"precision\": 8, \"issuerPlanetmint\": \"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB\", \"issuerLiquid\": \"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB\", \"machineId\": \"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB\", \"metadata\": {\"additionalDataCID\": \"CID\", \"gps\": \"{'Latitude':'-48.876667','Longitude':'-123.393333'}\"}}",
		// addr.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machine.CmdAttestMachine(), args)
	s.Require().NoError(err)
}
