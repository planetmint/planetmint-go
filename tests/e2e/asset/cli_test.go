package asset

import (
	"testing"
	"time"

	"github.com/planetmint/planetmint-go/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestE2EAssetTestSuite(t *testing.T) {
	time.Sleep(2 * time.Second)
	cfg := network.LoaderDefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewE2ETestSuite(cfg))
}
