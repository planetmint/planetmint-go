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
