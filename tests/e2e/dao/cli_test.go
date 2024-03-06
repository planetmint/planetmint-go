package dao

import (
	"testing"

	"github.com/planetmint/planetmint-go/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestE2EDaoTestSuite(t *testing.T) {
	cfg := network.LoaderDefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewE2ETestSuite(cfg))
}

func TestPopE2EDaoTestSuite(t *testing.T) {
	cfg := network.LoaderDefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewPopSelectionE2ETestSuite(cfg))
}

func TestGasConsumptionE2EDaoTestSuite(t *testing.T) {
	cfg := network.LoaderDefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewGasConsumptionE2ETestSuite(cfg))
}

func TestRestrictedMsgsE2EDaoTestSuite(t *testing.T) {
	cfg := network.LoaderDefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewRestrictedMsgsE2ESuite(cfg))
}

func TestAssetDistributionE2EDaoTestSuite(t *testing.T) {
	cfg := network.LoaderDefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewAssetDistributionE2ETestSuite(cfg))
}
