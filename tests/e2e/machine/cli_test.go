package machine

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/planetmint/planetmint-go/app"
	"github.com/stretchr/testify/suite"
)

func TestE2EMachineTestSuite(t *testing.T) {
	cfg, err := network.DefaultConfigWithAppConfig(app.AppConfig())
	if err != nil {
		panic("error while setting up application config")
	}
	cfg.NumValidators = 3
	cfg.MinGasPrices = "0.000003stake"
	suite.Run(t, NewE2ETestSuite(cfg))
}
