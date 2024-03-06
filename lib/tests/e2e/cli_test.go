package machine

import (
	"testing"

	"github.com/planetmint/planetmint-go/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestE2ELibTestSuite(t *testing.T) {
	t.Parallel()
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewE2ETestSuite(cfg))
}
