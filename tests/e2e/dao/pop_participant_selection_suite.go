package dao

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/testutil"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daocli "github.com/planetmint/planetmint-go/x/dao/client/cli"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
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

	cfg                network.Config
	network            *network.Network
	popEpochs          int64
	reissuanceEpochs   int64
	distributionOffset int64
	claimDenom         string
	feeDenom           string
}

func NewPopSelectionE2ETestSuite(cfg network.Config) *PopSelectionE2ETestSuite {
	return &PopSelectionE2ETestSuite{cfg: cfg}
}

func (s *PopSelectionE2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	s.popEpochs = 10
	s.reissuanceEpochs = 60
	s.distributionOffset = 2
	s.claimDenom = "crddl"
	s.feeDenom = sample.FeeDenom
	s.cfg.Mnemonics = []string{sample.Mnemonic}
	valAddr, err := s.createValAccount(s.cfg)
	s.Require().NoError(err)

	var daoGenState daotypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[daotypes.ModuleName], &daoGenState)
	daoGenState.Params.PopEpochs = s.popEpochs
	daoGenState.Params.ReissuanceEpochs = s.reissuanceEpochs
	daoGenState.Params.DistributionOffset = s.distributionOffset
	daoGenState.Params.MqttResponseTimeout = 200
	daoGenState.Params.FeeDenom = s.feeDenom
	daoGenState.Params.ClaimAddress = valAddr.String()
	s.cfg.GenesisState[daotypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&daoGenState)

	s.network = network.New(s.T(), s.cfg)
}

// TearDownSuite clean up after testing
func (s *PopSelectionE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *PopSelectionE2ETestSuite) perpareLocalTest() testutil.BufferWriter {
	val := s.network.Validators[0]

	latestHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)

	wait := int64(math.Ceil(float64(s.popEpochs) / 2.0))
	for {
		latestHeight, err = s.network.WaitForHeight(latestHeight + 1)
		s.Require().NoError(err)

		if latestHeight%s.popEpochs == wait {
			break
		}
	}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetChallenge(), []string{
		strconv.FormatInt(latestHeight-wait, 10),
	})
	s.Require().NoError(err)
	return out
}

type yamlChallenge struct {
	Initiator  string `yaml:"initiator"`
	Challenger string `yaml:"challenger"`
	Challengee string `yaml:"challengee"`
	Height     string `yaml:"height"`
	Success    bool   `yaml:"success"`
	Finished   bool   `yaml:"finished"`
}

func (s *PopSelectionE2ETestSuite) sendPoPResult(storedChallenge []byte, success bool) {
	val := s.network.Validators[0]
	var wrapper struct {
		Challenge yamlChallenge `yaml:"challenge"`
	}

	err := yaml.Unmarshal(storedChallenge, &wrapper)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	tmpChallenge := wrapper.Challenge
	var challenge daotypes.Challenge
	challenge.Challengee = tmpChallenge.Challengee
	challenge.Challenger = tmpChallenge.Challenger
	challenge.Initiator = tmpChallenge.Initiator
	challenge.Height, err = strconv.ParseInt(tmpChallenge.Height, 10, 64)
	s.Require().NoError(err)
	challenge.Finished = true
	challenge.Success = success

	msg := daotypes.NewMsgReportPopResult(val.Address.String(), &challenge)
	_, err = e2etestutil.BuildSignBroadcastTx(s.T(), val.Address, msg)
	s.Require().NoError(err)
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionNoActors() {
	out := s.perpareLocalTest()

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), true)
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionOneActors() {
	err := e2etestutil.AttestMachine(s.network, machines[0].name, machines[0].mnemonic, 0, s.feeDenom)
	s.Require().NoError(err)

	out := s.perpareLocalTest()

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), true)
}

func (s *PopSelectionE2ETestSuite) TestPopSelectionTwoActors() {
	err := e2etestutil.AttestMachine(s.network, machines[1].name, machines[1].mnemonic, 1, s.feeDenom)
	s.Require().NoError(err)

	out := s.perpareLocalTest()

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), true)
}

func (s *PopSelectionE2ETestSuite) VerifyTokens(token string) {
	val := s.network.Validators[0]
	// check balance for crddl
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetCmdQueryTotalSupply(), []string{
		fmt.Sprintf("--%s=%s", bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), token)
	assert.Equal(s.T(), "amount: \"17979452050\"\ndenom: "+token+"\n", out.String()) // Total supply 2 * 7990867578 (total supply) + 1 * 1997716894 (challenger) = 17979452050

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		machines[0].address,
		fmt.Sprintf("--%s=%s", bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), token)
	assert.Equal(s.T(), "amount: \"5993150682\"\ndenom: "+token+"\n", out.String()) // 3 * 1997716894 = 5993150682

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		machines[1].address,
		fmt.Sprintf("--%s=%s", bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), token)
	assert.Equal(s.T(), "amount: \"11986301368\"\ndenom: "+token+"\n", out.String()) // 2 * 5993150684 = 11986301368
}

func (s *PopSelectionE2ETestSuite) TestTokenDistribution1() {
	out := s.perpareLocalTest()

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), false)

	out = s.perpareLocalTest()

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), true)

	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	var daoGenState daotypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[daotypes.ModuleName], &daoGenState)

	s.VerifyTokens(daoGenState.Params.StagedDenom)

	// send Reissuance and DistributionResult implicitly
	latestHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)
	for {
		latestHeight, err := s.network.WaitForHeight(latestHeight + 1)
		s.Require().NoError(err)
		// s.Require().NoError(s.network.WaitForNextBlock())
		if latestHeight%s.reissuanceEpochs == s.distributionOffset {
			break
		}
	}
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	s.VerifyTokens(daoGenState.Params.ClaimDenom)
}

func (s *PopSelectionE2ETestSuite) TestTokenRedeemClaim() {
	val := s.network.Validators[0]

	k, err := val.ClientCtx.Keyring.Key(machines[0].name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	// Addr sends CreateRedeemClaim => accepted query redeem claim
	createClaimMsg := daotypes.NewMsgCreateRedeemClaim(addr.String(), "liquidAddress", 10000)
	out, err := lib.BroadcastTxWithFileLock(addr, createClaimMsg)
	s.Require().NoError(err)

	txResponse, err := lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(0), int(txResponse.Code))

	// WaitForBlock => Validator should implicitly send UpdateRedeemClaim
	s.Require().NoError(s.network.WaitForNextBlock())

	// Claim burned on CreateRedeemClaim
	balanceOut, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		addr.String(),
		fmt.Sprintf("--%s=%s", bank.FlagDenom, s.claimDenom),
	})
	s.Require().NoError(err)
	assert.Equal(s.T(), "amount: \"5993140682\"\ndenom: crddl\n", balanceOut.String()) // 3 * 1997716894 - 10000 = 5993140682

	// Addr sends ConfirmRedeemClaim => rejected not claim address
	confirmMsg := daotypes.NewMsgConfirmRedeemClaim(addr.String(), 0, "liquidAddress")
	out, err = lib.BroadcastTxWithFileLock(addr, confirmMsg)
	s.Require().NoError(err)

	txResponse, err = lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(21), int(txResponse.Code))

	// Validator with Claim Address sends ConfirmRedeemClaim => accepted
	valConfirmMsg := daotypes.NewMsgConfirmRedeemClaim(val.Address.String(), 0, "liquidAddress")
	out, err = lib.BroadcastTxWithFileLock(val.Address, valConfirmMsg)
	s.Require().NoError(err)

	txResponse, err = lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(0), int(txResponse.Code))

	// WaitForBlock before query
	s.Require().NoError(s.network.WaitForNextBlock())

	// QueryRedeemClaim
	qOut, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdShowRedeemClaim(), []string{"liquidAddress", "0"})
	s.Require().NoError(err)
	assert.Equal(s.T(), "redeemClaim:\n  amount: \"10000\"\n  beneficiary: liquidAddress\n  confirmed: true\n  creator: plmnt1kp93kns6hs2066d8qw0uz84fw3vlthewt2ck6p\n  id: \"0\"\n  liquidTxHash: \"0000000000000000000000000000000000000000000000000000000000000000\"\n", qOut.String())

	qOut, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdRedeemClaimByLiquidTxHash(), []string{"0000000000000000000000000000000000000000000000000000000000000000"})
	s.Require().NoError(err)
	assert.Equal(s.T(), "redeemClaim:\n  amount: \"10000\"\n  beneficiary: liquidAddress\n  confirmed: true\n  creator: plmnt1kp93kns6hs2066d8qw0uz84fw3vlthewt2ck6p\n  id: \"0\"\n  liquidTxHash: \"0000000000000000000000000000000000000000000000000000000000000000\"\n", qOut.String())
}

func (s *PopSelectionE2ETestSuite) createValAccount(cfg network.Config) (address sdk.AccAddress, err error) {
	buf := bufio.NewReader(os.Stdin)

	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, s.T().TempDir(), buf, cfg.Codec, cfg.KeyringOptions...)
	if err != nil {
		return nil, err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(cfg.SigningAlgo, keyringAlgos)
	if err != nil {
		return nil, err
	}

	mnemonic := cfg.Mnemonics[0]

	record, err := kb.NewAccount("node0", mnemonic, keyring.DefaultBIP39Passphrase, sdk.GetConfig().GetFullBIP44Path(), algo)
	if err != nil {
		return nil, err
	}

	addr, err := record.GetAddress()
	if err != nil {
		return nil, err
	}

	return addr, nil
}
