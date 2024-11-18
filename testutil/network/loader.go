package network

import (
	"strings"
	"testing"
	"time"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	pruningtypes "github.com/cosmos/cosmos-sdk/store/pruning/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/golang/mock/gomock"
	"github.com/planetmint/planetmint-go/app"
	"github.com/planetmint/planetmint-go/clients"
	"github.com/planetmint/planetmint-go/monitor"
	monitormocks "github.com/planetmint/planetmint-go/monitor/mocks"
	clientmocks "github.com/planetmint/planetmint-go/testutil/mocks"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/util/mocks"
	elements "github.com/rddl-network/elements-rpc"
	elementsmocks "github.com/rddl-network/elements-rpc/utils/mocks"
	rcctypes "github.com/rddl-network/rddl-claim-service/types"
	scctypes "github.com/rddl-network/shamir-coordinator-service/types"
	"github.com/stretchr/testify/require"
)

// Load creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func Load(t *testing.T, configs ...Config) *Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg Config
	if len(configs) == 0 {
		cfg = LoaderDefaultConfig()
	} else {
		cfg = configs[0]
	}
	validatorTmpDir := t.TempDir()

	// use mock client for testing
	util.MQTTClient = &mocks.MockMQTTClient{}
	monitor.MonitorMQTTClient = &mocks.MockMQTTClient{}
	monitor.SetMqttMonitorInstance(&monitormocks.MockMQTTMonitorClientI{})
	elements.Client = &elementsmocks.MockClient{}
	util.RegisterAssetServiceHTTPClient = &mocks.MockClient{}
	ctrl := gomock.NewController(t)
	claimMock := clientmocks.NewMockIRCClient(ctrl)
	claimMock.EXPECT().PostClaim(gomock.Any(), gomock.Any()).AnyTimes().Return(rcctypes.PostClaimResponse{
		TxID: "0000000000000000000000000000000000000000000000000000000000000000",
	}, nil)
	clients.ClaimServiceClient = claimMock

	shamirMock := clientmocks.NewMockIShamirCoordinatorClient(ctrl)
	shamirMock.EXPECT().SendTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(scctypes.SendTokensResponse{
		TxID: "7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9",
	}, nil)
	shamirMock.EXPECT().ReIssueAsset(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(scctypes.ReIssueResponse{
		TxID: "7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9",
	}, nil)
	shamirMock.EXPECT().IssueMachineNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(scctypes.IssueMachineNFTResponse{
		HexTX:    "0000000000000000000000000000000000000000000000000000000000000000",
		Contract: `{"entity":{"domain":"testnet-assets.rddl.io"}, "issuer_pubkey":"02", "machine_addr":"addr","name":"machine","precicion":8,"version":1}`,
		Asset:    "0000000000000000000000000000000000000000000000000000000000000000",
	}, nil)
	clients.ShamirCoordinatorServiceClient = shamirMock

	// enable application logger in tests
	appLogger := util.GetAppLogger()
	appLogger.SetTestingLogger(t)

	net, err := New(t, validatorTmpDir, cfg)
	// this is only done to support multi validator test
	// race conditions(load/unload) on the CI
	if err != nil && strings.Contains(err.Error(), "bind: address already in use") {
		net, err = New(t, validatorTmpDir, cfg)
	}
	require.NoError(t, err)

	_, err = net.WaitForHeight(1)
	require.NoError(t, err)
	t.Cleanup(net.Cleanup)
	return net
}

// LoaderDefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func LoaderDefaultConfig() Config {
	var (
		encoding = app.MakeEncodingConfig()
		chainID  = "chain-foobarbaz"
	)
	return Config{
		Codec:             encoding.Marshaler,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val ValidatorI) servertypes.Application {
			return app.New(
				val.GetCtx().Logger,
				tmdb.NewMemDB(),
				nil,
				true,
				map[int64]bool{},
				val.GetCtx().Config.RootDir,
				0,
				encoding,
				simtestutil.EmptyAppOptions{},
				baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
				baseapp.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
				baseapp.SetChainID(chainID),
			)
		},
		GenesisState:    app.ModuleBasics.DefaultGenesis(encoding.Marshaler),
		TimeoutCommit:   2 * time.Second,
		ChainID:         chainID,
		NumValidators:   1,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    "0.000003" + sample.FeeDenom,
		AccountTokens:   sdk.TokensFromConsensusPower(10000, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy: pruningtypes.PruningOptionNothing,
		CleanupDir:      true,
		SigningAlgo:     string(hd.Secp256k1Type),
		KeyringOptions:  []keyring.Option{},
		AccountDenom:    sample.FeeDenom,
	}
}
