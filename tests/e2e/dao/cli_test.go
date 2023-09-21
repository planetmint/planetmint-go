package dao

import (
	"github.com/planetmint/planetmint-go/testutil/network"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	suite.Run(t, NewE2ETestSuite(cfg))
}
