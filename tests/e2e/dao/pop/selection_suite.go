package pop

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/monitor"
	"github.com/planetmint/planetmint-go/testutil"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/util/mocks"
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

type SelectionE2ETestSuite struct {
	suite.Suite

	cfg                network.Config
	network            *network.Network
	popEpochs          int64
	reissuanceEpochs   int64
	distributionOffset int64
	claimDenom         string
	feeDenom           string
	errormsg           string
}

func NewSelectionE2ETestSuite(cfg network.Config) *SelectionE2ETestSuite {
	testsuite := &SelectionE2ETestSuite{cfg: cfg}
	testsuite.errormsg = "--%s=%s"
	return testsuite
}

func (s *SelectionE2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e dao pop selection test suite")

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
	daoGenState.Params.ClaimAddress = valAddr.String()
	s.cfg.GenesisState[daotypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&daoGenState)

	// setting up stagedClaims that are not part of PoP issuance (i.e.: past unresolved claims)
	machineBalances := []banktypes.Balance{
		{Address: machines[0].address, Coins: sdk.NewCoins(sdk.NewCoin(daoGenState.Params.StagedDenom, sdkmath.NewInt(10000)))},
		{Address: machines[1].address, Coins: sdk.NewCoins(sdk.NewCoin(daoGenState.Params.StagedDenom, sdkmath.NewInt(10000)))},
	}

	var bankGenState banktypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[banktypes.ModuleName], &bankGenState)
	bankGenState.Balances = append(bankGenState.Balances, machineBalances...)
	s.cfg.GenesisState[banktypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&bankGenState)

	s.network = network.Load(s.T(), s.cfg)
}

// TearDownSuite clean up after testing
func (s *SelectionE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e dao pop selection test suite")
}

func (s *SelectionE2ETestSuite) perpareLocalTest() testutil.BufferWriter {
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

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.GetCmdChallenge(), []string{
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

func (s *SelectionE2ETestSuite) sendPoPResult(storedChallenge []byte, success bool) {
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

	machineName := machines[0].name
	if challenge.Challenger != machines[0].address {
		machineName = machines[1].name
	}
	k, err := val.ClientCtx.Keyring.Key(machineName)
	s.Require().NoError(err)
	challengerAccAddress, _ := k.GetAddress()

	msg := daotypes.NewMsgReportPopResult(challengerAccAddress.String(), &challenge)
	_, err = e2etestutil.BuildSignBroadcastTx(s.T(), challengerAccAddress, msg)
	s.Require().NoError(err)
}

func (s *SelectionE2ETestSuite) TestPopSelectionNoActors() {
	out := s.perpareLocalTest()

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)
}

func (s *SelectionE2ETestSuite) TestPopSelectionOneActors() {
	err := monitor.AddParticipant(machines[0].address, time.Now().Unix())
	s.Require().NoError(err)
	err = e2etestutil.AttestMachine(s.network, machines[0].name, machines[0].mnemonic, 0, s.feeDenom)
	s.Require().NoError(err)

	out := s.perpareLocalTest()

	assert.NotContains(s.T(), out.String(), machines[0].address)
	assert.NotContains(s.T(), out.String(), machines[1].address)
}

func (s *SelectionE2ETestSuite) TestPopSelectionTwoActors() {
	err := monitor.AddParticipant(machines[1].address, time.Now().Unix())
	s.Require().NoError(err)
	err = e2etestutil.AttestMachine(s.network, machines[1].name, machines[1].mnemonic, 1, s.feeDenom)
	s.Require().NoError(err)

	out := s.perpareLocalTest()

	assert.Contains(s.T(), out.String(), machines[0].address)
	assert.Contains(s.T(), out.String(), machines[1].address)
	s.sendPoPResult(out.Bytes(), true)
}

func (s *SelectionE2ETestSuite) VerifyTokens(token string) {
	val := s.network.Validators[0]
	// check balance for crddl
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetCmdQueryTotalSupply(), []string{
		fmt.Sprintf(s.errormsg, bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), token)

	// Account for 1 additional unfinished PoP when checking balances after distribution
	if token == s.claimDenom {
		assert.Equal(s.T(), "amount: \"18579472050\"\ndenom: "+token+"\n", out.String()) // Total supply 2 * 7990867578 (total supply) + 1 * 1997716894 (challenger) + 6 * 100000000 (validator) + 2 * 10000 (past unresolved claims) = 18579472050
	} else {
		assert.Equal(s.T(), "amount: \"18479472050\"\ndenom: "+token+"\n", out.String()) // Total supply 2 * 7990867578 (total supply) + 1 * 1997716894 (challenger) + 5 * 100000000 (validator) + 2 * 10000 (past unresolved claims) = 18479472050
	}

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		machines[0].address,
		fmt.Sprintf(s.errormsg, bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), token)
	assert.Equal(s.T(), "amount: \"5993160682\"\ndenom: "+token+"\n", out.String()) // 3 * 1997716894 + 1 * 10000 = 5993160682

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		machines[1].address,
		fmt.Sprintf(s.errormsg, bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), token)
	assert.Equal(s.T(), "amount: \"11986311368\"\ndenom: "+token+"\n", out.String()) // 2 * 5993150684 + 1 * 10000 = 11986311368

	out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, bank.GetBalancesCmd(), []string{
		val.Address.String(),
		fmt.Sprintf(s.errormsg, bank.FlagDenom, token),
	})
	s.Require().NoError(err)
	assert.Contains(s.T(), out.String(), token)

	// Account for 1 additional unfinished PoP when checking balances after distribution
	if token == s.claimDenom {
		assert.Equal(s.T(), "amount: \"600000000\"\ndenom: "+token+"\n", out.String()) // 6 * 100000000
	} else {
		assert.Equal(s.T(), "amount: \"500000000\"\ndenom: "+token+"\n", out.String()) // 5 * 100000000
	}
}

func (s *SelectionE2ETestSuite) TestTokenDistribution1() {
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
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	s.VerifyTokens(daoGenState.Params.ClaimDenom)
}

func (s *SelectionE2ETestSuite) TestTokenRedeemClaim() {
	val := s.network.Validators[0]

	k, err := val.ClientCtx.Keyring.Key(machines[0].name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	// Addr sends CreateRedeemClaim => accepted query redeem claim
	createClaimMsg := daotypes.NewMsgCreateRedeemClaim(addr.String(), "liquidAddress")
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
		fmt.Sprintf(s.errormsg, bank.FlagDenom, s.claimDenom),
	})
	s.Require().NoError(err)
	assert.Equal(s.T(), "amount: \"0\"\ndenom: crddl\n", balanceOut.String()) // consumes all claims

	// Addr sends ConfirmRedeemClaim => rejected not claim address
	confirmMsg := daotypes.NewMsgConfirmRedeemClaim(addr.String(), 0, "liquidAddress")
	out, err = lib.BroadcastTxWithFileLock(addr, confirmMsg)
	s.Require().NoError(err)

	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock()) // added another waiting block to pass CI test cases (they are a bit slower)

	_, err = clitestutil.GetRawLogFromTxOut(val, out)
	s.Require().ErrorContains(err, "failed to execute message; message index: 0: expected: plmnt19cl05ztgt8ey6v86hjjjn3thfmpu6q2xtveehc; got: plmnt1kp93kns6hs2066d8qw0uz84fw3vlthewt2ck6p: invalid claim address")

	// Validator with Claim Address sends ConfirmRedeemClaim => accepted
	valConfirmMsg := daotypes.NewMsgConfirmRedeemClaim(val.Address.String(), 0, "liquidAddress")
	out, err = lib.BroadcastTxWithFileLock(val.Address, valConfirmMsg)
	s.Require().NoError(err)

	txResponse, err = lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(0), int(txResponse.Code))

	// WaitForBlock before query (2 blocks since 3 validators)
	s.Require().NoError(s.network.WaitForNextBlock())
	s.Require().NoError(s.network.WaitForNextBlock())

	// QueryRedeemClaim
	qOut, err := clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.GetCmdShowRedeemClaim(), []string{"liquidAddress", "0"})
	s.Require().NoError(err)
	assert.Equal(s.T(), "redeemClaim:\n  amount: \"5993160682\"\n  beneficiary: liquidAddress\n  confirmed: true\n  creator: plmnt1kp93kns6hs2066d8qw0uz84fw3vlthewt2ck6p\n  id: \"0\"\n  liquidTxHash: \"0000000000000000000000000000000000000000000000000000000000000000\"\n", qOut.String())

	qOut, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.GetCmdRedeemClaimByLiquidTxHash(), []string{"0000000000000000000000000000000000000000000000000000000000000000"})
	s.Require().NoError(err)
	assert.Equal(s.T(), "redeemClaim:\n  amount: \"5993160682\"\n  beneficiary: liquidAddress\n  confirmed: true\n  creator: plmnt1kp93kns6hs2066d8qw0uz84fw3vlthewt2ck6p\n  id: \"0\"\n  liquidTxHash: \"0000000000000000000000000000000000000000000000000000000000000000\"\n", qOut.String())

	// Make sure "Publish" has been called with PoPInit cmnd
	calls := mocks.GetCallLog()

	var popInitCalls []mocks.Call
	regex := regexp.MustCompile(`cmnd\/[a-zA-Z0-9]{15,50}\/PoPInit`)
	for _, call := range calls {
		if call.FuncName != "Publish" {
			continue
		}

		cmnd, ok := call.Params[0].(string)
		if !ok {
			assert.True(s.T(), ok) // fails test case if !ok
			continue
		}

		if regex.MatchString(cmnd) {
			popInitCalls = append(popInitCalls, call)
		}
	}
	assert.Greater(s.T(), len(popInitCalls), 0)
}

func (s *SelectionE2ETestSuite) createValAccount(cfg network.Config) (address sdk.AccAddress, err error) {
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
