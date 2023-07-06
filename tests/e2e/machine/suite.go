package machine

import (
	"fmt"
	clitestutil "planetmint-go/testutil/cli"
	"planetmint-go/testutil/network"
	machine "planetmint-go/x/machine/client/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
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
	val := s.network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, _, err := kb.NewMnemonic("machine", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	addr, _ := account.GetAddress()

	args := []string{
		"node0",
		addr.String(),
		"1000stake",
		"-y",
		fmt.Sprintf("--%s=%s", flags.FlagFees, "2stake"),
	}
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.NewSendTxCmd(), args)
	s.Require().NoError(err)

	s.T().Log(out.String())

	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *E2ETestSuite) TestAttestMachine() {
	val := s.network.Validators[0]

	account, err := val.ClientCtx.Keyring.Key("machine")
	s.Require().NoError(err)

	addr, _ := account.GetAddress()
	s.T().Log(fmt.Sprintf("address: %s", addr))

	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagChainID, s.network.Config.ChainID),
		"{\"name\": \"machine\", \"ticker\": \"machine_ticker\", \"issued\": 1, \"amount\": 1000, \"precision\": 8, \"issuerPlanetmint\": \"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB\", \"issuerLiquid\": \"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB\", \"machineId\": \"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB\", \"metadata\": {\"additionalDataCID\": \"CID\", \"gps\": \"{'Latitude':'-48.876667','Longitude':'-123.393333'}\"}}",
		fmt.Sprintf("--%s=%s", flags.FlagFrom, "machine"),
		"-y",
		fmt.Sprintf("--%s=%s", flags.FlagFees, "2stake"),
	}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, machine.CmdAttestMachine(), args)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())

	s.T().Log(out.String())

	args = []string{
		"A/ZrbETECRq5DNGJZ0aH0DjlV4Y1opMlRfGoEJH454eB",
	}

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machine.CmdGetMachineByPublicKey(), args)
	s.Require().NoError(err)

	s.T().Log(out.String())
}
