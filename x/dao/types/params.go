package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(mintAddress string, tokenDenom string, feeDenom string, stagedDenom string,
	claimDenom string, reissuanceAsset string, reissuanceEpochs int64, popEpochs int64,
	distributionOffset int64, distributionAddressEarlyInv string, distributionAddressInvestor string,
	distributionAddressStrategic string, distributionAddressDao string, distributionAddressPop string,
	mqttResponseTimeout int64, claimAddress string) Params {
	return Params{
		MintAddress:     mintAddress,
		TokenDenom:      tokenDenom,
		FeeDenom:        feeDenom,
		StagedDenom:     stagedDenom,
		ClaimDenom:      claimDenom,
		ReissuanceAsset: reissuanceAsset,
		// `ReissuanceEpochs` is a configuration parameter that determines the number of CometBFT epochs
		// required for reissuance. In the context of Planetmint, reissuance refers to the process of
		// issuing new tokens. This configuration parameter specifies the number of epochs (each epoch is 5
		// seconds) that need to pass before reissuance can occur. In this case, `ReissuanceEpochs` is set
		// to 17280, which means that reissuance can occur after 1 day (12*60*24) of epochs.
		ReissuanceEpochs:   reissuanceEpochs,
		PopEpochs:          popEpochs,
		DistributionOffset: distributionOffset,
		// `DistributionOffset` relative to `ReissuanceEpochs`. CometBFT epochs of 5s equate 30 min (12*30)
		// to wait for confirmations on the reissuance
		DistributionAddressEarlyInv:  distributionAddressEarlyInv,
		DistributionAddressInvestor:  distributionAddressInvestor,
		DistributionAddressStrategic: distributionAddressStrategic,
		DistributionAddressDao:       distributionAddressDao,
		DistributionAddressPop:       distributionAddressPop,
		MqttResponseTimeout:          mqttResponseTimeout,
		ClaimAddress:                 claimAddress,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		"plmnt1dyuhg8ldu3d6nvhrvzzemtc3893dys9v9lvdty",
		"plmnt",
		"plmnt",
		"stagedcrddl",
		"crddl",
		"7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9",
		17280,
		24,
		360,
		"vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		"vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		"vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		"vjU8eMzU3JbUWZEpVANt2ePJuPWSPixgjiSj2jDMvkVVQQi2DDnZuBRVX4Ygt5YGBf5zvTWCr1ntdqYH",
		"vjTvXCFSReRsZ7grdsAreRR12KuKpDw8idueQJK9Yh1BYS7ggAqgvCxCgwh13KGK6M52y37HUmvr4GdD",
		2000,
		"plmnt1dyuhg8ldu3d6nvhrvzzemtc3893dys9v9lvdty")
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
