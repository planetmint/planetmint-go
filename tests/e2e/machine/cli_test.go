package machine

import (
	"testing"

	"cosmossdk.io/depinject"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/planetmint/planetmint-go/app"
	"github.com/planetmint/planetmint-go/app/ante"
	"github.com/stretchr/testify/suite"

	pruningtypes "cosmossdk.io/store/pruning/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

var (
	cfg       network.Config
	appConfig depinject.Config
)

func TestE2EMachineTestSuite(t *testing.T) {
	appConfig = app.AppConfig()
	var err error
	cfg, err = network.DefaultConfigWithAppConfig(appConfig)
	if err != nil {
		panic("error while setting up application config")
	}
	cfg.NumValidators = 3
	cfg.MinGasPrices = "0.000003stake"

	cfg.AppConstructor = appConstructor

	suite.Run(t, NewE2ETestSuite(cfg))
}

func appConstructor(val network.ValidatorI) servertypes.Application {
	// we build a unique app instance for every validator here
	var appBuilder *runtime.AppBuilder
	if err := depinject.Inject(
		depinject.Configs(
			appConfig,
			depinject.Supply(val.GetCtx().Logger),
		),
		&appBuilder); err != nil {
		panic(err)
	}
	app := appBuilder.Build(
		dbm.NewMemDB(),
		nil,
		baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
		baseapp.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
		baseapp.SetChainID(cfg.ChainID),
	)

	testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})

	if err := app.Load(true); err != nil {
		panic(err)
	}

	anteOpts := ante.HandlerOptions{}
	anteHandler, err := ante.NewAnteHandler(anteOpts)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)

	return app
}
