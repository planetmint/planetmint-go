package machine

import (
	"testing"
	"time"

	"github.com/planetmint/planetmint-go/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestE2EMachineTestSuite(t *testing.T) {
	time.Sleep(2 * time.Second)
	cfg := network.LoaderDefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewE2ETestSuite(cfg))
}

func TestE2EProsumeTestSuite(t *testing.T) {
	cfg := network.LoaderDefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewProsumeE2ETestSuite(cfg))
}
