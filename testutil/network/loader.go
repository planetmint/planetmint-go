package network

import (
	"bytes"
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
	"github.com/planetmint/planetmint-go/app"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/util/mocks"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	elements "github.com/rddl-network/elements-rpc"
	elementsmocks "github.com/rddl-network/elements-rpc/utils/mocks"
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
	elements.Client = &elementsmocks.MockClient{}

	// enable application logger in tests
	appLogger := util.GetAppLogger()
	appLogger.SetTestingLogger(t)

	// set the proper root dir for the test environment so that the abci.go logic works
	conf := config.GetConfig()
	conf.SetRoot(validatorTmpDir + "/node0/simd")

	net, err := New(t, validatorTmpDir, cfg)
	require.NoError(t, err)

	conf.ValidatorAddress = net.Validators[0].Address.String()
	// set missing validator client context values for sending txs
	var output bytes.Buffer
	net.Validators[0].ClientCtx.BroadcastMode = "sync"
	net.Validators[0].ClientCtx.FromAddress = net.Validators[0].Address
	net.Validators[0].ClientCtx.FromName = net.Validators[0].Moniker
	net.Validators[0].ClientCtx.NodeURI = net.Validators[0].RPCAddress
	net.Validators[0].ClientCtx.Output = &output
	net.Validators[0].ClientCtx.SkipConfirm = true

	var daoGenState daotypes.GenesisState
	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[daotypes.ModuleName], &daoGenState)

	libConfig := lib.GetConfig()
	libConfig.SetClientCtx(net.Validators[0].ClientCtx)
	libConfig.SetRoot(validatorTmpDir + "/node0/simd")

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