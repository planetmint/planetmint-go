package machine

import (
	"fmt"
	"strings"

	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/planetmint/planetmint-go/lib"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
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

// ProsumeE2ETestSuite struct definition of prosume e2e test suite
type ProsumeE2ETestSuite struct {
	suite.Suite

	cfg      network.Config
	network  *network.Network
	feeDenom string
}

// NewProsumeE2ETestSuite returns configured prosume e2e test suite
func NewProsumeE2ETestSuite(cfg network.Config) *ProsumeE2ETestSuite {
	return &ProsumeE2ETestSuite{cfg: cfg}
}

func (s *ProsumeE2ETestSuite) SetupSuite() {
	s.T().Log("setting up prosume e2e test suite")

	s.feeDenom = sample.FeeDenom
	s.network = network.Load(s.T(), s.cfg)

	// create machines
	for i, machine := range machines {
		s.Require().NoError(e2etestutil.AttestMachine(s.network, machine.name, machine.mnemonic, i, s.feeDenom))
	}
}

func (s *ProsumeE2ETestSuite) TestProsume() {
	val := s.network.Validators[0]

	producerKeyring, err := val.ClientCtx.Keyring.Key(machines[0].name)
	s.Require().NoError(err)
	producerAddress, err := producerKeyring.GetAddress()
	s.Require().NoError(err)

	production := sdk.NewCoins(sdk.NewCoin("kwh", sdk.NewInt(1000)))
	mintProdMsg := machinetypes.NewMsgMintProduction(machines[0].address, production)
	out, err := e2etestutil.BuildSignBroadcastTx(s.T(), producerAddress, mintProdMsg)
	s.Require().NoError(err)

	txRes, err := lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(0), int(txRes.Code))

	s.Require().NoError(s.network.WaitForNextBlock())

	// check balance
	var queryBankBalanceResp banktypes.QueryAllBalancesResponse
	baseURL := strings.ReplaceAll(strings.ReplaceAll(val.APIAddress, "[", ""), "]", "")
	url := fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s", baseURL, machines[0].address)
	resp, err := sdktestutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, &queryBankBalanceResp))
	s.Require().Equal(queryBankBalanceResp.Balances.AmountOf("kwh"), sdk.NewInt(1000))

	consumerKeyring, err := val.ClientCtx.Keyring.Key(machines[1].name)
	s.Require().NoError(err)
	consumerAddress, err := consumerKeyring.GetAddress()
	s.Require().NoError(err)

	sendMsg := banktypes.NewMsgSend(producerAddress, consumerAddress, production)
	out, err = e2etestutil.BuildSignBroadcastTx(s.T(), producerAddress, sendMsg)
	s.Require().NoError(err)

	txRes, err = lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(0), int(txRes.Code))

	s.Require().NoError(s.network.WaitForNextBlock())

	// check balances
	url = fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s", baseURL, machines[0].address)
	resp, err = sdktestutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, &queryBankBalanceResp))
	s.Require().Equal(queryBankBalanceResp.Balances.AmountOf("kwh"), sdk.NewInt(0))

	url = fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s", baseURL, machines[1].address)
	resp, err = sdktestutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, &queryBankBalanceResp))
	s.Require().Equal(queryBankBalanceResp.Balances.AmountOf("kwh"), sdk.NewInt(1000))

	consumption := sdk.NewCoins(sdk.NewCoin("kwh", sdk.NewInt(500)))
	burnConsMsg := machinetypes.NewMsgBurnConsumption(machines[1].address, consumption)
	out, err = e2etestutil.BuildSignBroadcastTx(s.T(), consumerAddress, burnConsMsg)
	s.Require().NoError(err)

	txRes, err = lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(0), int(txRes.Code))

	s.Require().NoError(s.network.WaitForNextBlock())

	// check balance
	url = fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s", baseURL, machines[1].address)
	resp, err = sdktestutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, &queryBankBalanceResp))
	s.Require().Equal(queryBankBalanceResp.Balances.AmountOf("kwh"), sdk.NewInt(500))
}
