package dao

import (
	"math"
	"strconv"

	"github.com/planetmint/planetmint-go/config"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	"github.com/planetmint/planetmint-go/testutil/network"
	daocli "github.com/planetmint/planetmint-go/x/dao/client/cli"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/suite"
)

type AssetDistributionE2ETestSuite struct {
	suite.Suite

	cfg                network.Config
	network            *network.Network
	reissaunceEpochs   int64
	distributionOffset int64
}

func NewAssetDistributionE2ETestSuite(cfg network.Config) *AssetDistributionE2ETestSuite {
	return &AssetDistributionE2ETestSuite{cfg: cfg}
}

func (s *AssetDistributionE2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	// set epochs: make sure to start after initial height of 7
	s.distributionOffset = 5
	s.reissaunceEpochs = 10

	var daoGenState daotypes.GenesisState
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[daotypes.ModuleName], &daoGenState)
	daoGenState.Params.DistributionOffset = s.distributionOffset
	daoGenState.Params.ReissuanceEpochs = s.reissaunceEpochs
	s.cfg.GenesisState[daotypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(&daoGenState)

	conf := config.GetConfig()
	conf.DistributionOffset = int(s.distributionOffset)
	conf.ReissuanceEpochs = int(s.reissaunceEpochs)

	s.network = network.Load(s.T(), s.cfg)
}

func (s *AssetDistributionE2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suites")
}

func (s *AssetDistributionE2ETestSuite) TestAssetDistribution() {
	val := s.network.Validators[0]

	latestHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)

	// wait so that we see exactly on reissuance/distribution round, e.g.
	// wait = ceil((10 - 5) / 2) = ceil(2.5) = 3
	wait := int64(math.Ceil((float64(s.reissaunceEpochs - s.distributionOffset)) / 2.0))
	height := s.reissaunceEpochs + s.distributionOffset + wait
	for {
		latestHeight, err = s.network.WaitForHeight(latestHeight + 1)
		s.Require().NoError(err)

		if latestHeight == height {
			break
		}
	}
	errmsg := "rpc error: code = NotFound desc = distribution not found: key not found"
	testCases := []struct {
		name          string
		requestHeight int64
		expectedErr   string
	}{
		{
			"request height too low",
			s.distributionOffset,
			errmsg,
		},
		{
			"wrong request height",
			height,
			errmsg,
		},
		{
			"request height too high",
			2*s.reissaunceEpochs + s.distributionOffset,
			errmsg,
		},
		{
			"valid distribution request",
			s.reissaunceEpochs + s.distributionOffset,
			"",
		},
	}

	for _, tc := range testCases {
		_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, daocli.CmdGetDistribution(), []string{
			strconv.FormatInt(tc.requestHeight, 10),
		})
		if tc.expectedErr == "" {
			s.Require().NoError(err)
		} else {
			s.Require().Error(err)
		}
	}
}
