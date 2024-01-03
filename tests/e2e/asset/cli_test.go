package asset

import (
	"testing"

	"github.com/planetmint/planetmint-go/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestE2ETestSuite(t *testing.T) {
	t.Parallel()
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewE2ETestSuite(cfg))
}

func TestRESTMachineAttestationSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	suite.Run(t, NewRestE2ETestSuite(cfg))
}
