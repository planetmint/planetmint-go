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
func NewParams() Params {
	return Params{
		MintAddress:     "plmnt1dyuhg8ldu3d6nvhrvzzemtc3893dys9v9lvdty",
		TokenDenom:      "plmnt",
		FeeDenom:        "plmnt",
		StagedDenom:     "stagedcrddl",
		ClaimDenom:      "crddl",
		ReissuanceAsset: "7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9",
		// `ReissuanceEpochs` is a configuration parameter that determines the number of CometBFT epochs
		// required for reissuance. In the context of Planetmint, reissuance refers to the process of
		// issuing new tokens. This configuration parameter specifies the number of epochs (each epoch is 5
		// seconds) that need to pass before reissuance can occur. In this case, `ReissuanceEpochs` is set
		// to 17280, which means that reissuance can occur after 1 day (12*60*24) of epochs.
		ReissuanceEpochs: 17280,
		PopEpochs:        24, // 24 CometBFT epochs of 5s equate 120s
		// `DistributionOffset` relative to `ReissuanceEpochs`. CometBFT epochs of 5s equate 30 min (12*30)
		// to wait for confirmations on the reissuance
		DistributionOffset:               360,
		DistributionAddressEarlyInvestor: "vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		DistributionAddressInvestor:      "vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		DistributionAddressStrategic:     "vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		DistributionAddressDao:           "vjU8eMzU3JbUWZEpVANt2ePJuPWSPixgjiSj2jDMvkVVQQi2DDnZuBRVX4Ygt5YGBf5zvTWCr1ntdqYH",
		DistributionAddressPop:           "vjTvXCFSReRsZ7grdsAreRR12KuKpDw8idueQJK9Yh1BYS7ggAqgvCxCgwh13KGK6M52y37HUmvr4GdD",
		MqttResponseTimeout:              2000, // the value is defined in milliseconds
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
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
