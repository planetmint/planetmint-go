package asset

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"planetmint-go/testutil/network"
	"planetmint-go/testutil/sample"
	"regexp"

	clitestutil "planetmint-go/testutil/cli"
	assetcli "planetmint-go/x/asset/client/cli"
	machinecli "planetmint-go/x/machine/client/cli"

	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"sigs.k8s.io/yaml"
)

// Queryable pubkey for TestNotarizeAsset
const mnemonic = "helmet hedgehog lab actor weekend elbow pelican valid obtain hungry rocket decade tower gallery fit practice cart cherry giggle hair snack glance bulb farm"

// E2ETestSuite struct definition of asset suite
type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// NewE2ETestSuite returns configured asset E2ETestSuite
func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

// SetupSuite initializes asset E2ETestSuite
func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	s.network = network.New(s.T())
	val := s.network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, err := kb.NewAccount("machine", mnemonic, keyring.DefaultBIP39Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)
	s.Require().NoError(err)
	pk, err := account.GetPubKey()
	pkHex := hex.EncodeToString(pk.Bytes())
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

	machine := sample.Machine("machine", pkHex)
	machineJSON, err := json.Marshal(&machine)
	s.Require().NoError(err)

	args = []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, "machine"),
		fmt.Sprintf("--%s=%s", flags.FlagFees, "2stake"),
		"--yes",
		string(machineJSON),
	}

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, machinecli.CmdAttestMachine(), args)
	s.Require().NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())
}

// TearDownSuite clean up after testing
func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

// Needed to export private key from Keyring
type unsafeExporter interface {
	ExportPrivateKeyObject(uid string) (types.PrivKey, error)
}

// TestNotarizeAsset notarizes asset over cli
func (s *E2ETestSuite) TestNotarizeAsset() {
	val := s.network.Validators[0]

	privKey, err := val.ClientCtx.Keyring.(unsafeExporter).ExportPrivateKeyObject("machine")
	s.Require().NoError(err)

	sk := hex.EncodeToString(privKey.Bytes())

	cidHash, signature := sample.Asset(sk)

	testCases := []struct {
		name   string
		args   []string
		rawLog string
	}{
		{
			"machine not found",
			[]string{
				cidHash,
				signature,
				"pubkey",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "machine"),
				fmt.Sprintf("--%s=%s", flags.FlagFees, "2stake"),
				"--yes",
			},
			"machine not found",
		},
		{
			"invalid signature",
			[]string{
				"cid",
				"signature",
				hex.EncodeToString(privKey.PubKey().Bytes()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "machine"),
				fmt.Sprintf("--%s=%s", flags.FlagFees, "2stake"),
				"--yes",
			},
			"invalid signature",
		},
		{
			"valid notarization",
			[]string{
				cidHash,
				signature,
				hex.EncodeToString(privKey.PubKey().Bytes()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "machine"),
				fmt.Sprintf("--%s=%s", flags.FlagFees, "2stake"),
				"--yes",
			},
			"planetmintgo.asset.MsgNotarizeAsset",
		},
	}

	for _, tc := range testCases {
		out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, assetcli.CmdNotarizeAsset(), tc.args)
		s.Require().NoError(err)
		// Hack: numbers come back as strings and cannot be unmarshalled into TxResponse struct
		m := regexp.MustCompile(`"([0-9]+?)"`)
		str := m.ReplaceAllString(out.String(), "${1}")

		var txResponse sdk.TxResponse
		err = json.Unmarshal([]byte(str), &txResponse)
		s.Require().NoError(err)

		s.Require().NoError(s.network.WaitForNextBlock())
		args := []string{
			txResponse.TxHash,
		}
		out, err = clitestutil.ExecTestCLICmd(val.ClientCtx, authcmd.QueryTxCmd(), args)
		s.Require().NoError(err)

		str = m.ReplaceAllString(out.String(), "${1}")
		// Need to convert to JSON first, because TxResponse struct lacks `yaml:"height,omitempty"`, etc.
		j, err := yaml.YAMLToJSON([]byte(str))
		s.Require().NoError(err)

		err = json.Unmarshal(j, &txResponse)
		s.Require().NoError(err)

		assert.Contains(s.T(), txResponse.RawLog, tc.rawLog)
	}
}
