package dao

import (
	"testing"

	"github.com/planetmint/planetmint-go/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	suite.Run(t, NewE2ETestSuite(cfg))
}

func TestPopE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	suite.Run(t, NewPopSelectionE2ETestSuite(cfg))
}

func TestGasConsumptionE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	suite.Run(t, NewGasConsumptionE2ETestSuite(cfg))
}

func TestRestrictedMsgsE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	suite.Run(t, NewRestrictedMsgsE2ESuite(cfg))
}

func TestAssetDistributionE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	suite.Run(t, NewAssetDistributionE2ETestSuite(cfg))
}
