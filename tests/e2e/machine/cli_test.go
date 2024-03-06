package machine

import (
	"testing"

	"github.com/planetmint/planetmint-go/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestE2EMachineTestSuite(t *testing.T) {
	t.Parallel()
	cfg := network.LoaderDefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewE2ETestSuite(cfg))
}
