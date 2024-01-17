package dao

import (
	"fmt"
	"math"
	"strconv"

	"github.com/planetmint/planetmint-go/config"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daocli "github.com/planetmint/planetmint-go/x/dao/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AssetDistributionE2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewAssetDistributionE2ETestSuite(cfg network.Config) *AssetDistributionE2ETestSuite {
	return &AssetDistributionE2ETestSuite{cfg: cfg}
}

func (s *AssetDistributionE2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	// set fee denomination
	conf := config.GetConfig()
	conf.FeeDenom = sample.FeeDenom
	// set epochs: make sure to start after initial height of 7
	conf.DistributionOffset = 5
	conf.ReissuanceEpochs = 10

	s.network = network.New(s.T(), s.cfg)
}

func (s *AssetDistributionE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suites")
}

func (s *AssetDistributionE2ETestSuite) TestAssetDistribution() {
	val := s.network.Validators[0]
	conf := config.GetConfig()

	latestHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)

	// wait so that we see exactly on reissuance/distribution round, e.g.
	// wait = ceil((10 - 5) / 2) = ceil(2.5) = 3
	wait := int(math.Ceil((float64(conf.ReissuanceEpochs) - float64(conf.DistributionOffset)) / 2.0))
	height := conf.ReissuanceEpochs + conf.DistributionOffset + wait
	for {
		latestHeight, err = s.network.WaitForHeight(latestHeight + 1)
		s.Require().NoError(err)

		if latestHeight == int64(height) {
			break
		}
	}

	testCases := []struct {
		name          string
		requestHeight int
		expectedErr   string
	}{
		{
			"request height too low",
			conf.DistributionOffset,
			fmt.Sprintf("distribution wrong height: must be equal to or greater then %d",
				conf.ReissuanceEpochs+conf.DistributionOffset),
		},
		{
			"wrong request height",
			height,
			fmt.Sprintf("distribution wrong height: must equal to (n * %d) + %d, where n = 1, 2, 3, and so on",
				conf.ReissuanceEpochs, conf.DistributionOffset),
		},
		{
			"request height too high",
			2*conf.ReissuanceEpochs + conf.DistributionOffset,
			fmt.Sprintf("height %d must be less than or equal to the current blockchain height %d",
				2*conf.ReissuanceEpochs+conf.DistributionOffset, height),
		},
		{
			"valid distribution request",
			conf.ReissuanceEpochs + conf.DistributionOffset,
			"",
		},
	}

	for _, tc := range testCases {
		_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetDistribution(), []string{
			strconv.Itoa(tc.requestHeight),
		})
		if tc.expectedErr == "" {
			s.Require().NoError(err)
		} else {
			s.Require().Error(err)
			assert.Equal(s.T(), tc.expectedErr, err.Error())
		}
	}
}
