package asset

import (
	"fmt"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/testutil"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/stretchr/testify/suite"

	assettypes "github.com/planetmint/planetmint-go/x/asset/types"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

type RestE2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewRestE2ETestSuite(cfg network.Config) *RestE2ETestSuite {
	return &RestE2ETestSuite{cfg: cfg}
}

// SetupSuite initializes asset E2ETestSuite
func (s *RestE2ETestSuite) SetupSuite() {
	conf := config.GetConfig()
	conf.FeeDenom = "stake"

	s.T().Log("setting up e2e test suite")

	s.network = network.New(s.T())
	err := e2etestutil.AttestMachine(s.network, sample.Name, sample.Mnemonic, 0)
	s.Require().NoError(err)

	val := s.network.Validators[0]
	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)

	addr, _ := k.GetAddress()

	clientCtx := val.ClientCtx.
		WithFromAddress(addr).
		WithFromName(sample.Name)
	libConfig := lib.GetConfig()
	libConfig.SetClientCtx(clientCtx)
}

// TearDownSuite clean up after testing
func (s *RestE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

// TestNotarizeAssetREST notarizes asset over REST endpoint
func (s *RestE2ETestSuite) TestNotarizeAssetREST() {
	val := s.network.Validators[0]

	// Create Msg
	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)

	addr, err := k.GetAddress()
	s.Require().NoError(err)
	cid := sample.Asset()
	testCases := []struct {
		name             string
		msg              assettypes.MsgNotarizeAsset
		rawLog           string
		expectCheckTxErr bool
	}{
		{
			"invalid address",
			assettypes.MsgNotarizeAsset{
				Creator: "invalid creator address",
				Cid:     cid,
			},
			"invalid address",
			true,
		},
		{
			"machine not found",
			assettypes.MsgNotarizeAsset{
				Creator: "plmnt1v5394e8vmfrp4qzdav7xkze0f567w3tsgxf09j",
				Cid:     cid,
			},
			"machine not found",
			true,
		},
		{
			"valid notarization",
			assettypes.MsgNotarizeAsset{
				Creator: addr.String(),
				Cid:     cid,
			},
			"[]",
			true,
		},
	}

	for _, tc := range testCases {
		// Prepare Tx
		txBytes, err := testutil.PrepareTx(val, &tc.msg, sample.Name)
		s.Require().NoError(err)

		// Broadcast Tx
		broadcastTxResponse, err := testutil.BroadcastTx(val, txBytes)
		s.Require().NoError(err)

		s.Require().NoError(s.network.WaitForNextBlock())

		tx, err := testutil.GetRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", val.APIAddress, broadcastTxResponse.TxResponse.TxHash))
		s.Require().NoError(err)

		if !tc.expectCheckTxErr {
			var txRes txtypes.GetTxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(tx, &txRes)
			s.Require().NoError(err)
			s.Require().Contains(txRes.TxResponse.RawLog, tc.rawLog)
		} else {
			s.Require().Contains(broadcastTxResponse.TxResponse.RawLog, tc.rawLog)
		}
	}
}
